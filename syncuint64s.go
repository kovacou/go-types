package types

import (
	"sync"
)

type TSafeUint64s interface {
	// Reset the slice.
	Reset()

	// Contains say if "s" contains "values".
	Contains(...uint64) bool

	// ContainsOneOf says if "s" contains one of the "values".
	ContainsOneOf(...uint64) bool

	// Copy create a new copy of the slice.
	Copy() TSafeUint64s

	// Diff returns the difference between "s" and "s2".
	Diff(Uint64s) Uint64s

	// Empty says if the slice is empty.
	Empty() bool

	// Equal says if "s" and "s2" are equal.
	Equal(Uint64s) bool

	// Find the first element matching the pattern.
	Find(func(v uint64) bool) (uint64, bool)

	// FindAll elements matching the pattern.
	FindAll(func(v uint64) bool) Uint64s

	// First return the value of the first element.
	First() (uint64, bool)

	// Get the element "i" and say if it has been found.
	Get(int) (uint64, bool)

	// Intersect return the intersection between "s" and "s2".
	Intersect(Uint64s) Uint64s

	// Last return the value of the last element.
	Last() (uint64, bool)

	// Len returns the size of the slice.
	Len() int

	// Take n element and return a new slice.
	Take(int) Uint64s

	// S convert s into []interface{}
	S() []interface{}

	// S convert s into Uint64s
	Uint64s() Uint64s
}

func SyncUint64s() TSafeUint64s {
	return &tsafeUint64s{&sync.RWMutex{}, Uint64s{}}
}

type tsafeUint64s struct {
	mu     *sync.RWMutex
	values Uint64s
}

func (s *tsafeUint64s) Reset() {
	s.mu.Lock()
	s.values.Reset()
	s.mu.Unlock()
}

func (s *tsafeUint64s) Contains(values ...uint64) (ok bool) {
	s.mu.RLock()
	ok = s.values.Contains(values...)
	s.mu.RUnlock()
	return
}

func (s *tsafeUint64s) ContainsOneOf(values ...uint64) (ok bool) {
	s.mu.RLock()
	ok = s.values.ContainsOneOf(values...)
	s.mu.RUnlock()
	return
}

func (s *tsafeUint64s) Copy() TSafeUint64s {
	s2 := &tsafeUint64s{mu: &sync.RWMutex{}}
	s.mu.RLock()
	s2.values = s.values.Copy()
	s.mu.RUnlock()
	return s2
}

func (s *tsafeUint64s) Diff(s2 Uint64s) (out Uint64s) {
	s.mu.RLock()
	out = s.values.Diff(s2)
	s.mu.RUnlock()
	return
}

func (s *tsafeUint64s) Empty() (ok bool) {
	s.mu.RLock()
	ok = s.values.Empty()
	s.mu.RUnlock()
	return
}

func (s *tsafeUint64s) Equal(s2 Uint64s) (ok bool) {
	s.mu.RLock()
	ok = s.values.Equal(s2)
	s.mu.RUnlock()
	return
}

func (s *tsafeUint64s) Find(matcher func(v uint64) bool) (v uint64, ok bool) {
	s.mu.RLock()
	v, ok = s.values.Find(matcher)
	s.mu.RUnlock()
	return
}

func (s *tsafeUint64s) FindAll(matcher func(v uint64) bool) (v Uint64s) {
	s.mu.RLock()
	v = s.values.FindAll(matcher)
	s.mu.RUnlock()
	return
}

func (s *tsafeUint64s) First() (v uint64, ok bool) {
	s.mu.RLock()
	v, ok = s.values.First()
	s.mu.RUnlock()
	return
}

func (s *tsafeUint64s) Get(i int) (v uint64, ok bool) {
	s.mu.RLock()
	v, ok = s.values.Get(i)
	s.mu.RUnlock()
	return
}

func (s *tsafeUint64s) Intersect(s2 Uint64s) (v Uint64s) {
	s.mu.RLock()
	v = s.values.Intersect(s2)
	s.mu.RUnlock()
	return
}

func (s *tsafeUint64s) Last() (v uint64, ok bool) {
	s.mu.RLock()
	v, ok = s.values.Last()
	s.mu.RUnlock()
	return
}

func (s *tsafeUint64s) Len() (v int) {
	s.mu.RLock()
	v = s.values.Len()
	s.mu.RUnlock()
	return
}

func (s *tsafeUint64s) Take(n int) (v Uint64s) {
	s.mu.RLock()
	v = s.values.Take(n)
	s.mu.RUnlock()
	return
}

func (s *tsafeUint64s) S() (out []interface{}) {
	s.mu.RLock()
	out = s.values.S()
	s.mu.RUnlock()
	return
}

func (s *tsafeUint64s) Uint64s() (v Uint64s) {
	s.mu.RLock()
	v = s.values.Copy()
	s.mu.RUnlock()
	return
}
