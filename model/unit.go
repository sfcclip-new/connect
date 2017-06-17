package model

import (
	"strconv"
	"time"
)

// Unit represents an ad unit
type Unit struct {
	ID int64 `json:"-" xorm:"pk autoincr"`

	Name       string `json:"name"`
	Advertizer string `json:"advertizer"`
	ImageURL   string `json:"imageURL"`
	TargetURL  string `json:"targetURL"`

	ImageCount int64 `json:"image_count"`
	OpenCount  int64 `json:"open_count"`

	CreatedAt time.Time `json:"created_at" xorm:"created"`
	UpdatedAt time.Time `json:"updated_at" xorm:"updated"`
}

// GetID to satisfy jsonapi.MarshalIdentifier interface
func (u Unit) GetID() string {
	return strconv.FormatInt(u.ID, 10)
}

// SetID to satisfy jsonapi.UnmarshalIdentifier interface
func (u *Unit) SetID(id string) error {
	u.ID, _ = strconv.ParseInt(id, 10, 64)
	return nil
}
