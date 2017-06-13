package model

import (
	"net/http"
	"time"
)

// AccessType represents the type of a record
type AccessType int

const (
	// ImageAccessType recorded when image URL get accessed
	ImageAccessType AccessType = iota

	// OpenAccessType recorded when target URL get accessed
	OpenAccessType
)

// Record represents an access
type Record struct {
	ID int64 `json:"-" xorm:"pk autoincr"`

	UnitID     int64      `xorm:"notnull"`
	AccessType AccessType `xorm:"notnull"`

	Host       string
	RemoteAddr string
	RequestURI string
	UserAgent  string
	Referer    string

	CreatedAt time.Time `json:"created_at" xorm:"created"`
}

// NewRecord creates a new record
func NewRecord(unitID int64, accessType AccessType, r *http.Request) *Record {
	record := new(Record)
	record.UnitID = unitID
	record.AccessType = accessType
	record.Host = r.Host
	record.RemoteAddr = r.RemoteAddr
	record.RequestURI = r.RequestURI
	record.UserAgent = r.UserAgent()
	record.Referer = r.Referer()
	return record
}
