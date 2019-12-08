// Copyright © 2019 Alexandre Kovac <contact@kovacou.com>.
//
// Use of this source code is governed by an MIT-style
// license that can be found in the LICENSE file.

package types

import "sync"

// SyncMap return a new thread safe map.
func SyncMap() TSafeMap {
	m := TSafeMap{}
	m.Init(0)
	return m
}

// TSafeMap is a map thread safe.
type TSafeMap struct {
	mu     *sync.RWMutex
	values Map
}

// Init the map with the given buffer.
func (m *TSafeMap) Init(n int) {
	if m.values == nil {
		m.values = make(Map, n)
	}

	m.mu = &sync.RWMutex{}
}

// Map convert TSafeMap to Map.
func (m *TSafeMap) Map() (out Map) {
	m.mu.RLock()
	out = m.values.Copy()
	m.mu.RUnlock()
	return
}

// Set a new entry or change an entry for the given key "k".
func (m *TSafeMap) Set(k string, v interface{}) {
	m.mu.Lock()
	m.values[k] = v
	m.mu.Unlock()
}

// Add a new entry if the given key is not filled.
func (m *TSafeMap) Add(k string, v interface{}) {
	m.mu.Lock()
	if _, ok := m.values[k]; !ok {
		m.values[k] = v
	}
	m.mu.Unlock()
}

// Reset the values.
func (m *TSafeMap) Reset() {
	m.mu.Lock()
	m.values.Reset()
	m.mu.Unlock()
}
