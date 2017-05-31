package main

import (
	"math"
	"math/rand"
	"time"

	hashids "github.com/speps/go-hashids"
)

const hashidsSalt = "arz and clipchan"

// An utility function to generate a new hashid
func newHashID() string {
	data := hashids.NewData()
	data.Salt = hashidsSalt
	hashids, _ := hashids.NewWithData(data)
	rand.Seed(time.Now().UnixNano())
	seed := rand.Int63n(math.MaxInt16)
	hashid, _ := hashids.EncodeInt64([]int64{seed})
	return hashid
}

// Group representes a group of multiple `Ad`.
type Group struct {
	ID        string    `xorm:"pk index"`
	CreatedAt time.Time `xorm:"created"`
	UpdatedAt time.Time `xorm:"updated"`

	Units []Unit
}

// Unit representes an advertising information.
type Unit struct {
	ID        string    `xorm:"pk index"`
	CreatedAt time.Time `xorm:"created"`
	UpdatedAt time.Time `xorm:"updated"`

	Advertizer string
	TargetURL  string
	ImageURL   string
}
