package main

import (
	"github.com/go-xorm/core"
	"github.com/go-xorm/xorm"
	_ "github.com/mattn/go-sqlite3"
)

// Database ...
type Database struct {
	engine *xorm.Engine
}

// NewDatabase ...
func NewDatabase(production bool, name string) (*Database, error) {
	var (
		engine *xorm.Engine
		err    error
	)
	if production {
		panic("not implemented")
	} else {
		engine, err = xorm.NewEngine("sqlite3", name+".db")
	}
	if err != nil {
		return nil, err
	}

	engine.ShowSQL(true)
	engine.ShowExecTime(true)
	engine.SetMapper(core.GonicMapper{})

	if err := engine.Sync(new(Group)); err != nil {
		return nil, err
	}

	if err := engine.Sync(new(Unit)); err != nil {
		return nil, err
	}

	return &Database{engine}, nil
}

// Close ...
func (db *Database) Close() error {
	return db.engine.Close()
}

// AddUnit ...
func (db *Database) AddUnit(unit *Unit) error {
	if len(unit.ID) == 0 {
		unit.ID = newHashID()
	}
	if _, err := db.engine.Insert(unit); err != nil {
		return err
	}
	return nil
}

// GetUnit ...
func (db *Database) GetUnit(id string) (*Unit, error) {
	unit := Unit{ID: id}
	if found, err := db.engine.Get(&unit); !found || err != nil {
		return nil, err
	}
	return &unit, nil
}

// ListUnits ...
func (db Database) ListUnits() ([]Unit, error) {
	var units []Unit
	if err := db.engine.Find(&units); err != nil {
		return nil, err
	}
	return units, nil
}

// AddGroup ...
func (db *Database) AddGroup(group *Group) error {
	if len(group.ID) == 0 {
		group.ID = newHashID()
	}
	if _, err := db.engine.Insert(group); err != nil {
		return err
	}
	return nil
}

// GetGroup ...
func (db *Database) GetGroup(id string) (*Group, error) {
	group := Group{ID: id}
	if found, err := db.engine.Get(&group); !found || err != nil {
		return nil, err
	}
	return &group, nil
}

// ListGroups ...
func (db Database) ListGroups() ([]Group, error) {
	var groups []Group
	if err := db.engine.Find(&groups); err != nil {
		return nil, err
	}
	return groups, nil
}

// AddUnitToGroup ...
func (db Database) AddUnitToGroup(unit *Unit, group *Group) error {
	group.Units = append(group.Units, *unit)
	if _, err := db.engine.ID(group.ID).Update(group); err != nil {
		return err
	}
	return nil
}
