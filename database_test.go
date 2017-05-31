package main

import (
	"os"
	"testing"
)

func TestNewDatabaseAndClose(t *testing.T) {
	name := newHashID()
	db, err := NewDatabase(false, name)
	if err != nil {
		t.Fatal(err)
	}
	if db == nil {
		t.Fatalf("Database could not be initialized")
	}

	if db.engine.DriverName() != "sqlite3" {
		t.Fatal("Invalid driver name")
	}

	if err = db.Close(); err != nil {
		t.Fatal(err)
	}

	os.Remove(name + ".db")
}

func TestUnit(t *testing.T) {
	// Setup
	name := newHashID()
	db, _ := NewDatabase(false, name)
	defer db.Close()
	defer os.Remove(name + ".db")

	testUnitA := new(Unit)
	testUnitB := new(Unit)

	for _, testUnit := range []*Unit{testUnitA, testUnitB} {
		db.AddUnit(testUnit)

		// Test GetUnit
		unit, err := db.GetUnit(testUnit.ID)
		if err != nil {
			t.Fatal(err)
			return
		}

		if unit == nil || unit.ID != testUnit.ID {
			t.Fatalf("Invalid unit")
		}
	}

	// Test GetUnit for an invalid id
	if unit, _ := db.GetUnit(newHashID()); unit != nil {
		t.Fatalf("Invalid unit")
		return
	}

	// Test ListUnits
	units, err := db.ListUnits()
	if err != nil {
		t.Fatal(err)
		return
	}

	if len := len(units); len != 2 {
		t.Fatalf("Invalid number of units: %v", units)
	}
}

func TestGroup(t *testing.T) {
	// Setup
	name := newHashID()
	db, _ := NewDatabase(false, name)
	defer db.Close()
	defer os.Remove(name + ".db")

	testUnitA := new(Unit)
	db.AddUnit(testUnitA)
	testUnitB := new(Unit)
	db.AddUnit(testUnitB)

	// Test AddGroup
	testGroup := new(Group)
	db.AddGroup(testGroup)

	// Test GetGroup
	group, err := db.GetGroup(testGroup.ID)
	if err != nil {
		t.Fatal(err)
		return
	}

	if group == nil || group.ID != testGroup.ID {
		t.Fatalf("Invalid group")
	}

	// Test AddUnitToGroup; confirm groups have no unit at their initial state

	if len := len(group.Units); len != 0 {
		t.Fatalf("Invalid number of units in a group: %v", group)
	}

	// Test AddUnitToGroup
	if err = db.AddUnitToGroup(testUnitA, group); err != nil {
		t.Fatal(err)
	}

	if err = db.AddUnitToGroup(testUnitB, group); err != nil {
		t.Fatal(err)
	}

	// Confirm now the group has two units
	group, _ = db.GetGroup(testGroup.ID)
	if len := len(group.Units); len != 2 {
		t.Fatalf("Invalid number of units in a group: %v", group)
	}

	// Test ListGroups
	groups, err := db.ListGroups()
	if err != nil {
		t.Fatal(err)
		return
	}

	if len := len(groups); len != 1 {
		t.Fatalf("Invalid number of groups: %v", groups)
	}
}
