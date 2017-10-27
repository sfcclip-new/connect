package model

import (
	"net/http"
	"strconv"
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

	UnitID     int64      `json:"unit_id" xorm:"notnull"`
	AccessType AccessType `json:"access_type" xorm:"notnull"`

	Host       string `json:"host"`
	RemoteAddr string `json:"remote_addr"`
	RequestURI string `json:"request_uri"`
	UserAgent  string `json:"user_agent"`
	Referer    string `json:"referer"`

	CreatedAt time.Time `json:"created_at" xorm:"created"`
}

// NewRecord creates a new record
func NewRecord(unitID int64, accessType AccessType, r *http.Request) *Record {
	record := new(Record)
	record.UnitID = unitID
	record.AccessType = accessType
	record.Host = r.Host
	if realIP := r.Header.Get("X-Real-IP"); len(realIP) != 0 {
		record.RemoteAddr = r.Header.Get("X-Real-IP")
	} else {
		record.RemoteAddr = r.RemoteAddr
	}
	record.RequestURI = r.RequestURI
	record.UserAgent = r.UserAgent()
	record.Referer = r.Referer()
	return record
}

// GetID to satisfy jsonapi.MarshalIdentifier interface
func (r Record) GetID() string {
	return strconv.FormatInt(r.ID, 10)
}
