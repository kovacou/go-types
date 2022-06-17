package types

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func makeDefaultUint64s() Uint64s {
	return Uint64s{1, 2, 3, 4, 5}
}

func TestUint64s_Add(t *testing.T) {
	s := Uint64s{}
	s.Add(1, 2, 3)

	assert.Len(t, s, 3)
	assert.Equal(t, s, Uint64s{1, 2, 3})
}

func TestUint64sReset(t *testing.T) {
	s := &Uint64s{1, 2, 3, 4, 5}
	s2 := s.Copy()
	s.Reset()

	assert.Empty(t, s)
	assert.NotEmpty(t, s2)

	s2.Reset()
	assert.Empty(t, s2)
}

func TestUint64sCopy(t *testing.T) {
	s := makeDefaultUint64s()
	s2 := s.Copy()

	assert.Equal(t, s, s2)
	s2 = append(s2, 6, 7, 8)
	assert.NotEqual(t, s, s2)
}

func TestUint64sDiff(t *testing.T) {
	s := makeDefaultUint64s()
	s2 := Uint64s{}

	assert.Equal(t, s, s2.Diff(s))
	assert.Equal(t, s, s.Diff(s2))

	s3 := Uint64s{1, 3, 5}
	assert.Equal(t, Uint64s{2, 4}, s.Diff(s3))
	assert.Equal(t, Uint64s{2, 4}, s3.Diff(s))
}

func TestUint64sContains(t *testing.T) {
	s := makeDefaultUint64s()
	assert.True(t, s.Contains(1))
	assert.True(t, s.Contains(1, 2, 3))
	assert.True(t, s.Contains(1, 4, 5))
	assert.False(t, s.Contains(0))
	assert.False(t, s.Contains(0, 5, 1))
	assert.True(t, s.Contains(makeDefaultUint64s()...))
}

func TestUint64sContainsOneOf(t *testing.T) {
	s := makeDefaultUint64s()
	assert.True(t, s.ContainsOneOf(1))
	assert.True(t, s.ContainsOneOf(1, 2, 3))
	assert.True(t, s.ContainsOneOf(1, 4, 5))
	assert.True(t, s.ContainsOneOf(0, 5, 1))
	assert.False(t, s.ContainsOneOf(0))
	assert.False(t, s.ContainsOneOf(0, 6))

}

func TestUint64sEmpty(t *testing.T) {
	s := Uint64s{}
	assert.True(t, s.Empty())
	s = makeDefaultUint64s()
	assert.False(t, s.Empty())
}

func TestUint64sEqual(t *testing.T) {
	s := makeDefaultUint64s()
	s2 := makeDefaultUint64s()
	s3 := Uint64s{}
	s4 := Uint64s{6, 7, 8, 9, 10}
	assert.True(t, s.Equal(s2))
	assert.False(t, s.Equal(s3))
	assert.False(t, s3.Equal(s2))
	assert.False(t, s4.Equal(s2))
	assert.False(t, s2.Equal(s4))
}

func TestUint64sExcludes(t *testing.T) {
	s := makeDefaultUint64s()
	s2 := makeDefaultUint64s()
	s3 := Uint64s{3, 4, 5}
	s4 := Uint64s{6, 7, 8, 9, 10}

	assert.Empty(t, s.Excludes(s2))
	assert.Equal(t, s4, s4.Excludes(s))
	assert.Equal(t, Uint64s{1, 2}, s.Excludes(s3))
}

func TestUint64s_Filter(t *testing.T) {
	s := makeDefaultUint64s()

	assert.Equal(t, Uint64s{4, 5}, s.Filter(func(v uint64) bool {
		return v > 3
	}))

	assert.Equal(t, Uint64s{1, 2}, s.Filter(func(v uint64) bool {
		return v < 3
	}))

	assert.Empty(t, s.Filter(func(v uint64) bool {
		return v == 0
	}))
}

func TestUint64sFirst(t *testing.T) {
	s := Uint64s{1, 2, 3}
	s2 := Uint64s{}
	s3 := Uint64s{2, 4, 5}

	{
		v, ok := s.First()
		assert.True(t, ok)
		assert.Equal(t, uint64(1), v)
	}

	{
		v, ok := s2.First()
		assert.False(t, ok)
		assert.Equal(t, uint64(0), v)
	}

	{
		v, ok := s3.First()
		assert.True(t, ok)
		assert.Equal(t, uint64(2), v)
	}
}

func TestUint64sGet(t *testing.T) {
	s := Uint64s{1, 2, 3}
	s2 := Uint64s{}
	s3 := Uint64s{2, 4, 5}

	{
		v, ok := s.Get(0)
		assert.True(t, ok)
		assert.Equal(t, uint64(1), v)
	}

	{
		v, ok := s2.Get(6)
		assert.False(t, ok)
		assert.Equal(t, uint64(0), v)
	}

	{
		v, ok := s3.Get(2)
		assert.True(t, ok)
		assert.Equal(t, uint64(5), v)
	}
}

func TestUint64sIntersect(t *testing.T) {
	s := Uint64s{1, 2, 3}
	s2 := Uint64s{2, 3, 4, 5}
	s3 := Uint64s{}

	assert.Equal(t, Uint64s{2, 3}, s.Intersect(s2))
	assert.Equal(t, s3, s3.Intersect(s))
	assert.Equal(t, s3, s.Intersect(s3))
}

func TestUint64sLast(t *testing.T) {
	s := Uint64s{1, 2, 3}
	s2 := Uint64s{}
	s3 := Uint64s{2, 4, 5}

	{
		v, ok := s.Last()
		assert.True(t, ok)
		assert.Equal(t, uint64(3), v)
	}

	{
		v, ok := s2.Last()
		assert.False(t, ok)
		assert.Equal(t, uint64(0), v)
	}

	{
		v, ok := s3.Last()
		assert.True(t, ok)
		assert.Equal(t, uint64(5), v)
	}
}

func TestUint64sTake(t *testing.T) {
	s := Uint64s{1, 2, 3}
	s2 := Uint64s{}

	assert.Equal(t, s, s.Take(3))
	assert.Equal(t, Uint64s{1, 2}, s.Take(2))
	assert.Equal(t, s, s.Take(50))
	assert.Empty(t, s.Take(0))
	assert.Empty(t, s2.Take(10))
}
