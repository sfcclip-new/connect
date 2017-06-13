package server

import (
	"fmt"
	"net/http"
	"os"
	"strconv"

	"github.com/akkyie/connect.sfcclip.net/model"
	"github.com/akkyie/connect.sfcclip.net/resource"
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/core"
	"github.com/go-xorm/xorm"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"github.com/manyminds/api2go"
	"github.com/manyminds/api2go-adapter/gorillamux"
	_ "github.com/mattn/go-sqlite3"
	log "github.com/sirupsen/logrus"
)

// Server handles http requests
type Server struct {
	orm     *xorm.Engine
	handler http.Handler
}

// NewServer returns a new server
func NewServer(production bool) (*Server, error) {
	var (
		orm *xorm.Engine
		err error
	)
	if production {
		if err = godotenv.Load(); err != nil {
			return nil, err
		}

		user := os.Getenv("MYSQL_USER")
		password := os.Getenv("MYSQL_PASSWORD")
		host := os.Getenv("MYSQL_HOST")
		database := os.Getenv("MYSQL_DATABASE")

		orm, err = xorm.NewEngine("mysql", fmt.Sprintf("%s:%s@%s/%s", user, password, host, database))
		if err != nil {
			return nil, err
		}
	} else {
		orm, err = xorm.NewEngine("sqlite3", "./test.db")
		if err != nil {
			return nil, err
		}
	}

	orm.ShowSQL(true)
	orm.ShowExecTime(true)
	orm.SetMapper(core.GonicMapper{})

	router := mux.NewRouter()
	router.PathPrefix("/console").Handler(http.StripPrefix("/console", http.FileServer(http.Dir("./console"))))
	router.HandleFunc("/any/{GroupID}", handleAnyUnitInGroupRequest(orm))
	router.HandleFunc("/img/{UnitID}", handleImageRequest(orm))
	router.HandleFunc("/open/{UnitID}", handleOpenRequest(orm))

	var handler = handlers.RecoveryHandler(
		handlers.PrintRecoveryStack(true),
	)(router)

	handler = handlers.CORS(
		handlers.AllowedOrigins([]string{"*"}),
		handlers.AllowedHeaders([]string{"Authorization", "Content-Type"}),
		handlers.AllowedMethods([]string{"POST", "GET", "PATCH", "DELETE"}),
	)(handler)

	api := api2go.NewAPIWithRouting(
		"api",
		api2go.NewStaticResolver("/"),
		gorillamux.New(router),
	)

	api.AddResource(model.Unit{}, resource.NewUnitResource(orm))
	if err := orm.Sync(model.Unit{}); err != nil {
		return nil, err
	}

	api.AddResource(model.Group{}, resource.NewGroupResource(orm))
	if err := orm.Sync(model.Group{}); err != nil {
		return nil, err
	}

	if err := orm.Sync(model.Record{}); err != nil {
		return nil, err
	}

	return &Server{
		orm:     orm,
		handler: handler,
	}, nil
}

// Start server
func (s Server) Start(port int) error {
	log.Infof("Listening on :%d", port)
	return http.ListenAndServe(fmt.Sprintf(":%d", port), s.handler)
}

func anyUnit(orm *xorm.Engine, groupID int64) (*model.Unit, error) {
	group := model.Group{ID: groupID}
	if _, err := orm.Get(&group); err != nil {
		return nil, err
	}

	var units []model.Unit
	if err := orm.In("id", group.UnitIDs).Asc("image_count").Limit(1).Find(&units); err != nil {
		return nil, err
	}

	if len(units) != 1 {
		return nil, nil
	}

	return &units[0], nil
}

func handleAnyUnitInGroupRequest(orm *xorm.Engine) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		groupIDStr, has := vars["GroupID"]
		if !has {
			http.Error(w, "Group ID not found", http.StatusBadRequest)
			return
		}

		groupID, err := strconv.ParseInt(groupIDStr, 10, 0)
		if err != nil {
			http.Error(w, "Invalid group ID", http.StatusBadRequest)
			return
		}

		unit, err := anyUnit(orm, groupID)
		if err != nil {
			http.NotFound(w, r)
			return
		}

		fmt.Fprintf(w, "%d\n", unit.ID)
	}
}

func handleImageRequest(orm *xorm.Engine) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		unitIDStr, has := vars["UnitID"]
		if !has {
			http.Error(w, "Unit ID not found", http.StatusBadRequest)
			return
		}

		unitID, err := strconv.ParseInt(unitIDStr, 10, 0)
		if err != nil {
			http.Error(w, "Invalid unit ID", http.StatusBadRequest)
			return
		}

		unit := model.Unit{ID: unitID}
		has, err = orm.ID(unitID).Get(&unit)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		if !has {
			http.NotFound(w, r)
			return
		}

		go recordAccess(orm, unit, model.ImageAccessType, r)

		http.Redirect(w, r, unit.ImageURL, http.StatusSeeOther)
	}
}

func handleOpenRequest(orm *xorm.Engine) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		unitIDStr, has := vars["UnitID"]
		if !has {
			http.Error(w, "Unit ID not found", http.StatusBadRequest)
			return
		}

		unitID, err := strconv.ParseInt(unitIDStr, 10, 0)
		if err != nil {
			http.Error(w, "Invalid unit ID", http.StatusBadRequest)
			return
		}

		unit := model.Unit{ID: unitID}
		has, err = orm.ID(unitID).Get(&unit)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		if !has {
			http.NotFound(w, r)
			return
		}

		go recordAccess(orm, unit, model.OpenAccessType, r)

		http.Redirect(w, r, unit.TargetURL+"?utm_source=sfcclip", http.StatusSeeOther)
	}
}

func recordAccess(orm *xorm.Engine, unit model.Unit, accessType model.AccessType, r *http.Request) {
	record := model.NewRecord(unit.ID, accessType, r)
	if _, err := orm.Insert(record); err != nil {
		log.Fatal(err)
		return
	}

	if accessType == model.ImageAccessType {
		var (
			count int64
			err   error
		)

		if count, err = orm.Where("unit_id = ?", unit.ID).Count(new(model.Record)); err != nil {
			log.Fatal(err)
			return
		}

		unit.ImageAccessCount = count + 1
		if _, err := orm.ID(unit.ID).Update(unit); err != nil {
			log.Fatal(err)
			return
		}
	}
}
