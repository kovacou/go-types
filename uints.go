// Copyright © 2019 Alexandre Kovac <contact@kovacou.com>.
//
// Use of this source code is governed by an MIT-style
// license that can be found in the LICENSE file.

package types

// Uints is a slice of uint.
type Uints []uint

// Reset the slice.
func (s *Uints) Reset() {
	*s = []uint{}
}

// Contains says if "s" contains "values".
func (s Uints) Contains(values ...uint) bool {
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
func (s Uints) ContainsOneOf(values ...uint) bool {
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
func (s Uints) Copy() Uints {
	out := make(Uints, s.Len())
	copy(out, s)
	return out
}

// Diff return the difference between "s" and "s2".
func (s Uints) Diff(s2 Uints) Uints {
	if s.Empty() {
		return s2.Copy()
	} else if s2.Empty() {
		return s.Copy()
	}

	out := Uints{}

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
func (s Uints) Empty() bool {
	return len(s) == 0
}

// Equal says if "s" and "s2" are equal.
func (s Uints) Equal(s2 Uints) bool {
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
func (s Uints) Find(matcher func(v uint) bool) (uint, bool) {
	for _, val := range s {
		if matcher(val) {
			return val, true
		}
	}
	return 0, false
}

// FindAll elements matching the pattern.
func (s Uints) FindAll(matcher func(v uint) bool) Uints {
	out := Uints{}
	for _, val := range s {
		if matcher(val) {
			out = append(out, val)
		}
	}
	return out
}

// First return the value of the first element.
func (s Uints) First() (uint, bool) {
	if len(s) > 0 {
		return s[0], true
	}
	return 0, false
}

// Get the element "i" and say if it has been found.
func (s Uints) Get(i int) (uint, bool) {
	if i > s.Len() {
		return 0, false
	}
	return s[i], true
}

// Intersect return the intersection between "s" and "s2".
func (s Uints) Intersect(s2 Uints) Uints {
	out := Uints{}
	for _, v := range s {
		if s2.Contains(v) {
			out = append(out, v)
		}
	}
	return out
}

// Last return the value of the last element.
func (s Uints) Last() (uint, bool) {
	if n := len(s); n > 0 {
		return s[n-1], true
	}
	return 0, false
}

// Len return the size of the slice.
func (s Uints) Len() int {
	return len(s)
}

// Mean of the slice.
func (s Uints) Mean() (mean float64) {
	return float64(s.Sum()) / float64(s.Len())
}

// Sum of the slice.
func (s Uints) Sum() (sum uint) {
	for _, v := range s {
		sum += v
	}
	return
}

// Take n element and return a new slice.
func (s Uints) Take(n int) (out Uints) {
	if n < 0 || n > s.Len() {
		return s
	}
	return s[:n].Copy()
}

// ----------------- CONVERTING METHOD -----------------

// S convert Uints into []interface{}
func (s Uints) S() (out []interface{}) {
	for _, v := range s {
		out = append(out, v)
	}
	return
}
