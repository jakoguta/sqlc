// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0

package querytest

import (
	"database/sql"
	"time"
)

type Measurement struct {
	CityID    int32
	Logdate   time.Time
	Peaktemp  sql.NullInt32
	Unitsales sql.NullInt32
}

type MeasurementY2006m02 struct {
	CityID    int32
	Logdate   time.Time
	Peaktemp  sql.NullInt32
	Unitsales sql.NullInt32
}
