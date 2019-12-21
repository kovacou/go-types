// Copyright © 2019 Alexandre Kovac <contact@kovacou.com>.
//
// Use of this source code is governed by an MIT-style
// license that can be found in the LICENSE file.

package types

// Map is an hashmap.
type Map map[string]interface{}

// Reset the values of the map.
func (m *Map) Reset() {
	*m = Map{}
}

// Add a value to the map if the key doesn't exists.
func (m Map) Add(k string, v interface{}) {
	if _, ok := m[k]; !ok {
		m[k] = v
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
func (m Map) Find(matcher func(k string, v interface{}) bool) (string, interface{}, bool) {
	for k, v := range m {
		if matcher(k, v) {
			return k, v, true
		}
	}
	return "", nil, false
}

// FindAll elements matching the pattern.
func (m Map) FindAll(matcher func(k string, v interface{}) bool) Map {
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

// Len return the size of the map.
func (m Map) Len() int {
	return len(m)
}

// Set a new value in the map.
func (m Map) Set(k string, v interface{}) {
	m[k] = v
}

// Values return the list of values.
func (m Map) Values() []interface{} {
	out := make([]interface{}, 0, len(m))
	for k := range m {
		out = append(out, m[k])
	}
	return out
}
