// Copyright © 2020 Alexandre KOVAC <contact@kovacou.com>.
//
// Use of this source code is governed by an MIT-style
// license that can be found in the LICENSE file.

package types

import (
	"bytes"
	"time"
)

const (
	DateFormat     = "2006-01-02"
	DateTimeFormat = "2006-01-02 15:04:05"
)

// NewDate returns a new date from a time.Time.
func NewDate(t time.Time) Date {
	return Date{t}
}

// Date is a wrapper around time.Time.
type Date struct {
	time.Time
}

// MarshalJSON implements the json.Marshaler interface.
func (d Date) MarshalJSON() ([]byte, error) {
	d.Time.MarshalJSON()
	b := bytes.NewBuffer([]byte{})
	b.WriteRune('"')
	b.WriteString(d.Format(DateFormat))
	b.WriteRune('"')
	return b.Bytes(), nil
}

// NewDateTime create a new DateTime from a time.Time.
func NewDateTime(t time.Time) DateTime {
	return DateTime{t}
}

// DateTime is a wrapper around time.Time.
type DateTime struct {
	time.Time
}

// MarshalJSON implements the json.Marshaler interface.
func (d DateTime) MarshalJSON() ([]byte, error) {
	b := bytes.NewBuffer([]byte{})
	b.WriteRune('"')
	b.WriteString(d.Format(DateTimeFormat))
	b.WriteRune('"')
	return b.Bytes(), nil
}
