package model

import (
	"strconv"
	"time"

	"github.com/manyminds/api2go/jsonapi"
)

// Group represents an group of multiple units
type Group struct {
	ID   int64  `json:"-" xorm:"pk autoincr"`
	Name string `json:"name"`

	Units   []Unit  `json:"-" xorm:"-"`
	UnitIDs []int64 `json:"-" xorm:"'unit_ids'"`

	CreatedAt time.Time `json:"created_at" xorm:"created"`
	UpdatedAt time.Time `json:"updated_at" xorm:"updated"`
}

// GetID to satisfy jsonapi.MarshalIdentifier interface
func (g Group) GetID() string {
	return strconv.FormatInt(g.ID, 10)
}

// SetID to satisfy jsonapi.UnmarshalIdentifier interface
func (g *Group) SetID(id string) error {
	g.ID, _ = strconv.ParseInt(id, 10, 64)
	return nil
}

// GetReferences to satisfy the jsonapi.MarshalReferences interface
func (g Group) GetReferences() []jsonapi.Reference {
	return []jsonapi.Reference{
		{
			Type: "units",
			Name: "units",
		},
	}
}

// GetReferencedIDs to satisfy the jsonapi.MarshalLinkedRelations interface
func (g Group) GetReferencedIDs() []jsonapi.ReferenceID {
	results := []jsonapi.ReferenceID{}
	for _, unitID := range g.UnitIDs {
		results = append(results, jsonapi.ReferenceID{
			ID:   strconv.FormatInt(unitID, 10),
			Type: "units",
			Name: "units",
		})
	}
	return results
}

// GetReferencedStructs to satisfy the jsonapi.MarhsalIncludedRelations interface
func (g Group) GetReferencedStructs() []jsonapi.MarshalIdentifier {
	result := []jsonapi.MarshalIdentifier{}
	for key := range g.Units {
		result = append(result, g.Units[key])
	}
	return result
}

// SetToOneReferenceID to satisfy jsonapi.UnmarshalToOneRelations interface
func (g *Group) SetToOneReferenceID(name, ID string) error {
	if name == "units" {
		id, err := strconv.ParseInt(ID, 10, 64)
		if err != nil {
			return err
		}
		g.Units = append([]Unit{}, Unit{ID: id})
		g.UnitIDs = append([]int64{}, id)
	}
	return nil
}

// SetToManyReferenceIDs to satisfy jsonapi.UnmarshalToManyRelations interface
func (g *Group) SetToManyReferenceIDs(name string, IDs []string) error {
	if name == "units" {
		g.Units = []Unit{}
		g.UnitIDs = []int64{}
		for _, ID := range IDs {
			id, err := strconv.ParseInt(ID, 10, 64)
			if err != nil {
				return err
			}
			g.Units = append(g.Units, Unit{ID: id})
			g.UnitIDs = append(g.UnitIDs, id)
		}
	}
	return nil
}
