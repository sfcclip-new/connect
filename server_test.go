package main

import (
	"context"
	"fmt"
	"log"
	"testing"
	"time"
)

const (
	testGroupID = "abcde"
	testUnitID  = "12345"
)

type TestDatabase struct {
}

func (db *TestDatabase) AddUnit(unit *Unit) error {
	return nil
}

func (db *TestDatabase) GetGroup(id string) (*Group, error) {

	if id != testGroupID {
		return nil, fmt.Errorf("not found")
	}
	return &Group{ID: testGroupID}, nil
}

func (db *TestDatabase) AddGroup(group *Group) error {
	return nil
}

func (db *TestDatabase) ListGroups() ([]Group, error) {
	return []Group{Group{ID: testGroupID}}, nil
}

func (db *TestDatabase) GetUnit(id string) (*Unit, error) {
	if id != testUnitID {
		return nil, fmt.Errorf("not found")
	}
	return &Unit{ID: testUnitID}, nil
}

func (db *TestDatabase) ListUnits() ([]Unit, error) {
	return []Unit{Unit{ID: testUnitID}}, nil
}

func (db *TestDatabase) AddUnitToGroup(unit *Unit, group *Group) error {
	return nil
}

func TestNewServer(t *testing.T) {
	db := &TestDatabase{}
	server := NewServer("8080", db, false)
	if server == nil {
		t.Fatalf("Server could not initialized")
	}
}

func TestServerStartAndShutdown(t *testing.T) {
	db := &TestDatabase{}
	server := NewServer(":8080", db, false)

	go func() {
		if err := server.Start(); err != nil {
			t.Fatal(err)
		}
	}()

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	if err := server.Shutdown(ctx); err != nil {
		log.Fatal(err)
	}
}
