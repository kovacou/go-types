// Copyright © 2019 Alexandre KOVAC <contact@kovacou.com>.
//
// Use of this source code is governed by an MIT-style
// license that can be found in the LICENSE file.

package types

import (
	"database/sql/driver"

	"github.com/twpayne/go-polyline"
)

// Polyline is the representation of a Google standard polyline.
type Polyline string

// Decode the polyline to Points.
func (p Polyline) Decode() (out Points) {
	coords, _, _ := polyline.DecodeCoords([]byte(p))
	out = make(Points, len(coords))

	for _, coord := range coords {
		out = append(out, Point{
			Lat: coord[0],
			Lng: coord[1],
		})
	}
	return out
}

// String convert the polyline to string.
func (p Polyline) String() string {
	return string(p)
}

// Value implements the database.Valuer interface.
func (p Polyline) Value() (driver.Value, error) {
	return p.String(), nil
}
