// Copyright © 2019 Alexandre Kovac <contact@kovacou.com>.
//
// Use of this source code is governed by an MIT-style
// license that can be found in the LICENSE file.

package types

// Floats is a slice of float64.
type Floats []float64

// Float64NoZero is a filter for LenIf, SumIf, MeanIf.
func Float64NoZero(v float64) bool {
	return v != 0
}

// Reset the slice.
func (s *Floats) Reset() {
	*s = []float64{}
}

// Add new elements to the slice.
func (s *Floats) Add(values ...float64) {
	*s = append(*s, values...)
}

// Contains say if "s" contains "values".
func (s Floats) Contains(values ...float64) bool {
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
func (s Floats) ContainsOneOf(values ...float64) bool {
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
func (s Floats) Copy() Floats {
	out := make(Floats, s.Len())
	copy(out, s)
	return out
}

// Diff returns the difference between "s" and "s2".
func (s Floats) Diff(s2 Floats) Floats {
	if s.Empty() {
		return s2.Copy()
	} else if s2.Empty() {
		return s.Copy()
	}

	out := Floats{}

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
func (s Floats) Empty() bool {
	return len(s) == 0
}

// Equal says if "s" and "s2" are equal.
func (s Floats) Equal(s2 Floats) bool {
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
func (s Floats) Find(matcher func(v float64) bool) (float64, bool) {
	for _, val := range s {
		if matcher(val) {
			return val, true
		}
	}
	return 0, false
}

// FindAll elements matching the pattern.
func (s Floats) FindAll(matcher func(v float64) bool) Floats {
	out := Floats{}
	for _, val := range s {
		if matcher(val) {
			out = append(out, val)
		}
	}
	return out
}

// First return the value of the first element.
func (s Floats) First() (float64, bool) {
	if len(s) > 0 {
		return s[0], true
	}
	return 0, false
}

// Get the element "i" and say if it has been found.
func (s Floats) Get(i int) (float64, bool) {
	if i > s.Len() {
		return 0, false
	}
	return s[i], true
}

// Intersect return the intersection between "s" and "s2".
func (s Floats) Intersect(s2 Floats) Floats {
	out := Floats{}
	for _, v := range s {
		if s2.Contains(v) {
			out = append(out, v)
		}
	}
	return out
}

// Last return the value of the last element.
func (s Floats) Last() (float64, bool) {
	if n := len(s); n > 0 {
		return s[n-1], true
	}
	return 0, false
}

// Len returns the size of the slice.
func (s Floats) Len() int {
	return len(s)
}

// LenIf return the size of the slice if the filter is valid.
func (s Floats) LenIf(f func(v float64) bool) (n int) {
	for _, v := range s {
		if f(v) {
			n++
		}
	}
	return
}

// Mean of the slice.
func (s Floats) Mean() (mean float64) {
	return s.Sum() / float64(s.Len())
}

// MeanIf the filter is valid of the slice.
func (s Floats) MeanIf(f func(v float64) bool) (mean float64) {
	return s.SumIf(f) / float64(s.LenIf(f))
}

// Sum of the slice.
func (s Floats) Sum() (sum float64) {
	for _, v := range s {
		sum += v
	}
	return
}

// SumIf the filter is valid of the slice.
func (s Floats) SumIf(f func(v float64) bool) (sum float64) {
	for _, v := range s {
		if f(v) {
			sum += v
		}
	}
	return
}

// Take n element and return a new slice.
func (s Floats) Take(n int) (out Floats) {
	if n < 0 || n > s.Len() {
		return s
	}
	return s[:n].Copy()
}

// ----------------- CONVERTING METHOD -----------------

// S convert s to []interface{}
func (s Floats) S() (out []interface{}) {
	for _, v := range s {
		out = append(out, v)
	}
	return
}
