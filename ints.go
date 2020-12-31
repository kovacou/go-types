// Copyright © 2019 Alexandre Kovac <contact@kovacou.com>.
//
// Use of this source code is governed by an MIT-style
// license that can be found in the LICENSE file.

package types

// Ints is a slice of int.
type Ints []int

// Reset the slice.
func (s *Ints) Reset() {
	*s = []int{}
}

// Add new elements to the slice.
func (s *Ints) Add(values ...int) {
	*s = append(*s, values...)
}

// Contains say if "s" contains "values".
func (s Ints) Contains(values ...int) bool {
	findNum := 0
	for i := range s {
		for _, value := range values {
			if s[i] == value {
				findNum++
				break
			}
		}
	}
	return findNum == len(values)
}

// ContainsOneOf says if "s" contains one of the "values".
func (s Ints) ContainsOneOf(values ...int) bool {
	for _, value := range values {
		for i := range s {
			if s[i] == value {
				return true
			}
		}
	}
	return false
}

// Copy create a new copy of the slice.
func (s Ints) Copy() Ints {
	out := make(Ints, s.Len())
	copy(out, s)
	return out
}

// Diff returns the difference between "s" and "s2".
func (s Ints) Diff(s2 Ints) Ints {
	if s.Empty() {
		return s2.Copy()
	} else if s2.Empty() {
		return s.Copy()
	}

	out := Ints{}

	if len(s) >= len(s2) {
		for _, v := range s {
			if !s2.Contains(v) {
				out = append(out, v)
			}
		}
	}

	for _, v := range s2 {
		if !s.Contains(v) {
			out = append(out, v)
		}
	}

	return out
}

// Empty says if the slice is empty.
func (s Ints) Empty() bool {
	return len(s) == 0
}

// Equal says if "s" and "s2" are equal.
func (s Ints) Equal(s2 Ints) bool {
	if len(s) == len(s2) {
		for k := range s2 {
			if s2[k] != s[k] {
				return false
			}
		}
		return true
	}
	return false
}

// Find the first element matching the pattern.
func (s Ints) Find(matcher func(v int) bool) (int, bool) {
	for _, val := range s {
		if matcher(val) {
			return val, true
		}
	}
	return 0, false
}

// FindAll elements matching the pattern.
func (s Ints) FindAll(matcher func(v int) bool) Ints {
	out := Ints{}
	for _, val := range s {
		if matcher(val) {
			out = append(out, val)
		}
	}
	return out
}

// First return the value of the first element.
func (s Ints) First() (int, bool) {
	if len(s) > 0 {
		return s[0], true
	}
	return 0, false
}

// Get the element "i" and say if it has been found.
func (s Ints) Get(i int) (int, bool) {
	if i > s.Len() {
		return 0, false
	}
	return s[i], true
}

// Intersect return the intersection between "s" and "s2".
func (s Ints) Intersect(s2 Ints) Ints {
	out := Ints{}
	for _, v := range s {
		if s2.Contains(v) {
			out = append(out, v)
		}
	}
	return out
}

// Last return the value of the last element.
func (s Ints) Last() (int, bool) {
	if n := len(s); n > 0 {
		return s[n-1], true
	}
	return 0, false
}

// Len return the size of the slice.
func (s Ints) Len() int {
	return len(s)
}

// LenIf return the size of the slice if the filter is valid.
func (s Ints) LenIf(f func(v int) bool) (n int) {
	for _, v := range s {
		if f(v) {
			n++
		}
	}
	return
}

// Mean of the slice.
func (s Ints) Mean() (mean float64) {
	return float64(s.Sum()) / float64(s.Len())
}

// MeanIf the filter is valid of the slice.
func (s Ints) MeanIf(f func(v int) bool) (mean float64) {
	return float64(s.SumIf(f)) / float64(s.LenIf(f))
}

// Sum of the slice.
func (s Ints) Sum() (sum int) {
	for _, v := range s {
		sum += v
	}
	return
}

// SumIf the filter is valid of the slice.
func (s Ints) SumIf(f func(v int) bool) (sum int) {
	for _, v := range s {
		if f(v) {
			sum += v
		}
	}
	return
}

// Take n element and return a new slice.
func (s Ints) Take(n int) (out Ints) {
	if n < 0 || n > s.Len() {
		return s
	}
	return s[:n].Copy()
}

// ----------------- CONVERTING METHOD -----------------

// S convert s into []interface{}
func (s Ints) S() (out []interface{}) {
	for _, v := range s {
		out = append(out, v)
	}
	return
}
