// Copyright © 2019 Alexandre Kovac <contact@kovacou.fr>.
//
// Use of this source code is governed by an MIT-style
// license that can be found in the LICENSE file.

package types

// Bytes is a slice of byte.
type Bytes []byte

// Reset the slice.
func (s *Bytes) Reset() {
	*s = []byte{}
}

// Contains says if "s" contains "values".
func (s Bytes) Contains(values ...byte) bool {
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
func (s Bytes) ContainsOneOf(values ...byte) bool {
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
func (s Bytes) Copy() Bytes {
	out := make(Bytes, s.Len())
	copy(out, s)
	return out
}

// Diff return the difference between "s" and "s2".
func (s Bytes) Diff(s2 Bytes) Bytes {
	if s.Empty() {
		return s2.Copy()
	} else if s2.Empty() {
		return s.Copy()
	}

	out := Bytes{}

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
func (s Bytes) Empty() bool {
	return len(s) == 0
}

// Equal says if "s" and "s2" are equal.
func (s Bytes) Equal(s2 Bytes) bool {
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

// First return the value of the first element.
func (s Bytes) First() (byte, bool) {
	if len(s) > 0 {
		return s[0], true
	}
	return 0, false
}

// Get the element "i" and say if it has been found.
func (s Bytes) Get(i int) (byte, bool) {
	if i > s.Len() {
		return 0, false
	}
	return s[i], true
}

// Intersect return the intersection between "s" and "s2".
func (s Bytes) Intersect(s2 Bytes) Bytes {
	out := Bytes{}
	for _, v := range s {
		if s2.Contains(v) {
			out = append(out, v)
		}
	}
	return out
}

// Last return the value of the last element.
func (s Bytes) Last() (byte, bool) {
	if n := len(s); n > 0 {
		return s[n-1], true
	}
	return 0, false
}

// Len return the size of the slice.
func (s Bytes) Len() int {
	return len(s)
}

// Take n element and return a new slice.
func (s Bytes) Take(n int) (out Bytes) {
	if n < 0 || n > s.Len() {
		return s
	}
	return s[:n].Copy()
}

// ----------------- CONVERTING METHOD -----------------

// S convert Bytes into []interface{}
func (s Bytes) S() (out []interface{}) {
	for _, v := range s {
		out = append(out, v)
	}
	return
}
