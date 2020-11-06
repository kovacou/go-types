// Copyright © 2019 Alexandre Kovac <contact@kovacou.com>.
//
// Use of this source code is governed by an MIT-style
// license that can be found in the LICENSE file.

package types

// Bools is a slice of bool.
type Bools []bool

// Reset the slice.
func (s *Bools) Reset() {
	*s = []bool{}
}

// Add new elements to the slice.
func (s *Bools) Add(values ...bool) {
	*s = append(*s, values...)
}

// Contains say if "s" contains "values".
func (s Bools) Contains(values ...bool) bool {
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
func (s Bools) ContainsOneOf(values ...bool) bool {
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
func (s Bools) Copy() Bools {
	out := make(Bools, s.Len())
	copy(out, s)
	return out
}

// Diff returns the difference between "s" and "s2".
func (s Bools) Diff(s2 Bools) Bools {
	if s.Empty() {
		return s2.Copy()
	} else if s2.Empty() {
		return s.Copy()
	}

	out := Bools{}

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
func (s Bools) Empty() bool {
	return len(s) == 0
}

// Equal says if "s" and "s2" are equal.
func (s Bools) Equal(s2 Bools) bool {
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

// // Find the first element matching the pattern.
// func (s Bools) Find(matcher func(v string) bool) (string, bool) {
// 	for _, val := range s {
// 		if matcher(val) {
// 			return val, true
// 		}
// 	}
// 	return "", false
// }

// // FindAll elements matching the pattern.
// func (s Bools) FindAll(matcher func(v string) bool) Bools {
// 	out := Bools{}
// 	for _, val := range s {
// 		if matcher(val) {
// 			out = append(out, val)
// 		}
// 	}
// 	return out
// }

// First return the value of the first element.
func (s Bools) First() (bool, bool) {
	if len(s) > 0 {
		return s[0], true
	}
	return false, false
}

// Get the element "i" and say if it has been found.
func (s Bools) Get(i int) (bool, bool) {
	if i > s.Len() {
		return false, false
	}
	return s[i], true
}

// Intersect return the intersection between "s" and "s2".
func (s Bools) Intersect(s2 Bools) Bools {
	out := Bools{}
	for _, v := range s {
		if s2.Contains(v) {
			out = append(out, v)
		}
	}
	return out
}

// Last return the value of the last element.
func (s Bools) Last() (bool, bool) {
	if n := len(s); n > 0 {
		return s[n-1], true
	}
	return false, false
}

// Len returns the size of the slice.
func (s Bools) Len() int {
	return len(s)
}

// Take n element and return a new slice.
func (s Bools) Take(n int) (out Bools) {
	if n < 0 || n > s.Len() {
		return s
	}
	return s[:n].Copy()
}

// ----------------- CONVERTING METHOD -----------------

// S convert s into []interface{}
func (s Bools) S() (out []interface{}) {
	for _, v := range s {
		out = append(out, v)
	}
	return
}
