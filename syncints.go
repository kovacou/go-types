package types

import "sync"

type TSafeInts interface {
	// Reset the slice.
	Reset()

	// Contains say if "s" contains "values".
	Contains(...int) bool

	// ContainsOneOf says if "s" contains one of the "values".
	ContainsOneOf(...int) bool

	// Copy create a new copy of the slice.
	Copy() TSafeInts

	// Diff returns the difference between "s" and "s2".
	Diff(Ints) Ints

	// Empty says if the slice is empty.
	Empty() bool

	// Equal says if "s" and "s2" are equal.
	Equal(Ints) bool

	// Find the first element matching the pattern.
	Find(func(v int) bool) (int, bool)

	// FindAll elements matching the pattern.
	FindAll(func(v int) bool) Ints

	// First return the value of the first element.
	First() (int, bool)

	// Get the element "i" and say if it has been found.
	Get(int) (int, bool)

	// Intersect return the intersection between "s" and "s2".
	Intersect(Ints) Ints

	// Last return the value of the last element.
	Last() (int, bool)

	// Len returns the size of the slice.
	Len() int

	// Take n element and return a new slice.
	Take(int) Ints

	// S convert s into []any
	S() []any

	// Ints convert s into Ints
	Ints() Ints
}

func SyncInts() TSafeInts {
	return &tsafeInts{&sync.RWMutex{}, Ints{}}
}

type tsafeInts struct {
	mu     *sync.RWMutex
	values Ints
}

func (s *tsafeInts) Reset() {
	s.mu.Lock()
	s.values.Reset()
	s.mu.Unlock()
}

func (s *tsafeInts) Contains(values ...int) (ok bool) {
	s.mu.RLock()
	ok = s.values.Contains(values...)
	s.mu.RUnlock()
	return
}

func (s *tsafeInts) ContainsOneOf(values ...int) (ok bool) {
	s.mu.RLock()
	ok = s.values.ContainsOneOf(values...)
	s.mu.RUnlock()
	return
}

func (s *tsafeInts) Copy() TSafeInts {
	s2 := &tsafeInts{mu: &sync.RWMutex{}}
	s.mu.RLock()
	s2.values = s.values.Copy()
	s.mu.RUnlock()
	return s2
}

func (s *tsafeInts) Diff(s2 Ints) (out Ints) {
	s.mu.RLock()
	out = s.values.Diff(s2)
	s.mu.RUnlock()
	return
}

func (s *tsafeInts) Empty() (ok bool) {
	s.mu.RLock()
	ok = s.values.Empty()
	s.mu.RUnlock()
	return
}

func (s *tsafeInts) Equal(s2 Ints) (ok bool) {
	s.mu.RLock()
	ok = s.values.Equal(s2)
	s.mu.RUnlock()
	return
}

func (s *tsafeInts) Find(matcher func(v int) bool) (v int, ok bool) {
	s.mu.RLock()
	v, ok = s.values.Find(matcher)
	s.mu.RUnlock()
	return
}

func (s *tsafeInts) FindAll(matcher func(v int) bool) (v Ints) {
	s.mu.RLock()
	v = s.values.FindAll(matcher)
	s.mu.RUnlock()
	return
}

func (s *tsafeInts) First() (v int, ok bool) {
	s.mu.RLock()
	v, ok = s.values.First()
	s.mu.RUnlock()
	return
}

func (s *tsafeInts) Get(i int) (v int, ok bool) {
	s.mu.RLock()
	v, ok = s.values.Get(i)
	s.mu.RUnlock()
	return
}

func (s *tsafeInts) Intersect(s2 Ints) (v Ints) {
	s.mu.RLock()
	v = s.values.Intersect(s2)
	s.mu.RUnlock()
	return
}

func (s *tsafeInts) Last() (v int, ok bool) {
	s.mu.RLock()
	v, ok = s.values.Last()
	s.mu.RUnlock()
	return
}

func (s *tsafeInts) Len() (v int) {
	s.mu.RLock()
	v = s.values.Len()
	s.mu.RUnlock()
	return
}

func (s *tsafeInts) Take(n int) (v Ints) {
	s.mu.RLock()
	v = s.values.Take(n)
	s.mu.RUnlock()
	return
}

func (s *tsafeInts) S() (out []any) {
	s.mu.RLock()
	out = s.values.S()
	s.mu.RUnlock()
	return
}

func (s *tsafeInts) Ints() (v Ints) {
	s.mu.RLock()
	v = s.values.Copy()
	s.mu.RUnlock()
	return
}
