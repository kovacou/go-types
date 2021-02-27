// Copyright © 2019 Alexandre KOVAC <contact@kovacou.com>.
//
// Use of this source code is governed by an MIT-style
// license that can be found in the LICENSE file.

package types

import "github.com/twpayne/go-polyline"

// Point is the representation of a couple of latitude and longitude.
type Point struct {
	Lat float64 `json:"lat"`
	Lng float64 `json:"lng"`
}

// Encode the point to a slice of coordinates.
func (p Point) Encode() []float64 {
	return []float64{
		p.Lat,
		p.Lng,
	}
}

// Points is a slice of Point.
type Points []Point

// Polyline encode the series of points into string polyline.
func (p Points) Polyline() Polyline {
	return Polyline(polyline.EncodeCoords(p.Encode()))
}

// Encode the series of points.
func (p Points) Encode() (out [][]float64) {
	out = make([][]float64, len(p))
	for _, pt := range p {
		out = append(out, pt.Encode())
	}
	return
}
