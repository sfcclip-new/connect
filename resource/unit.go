package resource

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/akkyie/connect.sfcclip.net/model"
	"github.com/go-xorm/xorm"
	"github.com/manyminds/api2go"
	log "github.com/sirupsen/logrus"
)

// UnitResource provides routing for units
type UnitResource struct {
	orm *xorm.Engine
}

// NewUnitResource returns new UnitResource
func NewUnitResource(orm *xorm.Engine) *UnitResource {
	return &UnitResource{orm}
}

// FindAll to satisfy api2go data source interface
func (r UnitResource) FindAll(req api2go.Request) (api2go.Responder, error) {
	var units []model.Unit
	if groupIDs, has := req.QueryParams["groupsID"]; has && len(groupIDs) > 0 {
		groupID, err := strconv.ParseInt(groupIDs[0], 10, 0)
		if err != nil {
			return &Response{}, err
		}
		group := model.Group{ID: groupID}
		if has, err := r.orm.Get(&group); !has || err != nil {
			return &Response{}, err
		}
		if err := r.orm.In("id", group.UnitIDs).Find(&units); err != nil {
			return &Response{}, err
		}
	} else {
		if err := r.orm.Find(&units); err != nil {
			return &Response{}, err
		}
	}

	return &Response{Res: units}, nil
}

// FindOne to satisfy `api2go.DataSource` interface
// this method should return the user with the given ID, otherwise an error
func (r UnitResource) FindOne(ID string, req api2go.Request) (api2go.Responder, error) {
	id, err := strconv.ParseInt(ID, 10, 0)
	if err != nil {
		return &Response{}, err
	}
	unit := model.Unit{ID: id}
	if _, err := r.orm.Get(&unit); err != nil {
		return &Response{}, err
	}
	return &Response{Res: unit}, nil
}

// Create method to satisfy `api2go.DataSource` interface
func (r UnitResource) Create(obj interface{}, req api2go.Request) (api2go.Responder, error) {
	unit, ok := obj.(model.Unit)
	if !ok {
		return &Response{}, api2go.NewHTTPError(errors.New(""), "", http.StatusBadRequest)
	}
	if _, err := r.orm.Insert(&unit); err != nil {
		return &Response{}, err
	}
	log.Info(unit)
	return &Response{Res: unit, Code: http.StatusCreated}, nil
}

// Delete to satisfy `api2go.DataSource` interface
func (r UnitResource) Delete(id string, req api2go.Request) (api2go.Responder, error) {
	ID, err := strconv.ParseInt(id, 10, 0)
	if err != nil {
		return &Response{}, err
	}
	if _, err := r.orm.Id(ID).Delete(&model.Unit{}); err != nil {
		return &Response{}, err
	}
	return &Response{Code: http.StatusNoContent}, nil
}

//Update stores all changes on the user
func (r UnitResource) Update(obj interface{}, req api2go.Request) (api2go.Responder, error) {
	unit, ok := obj.(model.Unit)
	if !ok {
		return &Response{}, api2go.NewHTTPError(errors.New(""), "", http.StatusBadRequest)
	}

	if _, err := r.orm.Id(unit.ID).Update(unit); err != nil {
		return &Response{}, err
	}
	return &Response{Code: http.StatusNoContent}, nil
}
