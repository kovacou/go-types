// Copyright © 2020 Alexandre KOVAC <contact@kovacou.com>.
//
// Use of this source code is governed by an MIT-style
// license that can be found in the LICENSE file.

package types

import (
	"bytes"
	"database/sql/driver"
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

// ParseDate returns the Date of the given string.
func ParseDate(t string) Date {
	dt, _ := time.Parse(DateFormat, t)
	return NewDate(dt)
}

// Date is a wrapper around time.Time.
type Date struct {
	time.Time
}

// MarshalJSON implements the json.Marshaler interface.
func (d Date) MarshalJSON() ([]byte, error) {
	b := bytes.NewBuffer([]byte{})
	b.WriteRune('"')
	b.WriteString(d.Format(DateFormat))
	b.WriteRune('"')
	return b.Bytes(), nil
}

// String returns the string representation of the date.
func (d Date) String() string {
	return d.Format(DateFormat)
}

// Value implements the database.Valuer interface.
func (d Date) Value() (driver.Value, error) {
	return d.String(), nil
}

// NewDateTime create a new DateTime from a time.Time.
func NewDateTime(t time.Time) DateTime {
	return DateTime{t}
}

// ParseDateTime returns the DateTime of the given string.
func ParseDateTime(t string) DateTime {
	dt, _ := time.Parse(DateTimeFormat, t)
	return NewDateTime(dt)
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

// String returns the string representation of the datetime.
func (d DateTime) String() string {
	return d.Format(DateTimeFormat)
}

// Value implements the database.Valuer interface.
func (d DateTime) Value() (driver.Value, error) {
	return d.String(), nil
}
