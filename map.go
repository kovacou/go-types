// Copyright © 2019 Alexandre KOVAC <contact@kovacou.com>.
//
// Use of this source code is governed by an MIT-style
// license that can be found in the LICENSE file.

package types

import (
	"time"
)

// Map is a hashmap.
type Map map[string]any

// Reset the values of the map.
func (m *Map) Reset() {
	*m = Map{}
}

// Add a value to the map if the key doesn't exists.
func (m Map) Add(k string, v any) {
	if _, ok := m[k]; !ok {
		m[k] = v
	}
}

// Merge another map.
func (m Map) Merge(sub map[string]any) {
	for k, v := range sub {
		m.Add(k, v)
	}
}

// Copy the keys and values into a new map.
func (m Map) Copy() Map {
	out := Map{}
	for k, v := range m {
		out[k] = v
	}
	return out
}

// Find the first element matching the pattern.
func (m Map) Find(matcher Matcher) (string, any, bool) {
	for k, v := range m {
		if matcher(k, v) {
			return k, v, true
		}
	}
	return "", nil, false
}

// FindAll elements matching the pattern.
func (m Map) FindAll(matcher Matcher) Map {
	out := Map{}
	for k, v := range m {
		if matcher(k, v) {
			out[k] = v
		}
	}
	return out
}

// KeyExists says if the list of keys exists.
func (m Map) KeyExists(keys ...string) bool {
	for _, k := range keys {
		if _, ok := m[k]; !ok {
			return false
		}
	}
	return true
}

// Keys return the list of keys.
func (m Map) Keys() []string {
	out := make([]string, 0, len(m))
	for k := range m {
		out = append(out, k)
	}
	return out
}

// Len returns the size of the map.
func (m Map) Len() int {
	return len(m)
}

// Set a new value in the map.
func (m Map) Set(k string, v any) {
	m[k] = v
}

// Values return the list of values.
func (m Map) Values() []any {
	out := make([]any, 0, len(m))
	for k := range m {
		out = append(out, m[k])
	}
	return out
}

// Get an element from the map.
func (m Map) Get(k string) (v any, ok bool) {
	v, ok = m[k]
	return
}

// String get an element from the map as string.
func (m Map) String(k string) string {
	return m[k].(string)
}

func (m Map) StringPtr(k string) *string {
	return StringPtr(m.String(k))
}

func (m Map) Int(k string) int {
	return m[k].(int)
}

func (m Map) IntPtr(k string) *int {
	return IntPtr(m.Int(k))
}

func (m Map) Int64(k string) int64 {
	return m[k].(int64)
}

func (m Map) Int64Ptr(k string) *int64 {
	return Int64Ptr(m.Int64(k))
}

func (m Map) Uint64(k string) uint64 {
	return uint64(m.Int64(k))
}

func (m Map) Uint64Ptr(k string) *uint64 {
	return Uint64Ptr(m.Uint64(k))
}

func (m Map) Time(k string) time.Time {
	return m[k].(time.Time)
}

func (m Map) TimePtr(k string) *time.Time {
	return TimePtr(m.Time(k))
}
