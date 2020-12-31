// Copyright © 2019 Alexandre Kovac <contact@kovacou.com>.
//
// Use of this source code is governed by an MIT-style
// license that can be found in the LICENSE file.

package types

// Int64s is a slice of int64.
type Int64s []int64

// Int64NoZero is a filter for LenIf, SumIf, MeanIf.
func Int64NoZero(v int64) bool {
	return v != 0
}

// Reset the slice.
func (s *Int64s) Reset() {
	*s = []int64{}
}

// Add new elements to the slice.
func (s *Int64s) Add(values ...int64) {
	*s = append(*s, values...)
}

// Contains say if "s" contains "values".
func (s Int64s) Contains(values ...int64) bool {
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
func (s Int64s) ContainsOneOf(values ...int64) bool {
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
func (s Int64s) Copy() Int64s {
	out := make(Int64s, s.Len())
	copy(out, s)
	return out
}

// Diff returns the difference between "s" and "s2".
func (s Int64s) Diff(s2 Int64s) Int64s {
	if s.Empty() {
		return s2.Copy()
	} else if s2.Empty() {
		return s.Copy()
	}

	out := Int64s{}

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
func (s Int64s) Empty() bool {
	return len(s) == 0
}

// Equal says if "s" and "s2" are equal.
func (s Int64s) Equal(s2 Int64s) bool {
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
func (s Int64s) Find(matcher func(v int64) bool) (int64, bool) {
	for _, val := range s {
		if matcher(val) {
			return val, true
		}
	}
	return 0, false
}

// FindAll elements matching the pattern.
func (s Int64s) FindAll(matcher func(v int64) bool) Int64s {
	out := Int64s{}
	for _, val := range s {
		if matcher(val) {
			out = append(out, val)
		}
	}
	return out
}

// First return the value of the first element.
func (s Int64s) First() (int64, bool) {
	if len(s) > 0 {
		return s[0], true
	}
	return 0, false
}

// Get the element "i" and say if it has been found.
func (s Int64s) Get(i int) (int64, bool) {
	if i > s.Len() {
		return 0, false
	}
	return s[i], true
}

// Intersect return the intersection between "s" and "s2".
func (s Int64s) Intersect(s2 Int64s) Int64s {
	out := Int64s{}
	for _, v := range s {
		if s2.Contains(v) {
			out = append(out, v)
		}
	}
	return out
}

// Last return the value of the last element.
func (s Int64s) Last() (int64, bool) {
	if n := len(s); n > 0 {
		return s[n-1], true
	}
	return 0, false
}

// Len return the size of the slice.
func (s Int64s) Len() int {
	return len(s)
}

// LenIf return the size of the slice if the filter is valid.
func (s Int64s) LenIf(f func(v int64) bool) (n int) {
	for _, v := range s {
		if f(v) {
			n++
		}
	}
	return
}

// Mean of the slice.
func (s Int64s) Mean() (mean float64) {
	return float64(s.Sum()) / float64(s.Len())
}

// MeanIf the filter is valid of the slice.
func (s Int64s) MeanIf(f func(v int64) bool) (mean float64) {
	return float64(s.SumIf(f)) / float64(s.LenIf(f))
}

// Sum of the slice.
func (s Int64s) Sum() (sum int64) {
	for _, v := range s {
		sum += v
	}
	return
}

// SumIf the filter is valid of the slice.
func (s Int64s) SumIf(f func(v int64) bool) (sum int64) {
	for _, v := range s {
		if f(v) {
			sum += v
		}
	}
	return
}

// Take n element and return a new slice.
func (s Int64s) Take(n int) (out Int64s) {
	if n < 0 || n > s.Len() {
		return s
	}
	return s[:n].Copy()
}

// ----------------- CONVERTING METHOD -----------------

// S convert s into []interface{}
func (s Int64s) S() (out []interface{}) {
	for _, v := range s {
		out = append(out, v)
	}
	return
}
