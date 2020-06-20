// Copyright © 2019 Alexandre Kovac <contact@kovacou.com>.
//
// Use of this source code is governed by an MIT-style
// license that can be found in the LICENSE file.

package types

import "sync"

// TSafeMap abstract the implementation of SyncMap.
type TSafeMap interface {
	// Add a new entry if the given key is not filled.
	Add(string, interface{})

	// Find the first element matching the pattern.
	Find(Matcher) (string, interface{}, bool)

	// FindAll elements matching the pattern.
	FindAll(Matcher) Map

	// Get an element from the key.
	Get(string) (interface{}, bool)

	// Map convert TSafeMap to Map.
	Map() Map

	// Set a new entry or change an entry for the given key "k".
	Set(string, interface{})

	// Reset the values.
	Reset()
}

// SyncMap return a new ThreadSafeMap.
func SyncMap() TSafeMap {
	return &tsafeMap{
		&sync.RWMutex{},
		make(Map, 0),
	}
}

// TSafeMap is a map thread safe.
type tsafeMap struct {
	mu     *sync.RWMutex
	values Map
}

func (m *tsafeMap) Add(k string, v interface{}) {
	m.mu.Lock()
	if _, ok := m.values[k]; !ok {
		m.values[k] = v
	}
	m.mu.Unlock()
}

func (m *tsafeMap) Find(matcher Matcher) (k string, v interface{}, ok bool) {
	m.mu.RLock()
	k, v, ok = m.values.Find(matcher)
	m.mu.RUnlock()
	return
}

func (m *tsafeMap) FindAll(matcher Matcher) (out Map) {
	m.mu.RLock()
	out = m.values.FindAll(matcher)
	m.mu.RUnlock()
	return
}

func (m *tsafeMap) Get(k string) (v interface{}, ok bool) {
	m.mu.RLock()
	v, ok = m.values[k]
	m.mu.RUnlock()
	return
}

func (m *tsafeMap) Map() (out Map) {
	m.mu.RLock()
	out = m.values.Copy()
	m.mu.RUnlock()
	return
}

func (m *tsafeMap) Set(k string, v interface{}) {
	m.mu.Lock()
	m.values[k] = v
	m.mu.Unlock()
}

func (m *tsafeMap) Reset() {
	m.mu.Lock()
	m.values.Reset()
	m.mu.Unlock()
}
