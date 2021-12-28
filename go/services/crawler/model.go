package crawler

import (
	"time"
)

type Seller struct {
	URL        string
	Key        string
	Name       *string
	Tel        *string
	Prefecture *string
	Address    *string
	Latitude   *float64
	Longitude  *float64
	OpenedAt   *time.Time
	Exported   *bool
	ExportedAt *time.Time
	Disabled   bool
}
