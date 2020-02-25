// Copyright © 2019 Alexandre Kovac <contact@kovacou.com>.
//
// Use of this source code is governed by an MIT-style
// license that can be found in the LICENSE file.

package types

import "sync"

// TSafeMap abstract the implementation of SyncMap.
type TSafeMap interface {
	Add(string, interface{})
	Find(Matcher) (string, interface{}, bool)
	FindAll(Matcher) Map
	Get(string) (interface{}, bool)
	Map() Map
	Set(string, interface{})
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

// Add a new entry if the given key is not filled.
func (m *tsafeMap) Add(k string, v interface{}) {
	m.mu.Lock()
	if _, ok := m.values[k]; !ok {
		m.values[k] = v
	}
	m.mu.Unlock()
}

// Find the first element matching the pattern.
func (m *tsafeMap) Find(matcher Matcher) (k string, v interface{}, ok bool) {
	m.mu.RLock()
	k, v, ok = m.values.Find(matcher)
	m.mu.RUnlock()
	return
}

// FindAll elements matching the pattern.
func (m *tsafeMap) FindAll(matcher Matcher) (out Map) {
	m.mu.RLock()
	out = m.values.FindAll(matcher)
	m.mu.RUnlock()
	return
}

// Get an element from the key.
func (m *tsafeMap) Get(k string) (v interface{}, ok bool) {
	m.mu.RLock()
	v, ok = m.values[k]
	m.mu.RUnlock()
	return
}

// Map convert TSafeMap to Map.
func (m *tsafeMap) Map() (out Map) {
	m.mu.RLock()
	out = m.values.Copy()
	m.mu.RUnlock()
	return
}

// Set a new entry or change an entry for the given key "k".
func (m *tsafeMap) Set(k string, v interface{}) {
	m.mu.Lock()
	m.values[k] = v
	m.mu.Unlock()
}

// Reset the values.
func (m *tsafeMap) Reset() {
	m.mu.Lock()
	m.values.Reset()
	m.mu.Unlock()
}
