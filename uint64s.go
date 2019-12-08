// Copyright © 2019 Alexandre Kovac <contact@kovacou.com>.
//
// Use of this source code is governed by an MIT-style
// license that can be found in the LICENSE file.

package types

// Uint64s is a slice of uint64.
type Uint64s []uint64

// Reset the slice.
func (s *Uint64s) Reset() {
	*s = []uint64{}
}

// Contains says if "s" contains "values".
func (s Uint64s) Contains(values ...uint64) bool {
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
func (s Uint64s) ContainsOneOf(values ...uint64) bool {
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
func (s Uint64s) Copy() Uint64s {
	out := make(Uint64s, s.Len())
	copy(out, s)
	return out
}

// Diff return the difference between "s" and "s2".
func (s Uint64s) Diff(s2 Uint64s) Uint64s {
	if s.Empty() {
		return s2.Copy()
	} else if s2.Empty() {
		return s.Copy()
	}

	out := Uint64s{}

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
func (s Uint64s) Empty() bool {
	return len(s) == 0
}

// Equal says if "s" and "s2" are equal.
func (s Uint64s) Equal(s2 Uint64s) bool {
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
func (s Uint64s) Find(matcher func(v uint64) bool) (uint64, bool) {
	for _, val := range s {
		if matcher(val) {
			return val, true
		}
	}
	return 0, false
}

// FindAll elements matching the pattern.
func (s Uint64s) FindAll(matcher func(v uint64) bool) Uint64s {
	out := Uint64s{}
	for _, val := range s {
		if matcher(val) {
			out = append(out, val)
		}
	}
	return out
}

// First return the value of the first element.
func (s Uint64s) First() (uint64, bool) {
	if len(s) > 0 {
		return s[0], true
	}
	return 0, false
}

// Get the element "i" and say if it has been found.
func (s Uint64s) Get(i int) (uint64, bool) {
	if i > s.Len() {
		return 0, false
	}
	return s[i], true
}

// Intersect return the intersection between "s" and "s2".
func (s Uint64s) Intersect(s2 Uint64s) Uint64s {
	out := Uint64s{}
	for _, v := range s {
		if s2.Contains(v) {
			out = append(out, v)
		}
	}
	return out
}

// Last return the value of the last element.
func (s Uint64s) Last() (uint64, bool) {
	if n := len(s); n > 0 {
		return s[n-1], true
	}
	return 0, false
}

// Len return the size of the slice.
func (s Uint64s) Len() int {
	return len(s)
}

// Take n element and return a new slice.
func (s Uint64s) Take(n int) (out Uint64s) {
	if n < 0 || n > s.Len() {
		return s
	}
	return s[:n].Copy()
}

// ----------------- CONVERTING METHOD -----------------

// S convert Uint64s into []interface{}
func (s Uint64s) S() (out []interface{}) {
	for _, v := range s {
		out = append(out, v)
	}
	return
}
