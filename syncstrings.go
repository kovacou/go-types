package types

import "sync"

type TSafeStrings interface {
	// Reset the slice.
	Reset()

	// Contains say if "s" contains "values".
	Contains(...string) bool

	// ContainsOneOf says if "s" contains one of the "values".
	ContainsOneOf(...string) bool

	// Copy create a new copy of the slice.
	Copy() TSafeStrings

	// Diff returns the difference between "s" and "s2".
	Diff(Strings) Strings

	// Empty says if the slice is empty.
	Empty() bool

	// Equal says if "s" and "s2" are equal.
	Equal(Strings) bool

	// Find the first element matching the pattern.
	Find(func(v string) bool) (string, bool)

	// FindAll elements matching the pattern.
	FindAll(func(v string) bool) Strings

	// First return the value of the first element.
	First() (string, bool)

	// Get the element "i" and say if it has been found.
	Get(int) (string, bool)

	// Intersect return the intersection between "s" and "s2".
	Intersect(Strings) Strings

	// Last return the value of the last element.
	Last() (string, bool)

	// Len returns the size of the slice.
	Len() int

	// Take n element and return a new slice.
	Take(int) Strings

	// S convert s into []any
	S() []any

	// S convert s into Strings
	Strings() Strings
}

func SyncStrings() TSafeStrings {
	return &tsafeStrings{&sync.RWMutex{}, Strings{}}
}

type tsafeStrings struct {
	mu     *sync.RWMutex
	values Strings
}

func (s *tsafeStrings) Reset() {
	s.mu.Lock()
	s.values.Reset()
	s.mu.Unlock()
}

func (s *tsafeStrings) Contains(values ...string) (ok bool) {
	s.mu.RLock()
	ok = s.values.Contains(values...)
	s.mu.RUnlock()
	return
}

func (s *tsafeStrings) ContainsOneOf(values ...string) (ok bool) {
	s.mu.RLock()
	ok = s.values.ContainsOneOf(values...)
	s.mu.RUnlock()
	return
}

func (s *tsafeStrings) Copy() TSafeStrings {
	s2 := &tsafeStrings{mu: &sync.RWMutex{}}
	s.mu.RLock()
	s2.values = s.values.Copy()
	s.mu.RUnlock()
	return s2
}

func (s *tsafeStrings) Diff(s2 Strings) (out Strings) {
	s.mu.RLock()
	out = s.values.Diff(s2)
	s.mu.RUnlock()
	return
}

func (s *tsafeStrings) Empty() (ok bool) {
	s.mu.RLock()
	ok = s.values.Empty()
	s.mu.RUnlock()
	return
}

func (s *tsafeStrings) Equal(s2 Strings) (ok bool) {
	s.mu.RLock()
	ok = s.values.Equal(s2)
	s.mu.RUnlock()
	return
}

func (s *tsafeStrings) Find(matcher func(v string) bool) (v string, ok bool) {
	s.mu.RLock()
	v, ok = s.values.Find(matcher)
	s.mu.RUnlock()
	return
}

func (s *tsafeStrings) FindAll(matcher func(v string) bool) (v Strings) {
	s.mu.RLock()
	v = s.values.FindAll(matcher)
	s.mu.RUnlock()
	return
}

func (s *tsafeStrings) First() (v string, ok bool) {
	s.mu.RLock()
	v, ok = s.values.First()
	s.mu.RUnlock()
	return
}

func (s *tsafeStrings) Get(i int) (v string, ok bool) {
	s.mu.RLock()
	v, ok = s.values.Get(i)
	s.mu.RUnlock()
	return
}

func (s *tsafeStrings) Intersect(s2 Strings) (v Strings) {
	s.mu.RLock()
	v = s.values.Intersect(s2)
	s.mu.RUnlock()
	return
}

func (s *tsafeStrings) Last() (v string, ok bool) {
	s.mu.RLock()
	v, ok = s.values.Last()
	s.mu.RUnlock()
	return
}

func (s *tsafeStrings) Len() (v int) {
	s.mu.RLock()
	v = s.values.Len()
	s.mu.RUnlock()
	return
}

func (s *tsafeStrings) Take(n int) (v Strings) {
	s.mu.RLock()
	v = s.values.Take(n)
	s.mu.RUnlock()
	return
}

func (s *tsafeStrings) S() (out []any) {
	s.mu.RLock()
	out = s.values.S()
	s.mu.RUnlock()
	return
}

func (s *tsafeStrings) Strings() (v Strings) {
	s.mu.RLock()
	v = s.values.Copy()
	s.mu.RUnlock()
	return
}
