package types

import "sync"

type TSafeInt64s interface {
	// Reset the slice.
	Reset()

	// Contains say if "s" contains "values".
	Contains(...int64) bool

	// ContainsOneOf says if "s" contains one of the "values".
	ContainsOneOf(...int64) bool

	// Copy create a new copy of the slice.
	Copy() TSafeInt64s

	// Diff returns the difference between "s" and "s2".
	Diff(Int64s) Int64s

	// Empty says if the slice is empty.
	Empty() bool

	// Equal says if "s" and "s2" are equal.
	Equal(Int64s) bool

	// Find the first element matching the pattern.
	Find(func(v int64) bool) (int64, bool)

	// FindAll elements matching the pattern.
	FindAll(func(v int64) bool) Int64s

	// First return the value of the first element.
	First() (int64, bool)

	// Get the element "i" and say if it has been found.
	Get(int) (int64, bool)

	// Intersect return the intersection between "s" and "s2".
	Intersect(Int64s) Int64s

	// Last return the value of the last element.
	Last() (int64, bool)

	// Len returns the size of the slice.
	Len() int

	// Take n element and return a new slice.
	Take(int) Int64s

	// S convert s into []any
	S() []any

	// S convert s into Int64s
	Int64s() Int64s
}

func SyncInt64s() TSafeInt64s {
	return &tsafeInt64s{&sync.RWMutex{}, Int64s{}}
}

type tsafeInt64s struct {
	mu     *sync.RWMutex
	values Int64s
}

func (s *tsafeInt64s) Reset() {
	s.mu.Lock()
	s.values.Reset()
	s.mu.Unlock()
}

func (s *tsafeInt64s) Contains(values ...int64) (ok bool) {
	s.mu.RLock()
	ok = s.values.Contains(values...)
	s.mu.RUnlock()
	return
}

func (s *tsafeInt64s) ContainsOneOf(values ...int64) (ok bool) {
	s.mu.RLock()
	ok = s.values.ContainsOneOf(values...)
	s.mu.RUnlock()
	return
}

func (s *tsafeInt64s) Copy() TSafeInt64s {
	s2 := &tsafeInt64s{mu: &sync.RWMutex{}}
	s.mu.RLock()
	s2.values = s.values.Copy()
	s.mu.RUnlock()
	return s2
}

func (s *tsafeInt64s) Diff(s2 Int64s) (out Int64s) {
	s.mu.RLock()
	out = s.values.Diff(s2)
	s.mu.RUnlock()
	return
}

func (s *tsafeInt64s) Empty() (ok bool) {
	s.mu.RLock()
	ok = s.values.Empty()
	s.mu.RUnlock()
	return
}

func (s *tsafeInt64s) Equal(s2 Int64s) (ok bool) {
	s.mu.RLock()
	ok = s.values.Equal(s2)
	s.mu.RUnlock()
	return
}

func (s *tsafeInt64s) Find(matcher func(v int64) bool) (v int64, ok bool) {
	s.mu.RLock()
	v, ok = s.values.Find(matcher)
	s.mu.RUnlock()
	return
}

func (s *tsafeInt64s) FindAll(matcher func(v int64) bool) (v Int64s) {
	s.mu.RLock()
	v = s.values.FindAll(matcher)
	s.mu.RUnlock()
	return
}

func (s *tsafeInt64s) First() (v int64, ok bool) {
	s.mu.RLock()
	v, ok = s.values.First()
	s.mu.RUnlock()
	return
}

func (s *tsafeInt64s) Get(i int) (v int64, ok bool) {
	s.mu.RLock()
	v, ok = s.values.Get(i)
	s.mu.RUnlock()
	return
}

func (s *tsafeInt64s) Intersect(s2 Int64s) (v Int64s) {
	s.mu.RLock()
	v = s.values.Intersect(s2)
	s.mu.RUnlock()
	return
}

func (s *tsafeInt64s) Last() (v int64, ok bool) {
	s.mu.RLock()
	v, ok = s.values.Last()
	s.mu.RUnlock()
	return
}

func (s *tsafeInt64s) Len() (v int) {
	s.mu.RLock()
	v = s.values.Len()
	s.mu.RUnlock()
	return
}

func (s *tsafeInt64s) Take(n int) (v Int64s) {
	s.mu.RLock()
	v = s.values.Take(n)
	s.mu.RUnlock()
	return
}

func (s *tsafeInt64s) S() (out []any) {
	s.mu.RLock()
	out = s.values.S()
	s.mu.RUnlock()
	return
}

func (s *tsafeInt64s) Int64s() (v Int64s) {
	s.mu.RLock()
	v = s.values.Copy()
	s.mu.RUnlock()
	return
}
