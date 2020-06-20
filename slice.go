// Copyright © 2019 Alexandre Kovac <contact@kovacou.fr>.
//
// Use of this source code is governed by an MIT-style
// license that can be found in the LICENSE file.

package types

// Slice is a slice of interface{}.
type Slice []interface{}

// Reset the slice.
func (s *Slice) Reset() {
	*s = []interface{}{}
}

// Contains say if "s" contains "values".
func (s Slice) Contains(values ...interface{}) bool {
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
func (s Slice) ContainsOneOf(values ...interface{}) bool {
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
func (s Slice) Copy() Slice {
	out := make(Slice, s.Len())
	copy(out, s)
	return out
}

// Diff returns the difference between "s" and "s2".
func (s Slice) Diff(s2 Slice) Slice {
	if s.Empty() {
		return s2.Copy()
	} else if s2.Empty() {
		return s.Copy()
	}

	out := Slice{}

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
func (s Slice) Empty() bool {
	return len(s) == 0
}

// Equal says if "s" and "s2" are equal.
func (s Slice) Equal(s2 Slice) bool {
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
func (s Slice) Find(matcher func(v interface{}) bool) (interface{}, bool) {
	for _, val := range s {
		if matcher(val) {
			return val, true
		}
	}
	return "", false
}

// FindAll elements matching the pattern.
func (s Slice) FindAll(matcher func(v interface{}) bool) Slice {
	out := Slice{}
	for _, val := range s {
		if matcher(val) {
			out = append(out, val)
		}
	}
	return out
}

// First return the value of the first element.
func (s Slice) First() (interface{}, bool) {
	if len(s) > 0 {
		return s[0], true
	}
	return nil, false
}

// Get the element "i" and say if it has been found.
func (s Slice) Get(i int) (interface{}, bool) {
	if i > s.Len() {
		return "", false
	}
	return s[i], true
}

// Intersect return the intersection between "s" and "s2".
func (s Slice) Intersect(s2 Slice) Slice {
	out := Slice{}
	for _, v := range s {
		if s2.Contains(v) {
			out = append(out, v)
		}
	}
	return out
}

// Last return the value of the last element.
func (s Slice) Last() (interface{}, bool) {
	if n := len(s); n > 0 {
		return s[n-1], true
	}
	return "", false
}

// Len returns the size of the slice.
func (s Slice) Len() int {
	return len(s)
}

// Take n element and return a new slice.
func (s Slice) Take(n int) (out Slice) {
	if n < 0 || n > s.Len() {
		return s
	}
	return s[:n].Copy()
}

// ----------------- CONVERTING METHOD -----------------

// S convert s into []interface{}
func (s Slice) S() (out []interface{}) {
	for _, v := range s {
		out = append(out, v)
	}
	return
}
