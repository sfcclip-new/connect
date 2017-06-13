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

// GroupResource provides routing for groups
type GroupResource struct {
	orm *xorm.Engine
}

// NewGroupResource returns new GroupResource
func NewGroupResource(orm *xorm.Engine) *GroupResource {
	return &GroupResource{orm}
}

// FindAll to satisfy api2go data source interface
func (r GroupResource) FindAll(req api2go.Request) (api2go.Responder, error) {
	var groups []model.Group
	if err := r.orm.Find(&groups); err != nil {
		return &Response{}, err
	}

	log.Info("FindAll", groups)
	for key, group := range groups {
		if len(group.UnitIDs) == 0 {
			continue
		}
		var units []model.Unit
		if err := r.orm.In("id", group.UnitIDs).Find(&units); err != nil {
			return &Response{}, err
		}
		groups[key].Units = units
	}

	return &Response{Res: groups}, nil
}

// FindOne to satisfy `api2go.DataSource` interface
// this method should return the user with the given ID, otherwise an error
func (r GroupResource) FindOne(ID string, req api2go.Request) (api2go.Responder, error) {
	id, err := strconv.ParseInt(ID, 10, 0)
	if err != nil {
		return &Response{}, err
	}

	group := model.Group{ID: id}
	if _, err := r.orm.Get(&group); err != nil {
		return &Response{}, err
	}

	var units []model.Unit
	if err := r.orm.In("id", group.UnitIDs).Find(&units); err != nil {
		return &Response{}, err
	}
	group.Units = units

	return &Response{Res: group}, nil
}

// Create method to satisfy `api2go.DataSource` interface
func (r GroupResource) Create(obj interface{}, req api2go.Request) (api2go.Responder, error) {
	group, ok := obj.(model.Group)
	if !ok {
		return &Response{}, api2go.NewHTTPError(errors.New(""), "", http.StatusBadRequest)
	}
	if _, err := r.orm.Insert(&group); err != nil {
		return &Response{}, err
	}
	log.Info(group)
	return &Response{Res: group, Code: http.StatusCreated}, nil
}

// Delete to satisfy `api2go.DataSource` interface
func (r GroupResource) Delete(id string, req api2go.Request) (api2go.Responder, error) {
	ID, err := strconv.ParseInt(id, 10, 0)
	if err != nil {
		return &Response{}, err
	}
	if _, err := r.orm.Id(ID).Delete(&model.Group{}); err != nil {
		return &Response{}, err
	}
	return &Response{Code: http.StatusNoContent}, nil
}

// Update stores all changes on the user
func (r GroupResource) Update(obj interface{}, req api2go.Request) (api2go.Responder, error) {
	group, ok := obj.(model.Group)
	if !ok {
		return &Response{}, api2go.NewHTTPError(errors.New(""), "", http.StatusBadRequest)
	}

	if _, err := r.orm.Id(group.ID).Update(group); err != nil {
		return &Response{}, err
	}
	return &Response{Code: http.StatusNoContent}, nil
}
