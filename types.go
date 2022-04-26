// Copyright © 2019 Alexandre Kovac <contact@kovacou.com>.
//
// Use of this source code is governed by an MIT-style
// license that can be found in the LICENSE file.

package types

import "time"

type Matcher func(k string, v any) bool

// Int returns the value of `v`.
func Int(v *int) int {
	if v != nil {
		return *v
	}
	return 0
}

// IntPtr return a pointer of `v`.
func IntPtr(v int) *int {
	return &v
}

// Uint returns the value of `v`.
func Uint(v *uint) uint {
	if v != nil {
		return *v
	}
	return 0
}

// UintPtr return a pointer of `v`.
func UintPtr(v uint) *uint {
	return &v
}

// Int64 return the value of `v`.
func Int64(v *int64) int64 {
	if v != nil {
		return *v
	}
	return 0
}

// Int64Ptr return a pointer of `v`.
func Int64Ptr(v int64) *int64 {
	return &v
}

// Uint64 return the value of `v`.
func Uint64(v *uint64) uint64 {
	if v != nil {
		return *v
	}
	return 0
}

// Uint64Ptr return a pointer of `v`.
func Uint64Ptr(v uint64) *uint64 {
	return &v
}

// Float64 return the value of `v`.
func Float64(v *float64) float64 {
	if v != nil {
		return *v
	}
	return 0
}

// Float64Ptr return a pointer of `v`.
func Float64Ptr(v float64) *float64 {
	return &v
}

// String returns the value of `v`.
func String(v *string) string {
	if v != nil {
		return *v
	}
	return ""
}

// StringPtr return the value of `v`.
func StringPtr(v string) *string {
	return &v
}

// Bool return the value of `v`.
func Bool(v *bool) bool {
	if v != nil {
		return *v
	}
	return false
}

// BoolPtr return a pointer of `v`.
func BoolPtr(v bool) *bool {
	return &v
}

// Time return the value of `v`.
func Time(v *time.Time) time.Time {
	if v != nil {
		return *v
	}
	return time.Time{}
}

// TimePtr return a pointer of `v`.
func TimePtr(v time.Time) *time.Time {
	return &v
}
