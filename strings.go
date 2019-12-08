// Copyright © 2019 Alexandre Kovac <contact@kovacou.com>.
//
// Use of this source code is governed by an MIT-style
// license that can be found in the LICENSE file.

package types

// Strings is a slice of string.
type Strings []string

// Reset the slice.
func (s *Strings) Reset() {
	*s = []string{}
}

// Contains says if "s" contains "values".
func (s Strings) Contains(values ...string) bool {
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
func (s Strings) ContainsOneOf(values ...string) bool {
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
func (s Strings) Copy() Strings {
	out := make(Strings, s.Len())
	copy(out, s)
	return out
}

// Diff return the difference between "s" and "s2".
func (s Strings) Diff(s2 Strings) Strings {
	if s.Empty() {
		return s2.Copy()
	} else if s2.Empty() {
		return s.Copy()
	}

	out := Strings{}

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
func (s Strings) Empty() bool {
	return len(s) == 0
}

// Equal says if "s" and "s2" are equal.
func (s Strings) Equal(s2 Strings) bool {
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
func (s Strings) Find(matcher func(v string) bool) (string, bool) {
	for _, val := range s {
		if matcher(val) {
			return val, true
		}
	}
	return "", false
}

// FindAll elements matching the pattern.
func (s Strings) FindAll(matcher func(v string) bool) Strings {
	out := Strings{}
	for _, val := range s {
		if matcher(val) {
			out = append(out, val)
		}
	}
	return out
}

// First return the value of the first element.
func (s Strings) First() (string, bool) {
	if len(s) > 0 {
		return s[0], true
	}
	return "", false
}

// Get the element "i" and say if it has been found.
func (s Strings) Get(i int) (string, bool) {
	if i > s.Len() {
		return "", false
	}
	return s[i], true
}

// Intersect return the intersection between "s" and "s2".
func (s Strings) Intersect(s2 Strings) Strings {
	out := Strings{}
	for _, v := range s {
		if s2.Contains(v) {
			out = append(out, v)
		}
	}
	return out
}

// Last return the value of the last element.
func (s Strings) Last() (string, bool) {
	if n := len(s); n > 0 {
		return s[n-1], true
	}
	return "", false
}

// Len return the size of the slice.
func (s Strings) Len() int {
	return len(s)
}

// Take n element and return a new slice.
func (s Strings) Take(n int) (out Strings) {
	if n < 0 || n > s.Len() {
		return s
	}
	return s[:n].Copy()
}

// ----------------- CONVERTING METHOD -----------------

// S convert Strings into []interface{}
func (s Strings) S() (out []interface{}) {
	for _, v := range s {
		out = append(out, v)
	}
	return
}
