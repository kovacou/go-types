package types

import "sync"

type TSafeUints interface {
	// Reset the slice.
	Reset()

	// Contains say if "s" contains "values".
	Contains(...uint) bool

	// ContainsOneOf says if "s" contains one of the "values".
	ContainsOneOf(...uint) bool

	// Copy create a new copy of the slice.
	Copy() TSafeUints

	// Diff returns the difference between "s" and "s2".
	Diff(Uints) Uints

	// Empty says if the slice is empty.
	Empty() bool

	// Equal says if "s" and "s2" are equal.
	Equal(Uints) bool

	// Find the first element matching the pattern.
	Find(func(v uint) bool) (uint, bool)

	// FindAll elements matching the pattern.
	FindAll(func(v uint) bool) Uints

	// First return the value of the first element.
	First() (uint, bool)

	// Get the element "i" and say if it has been found.
	Get(int) (uint, bool)

	// Intersect return the intersection between "s" and "s2".
	Intersect(Uints) Uints

	// Last return the value of the last element.
	Last() (uint, bool)

	// Len returns the size of the slice.
	Len() int

	// Take n element and return a new slice.
	Take(int) Uints

	// S convert s into []any
	S() []any

	// S convert s into Uints
	Uints() Uints
}

func SyncUints() TSafeUints {
	return &tsafeUints{&sync.RWMutex{}, Uints{}}
}

type tsafeUints struct {
	mu     *sync.RWMutex
	values Uints
}

func (s *tsafeUints) Reset() {
	s.mu.Lock()
	s.values.Reset()
	s.mu.Unlock()
}

func (s *tsafeUints) Contains(values ...uint) (ok bool) {
	s.mu.RLock()
	ok = s.values.Contains(values...)
	s.mu.RUnlock()
	return
}

func (s *tsafeUints) ContainsOneOf(values ...uint) (ok bool) {
	s.mu.RLock()
	ok = s.values.ContainsOneOf(values...)
	s.mu.RUnlock()
	return
}

func (s *tsafeUints) Copy() TSafeUints {
	s2 := &tsafeUints{mu: &sync.RWMutex{}}
	s.mu.RLock()
	s2.values = s.values.Copy()
	s.mu.RUnlock()
	return s2
}

func (s *tsafeUints) Diff(s2 Uints) (out Uints) {
	s.mu.RLock()
	out = s.values.Diff(s2)
	s.mu.RUnlock()
	return
}

func (s *tsafeUints) Empty() (ok bool) {
	s.mu.RLock()
	ok = s.values.Empty()
	s.mu.RUnlock()
	return
}

func (s *tsafeUints) Equal(s2 Uints) (ok bool) {
	s.mu.RLock()
	ok = s.values.Equal(s2)
	s.mu.RUnlock()
	return
}

func (s *tsafeUints) Find(matcher func(v uint) bool) (v uint, ok bool) {
	s.mu.RLock()
	v, ok = s.values.Find(matcher)
	s.mu.RUnlock()
	return
}

func (s *tsafeUints) FindAll(matcher func(v uint) bool) (v Uints) {
	s.mu.RLock()
	v = s.values.FindAll(matcher)
	s.mu.RUnlock()
	return
}

func (s *tsafeUints) First() (v uint, ok bool) {
	s.mu.RLock()
	v, ok = s.values.First()
	s.mu.RUnlock()
	return
}

func (s *tsafeUints) Get(i int) (v uint, ok bool) {
	s.mu.RLock()
	v, ok = s.values.Get(i)
	s.mu.RUnlock()
	return
}

func (s *tsafeUints) Intersect(s2 Uints) (v Uints) {
	s.mu.RLock()
	v = s.values.Intersect(s2)
	s.mu.RUnlock()
	return
}

func (s *tsafeUints) Last() (v uint, ok bool) {
	s.mu.RLock()
	v, ok = s.values.Last()
	s.mu.RUnlock()
	return
}

func (s *tsafeUints) Len() (v int) {
	s.mu.RLock()
	v = s.values.Len()
	s.mu.RUnlock()
	return
}

func (s *tsafeUints) Take(n int) (v Uints) {
	s.mu.RLock()
	v = s.values.Take(n)
	s.mu.RUnlock()
	return
}

func (s *tsafeUints) S() (out []any) {
	s.mu.RLock()
	out = s.values.S()
	s.mu.RUnlock()
	return
}

func (s *tsafeUints) Uints() (v Uints) {
	s.mu.RLock()
	v = s.values.Copy()
	s.mu.RUnlock()
	return
}
