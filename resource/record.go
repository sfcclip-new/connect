package resource

import (
	"strconv"

	"github.com/akkyie/connect.sfcclip.net/model"
	"github.com/go-xorm/xorm"
	"github.com/manyminds/api2go"
)

// RecordResource provides routing for records
type RecordResource struct {
	orm *xorm.Engine
}

// NewRecordResource returns new RecordResource
func NewRecordResource(orm *xorm.Engine) *RecordResource {
	return &RecordResource{orm}
}

// FindAll to satisfy api2go data source interface
func (r RecordResource) FindAll(req api2go.Request) (api2go.Responder, error) {
	var records []model.Record
	if err := r.orm.Find(&records); err != nil {
		return &Response{}, err
	}

	return &Response{Res: records}, nil
}

// FindOne to satisfy `api2go.DataSource` interface
// this method should return the user with the given ID, otherwise an error
func (r RecordResource) FindOne(ID string, req api2go.Request) (api2go.Responder, error) {
	id, err := strconv.ParseInt(ID, 10, 0)
	if err != nil {
		return &Response{}, err
	}
	record := model.Record{ID: id}
	if _, err := r.orm.Get(&record); err != nil {
		return &Response{}, err
	}
	return &Response{Res: record}, nil
}
