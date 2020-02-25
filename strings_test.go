// Copyright © 2019 Alexandre Kovac <contact@kovacou.com>.
//
// Use of this source code is governed by an MIT-style
// license that can be found in the LICENSE file.

package types

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func makeDefaultStrings() Strings {
	return Strings{"1", "2", "3", "4", "5"}
}

func TestStringsReset(t *testing.T) {
	s := &Strings{"1", "2", "3", "4", "5"}
	s2 := s.Copy()
	s.Reset()

	assert.Empty(t, s)
	assert.NotEmpty(t, s2)

	s2.Reset()
	assert.Empty(t, s2)
}

func TestStringsCopy(t *testing.T) {
	s := makeDefaultStrings()
	s2 := s.Copy()

	assert.Equal(t, s, s2)
	s2 = append(s2, "6", "7", "8")
	assert.NotEqual(t, s, s2)
}

func TestStringsDiff(t *testing.T) {
	s := makeDefaultStrings()
	s2 := Strings{}

	assert.Equal(t, s, s2.Diff(s))
	assert.Equal(t, s, s.Diff(s2))

	s3 := Strings{"1", "3", "5"}
	assert.Equal(t, Strings{"2", "4"}, s.Diff(s3))
	assert.Equal(t, Strings{"2", "4"}, s3.Diff(s))
}

func TestStringsContains(t *testing.T) {
	s := makeDefaultStrings()
	assert.True(t, s.Contains("1"))
	assert.True(t, s.Contains("1", "2", "3"))
	assert.True(t, s.Contains("1", "4", "5"))
	assert.False(t, s.Contains("0"))
	assert.False(t, s.Contains("0", "5", "1"))
	assert.True(t, s.Contains(makeDefaultStrings()...))
}

func TestStringsContainsOneOf(t *testing.T) {
	s := makeDefaultStrings()
	assert.True(t, s.ContainsOneOf("1"))
	assert.True(t, s.ContainsOneOf("1", "2", "3"))
	assert.True(t, s.ContainsOneOf("1", "4", "5"))
	assert.True(t, s.ContainsOneOf("0", "5", "1"))
	assert.False(t, s.ContainsOneOf("0"))
	assert.False(t, s.ContainsOneOf("0", "6"))

}

func TestStringsEmpty(t *testing.T) {
	s := Strings{}
	assert.True(t, s.Empty())
	s = makeDefaultStrings()
	assert.False(t, s.Empty())
}

func TestStringsEqual(t *testing.T) {
	s := makeDefaultStrings()
	s2 := makeDefaultStrings()
	s3 := Strings{}
	s4 := Strings{"6", "7", "8", "9", "10"}
	assert.True(t, s.Equal(s2))
	assert.False(t, s.Equal(s3))
	assert.False(t, s3.Equal(s2))
	assert.False(t, s4.Equal(s2))
	assert.False(t, s2.Equal(s4))
}

func TestStringsFirst(t *testing.T) {
	s := Strings{"1", "2", "3"}
	s2 := Strings{}
	s3 := Strings{"2", "4", "5"}

	{
		v, ok := s.First()
		assert.True(t, ok)
		assert.Equal(t, string("1"), v)
	}

	{
		v, ok := s2.First()
		assert.False(t, ok)
		assert.Equal(t, "", v)
	}

	{
		v, ok := s3.First()
		assert.True(t, ok)
		assert.Equal(t, "2", v)
	}
}

func TestStringsGet(t *testing.T) {
	s := Strings{"1", "2", "3"}
	s2 := Strings{}
	s3 := Strings{"2", "4", "5"}

	{
		v, ok := s.Get(0)
		assert.True(t, ok)
		assert.Equal(t, "1", v)
	}

	{
		v, ok := s2.Get(6)
		assert.False(t, ok)
		assert.Equal(t, "", v)
	}

	{
		v, ok := s3.Get(2)
		assert.True(t, ok)
		assert.Equal(t, "5", v)
	}
}

func TestStringsIntersect(t *testing.T) {
	s := Strings{"1", "2", "3"}
	s2 := Strings{"2", "3", "4", "5"}
	s3 := Strings{}

	assert.Equal(t, Strings{"2", "3"}, s.Intersect(s2))
	assert.Equal(t, s3, s3.Intersect(s))
	assert.Equal(t, s3, s.Intersect(s3))
}

func TestStringsLast(t *testing.T) {
	s := Strings{"1", "2", "3"}
	s2 := Strings{}
	s3 := Strings{"2", "4", "5"}

	{
		v, ok := s.Last()
		assert.True(t, ok)
		assert.Equal(t, "3", v)
	}

	{
		v, ok := s2.Last()
		assert.False(t, ok)
		assert.Equal(t, "", v)
	}

	{
		v, ok := s3.Last()
		assert.True(t, ok)
		assert.Equal(t, "5", v)
	}
}

func TestStringsTake(t *testing.T) {
	s := Strings{"1", "2", "3"}
	s2 := Strings{}

	assert.Equal(t, s, s.Take(3))
	assert.Equal(t, Strings{"1", "2"}, s.Take(2))
	assert.Equal(t, s, s.Take(50))
	assert.Empty(t, s.Take(0))
	assert.Empty(t, s2.Take(10))
}
