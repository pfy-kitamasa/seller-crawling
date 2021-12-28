// Code generated by sqlc. DO NOT EDIT.

package repository

import (
	"database/sql"
	"time"

	"github.com/tabbed/pqtype"
)

type Domain struct {
	ID             int32
	Key            string
	Url            string
	AllowDomains   []string
	LatestVisitUrl string
	CreatedAt      time.Time
	UpdatedAt      time.Time
}

type Seller struct {
	ID         int32
	Url        string
	Key        string
	Name       string
	Tel        string
	Prefecture string
	Address    string
	Longitude  float64
	Latitude   float64
	OpenedAt   time.Time
	Meta       pqtype.NullRawMessage
	Exported   bool
	ExportedAt sql.NullTime
	Disabled   bool
	CreatedAt  time.Time
	UpdatedAt  time.Time
}
