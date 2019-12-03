package types

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func makeDefaultInt64s() Int64s {
	return Int64s{1, 2, 3, 4, 5}
}

func TestInt64sReset(t *testing.T) {
	s := &Int64s{1, 2, 3, 4, 5}
	s2 := s.Copy()
	s.Reset()

	assert.Empty(t, s)
	assert.NotEmpty(t, s2)

	s2.Reset()
	assert.Empty(t, s2)
}

func TestInt64sCopy(t *testing.T) {
	s := makeDefaultInt64s()
	s2 := s.Copy()

	assert.Equal(t, s, s2)
	s2 = append(s2, 6, 7, 8)
	assert.NotEqual(t, s, s2)
}

func TestInt64sDiff(t *testing.T) {
	s := makeDefaultInt64s()
	s2 := Int64s{}

	assert.Equal(t, s, s2.Diff(s))
	assert.Equal(t, s, s.Diff(s2))

	s3 := Int64s{1, 3, 5}
	assert.Equal(t, Int64s{2, 4}, s.Diff(s3))
	assert.Equal(t, Int64s{2, 4}, s3.Diff(s))
}

func TestInt64sContains(t *testing.T) {
	s := makeDefaultInt64s()
	assert.True(t, s.Contains(1))
	assert.True(t, s.Contains(1, 2, 3))
	assert.True(t, s.Contains(1, 4, 5))
	assert.False(t, s.Contains(0))
	assert.False(t, s.Contains(0, 5, 1))
	assert.True(t, s.Contains(makeDefaultInt64s()...))
}

func TestInt64sContainsOneOf(t *testing.T) {
	s := makeDefaultInt64s()
	assert.True(t, s.ContainsOneOf(1))
	assert.True(t, s.ContainsOneOf(1, 2, 3))
	assert.True(t, s.ContainsOneOf(1, 4, 5))
	assert.True(t, s.ContainsOneOf(0, 5, 1))
	assert.False(t, s.ContainsOneOf(0))
	assert.False(t, s.ContainsOneOf(0, 6))

}

func TestInt64sEmpty(t *testing.T) {
	s := Int64s{}
	assert.True(t, s.Empty())
	s = makeDefaultInt64s()
	assert.False(t, s.Empty())
}

func TestInt64sEqual(t *testing.T) {
	s := makeDefaultInt64s()
	s2 := makeDefaultInt64s()
	s3 := Int64s{}
	s4 := Int64s{6, 7, 8, 9, 10}
	assert.True(t, s.Equal(s2))
	assert.False(t, s.Equal(s3))
	assert.False(t, s3.Equal(s2))
	assert.False(t, s4.Equal(s2))
	assert.False(t, s2.Equal(s4))
}

func TestInt64sFirst(t *testing.T) {
	s := Int64s{1, 2, 3}
	s2 := Int64s{}
	s3 := Int64s{2, 4, 5}

	{
		v, ok := s.First()
		assert.True(t, ok)
		assert.Equal(t, int64(1), v)
	}

	{
		v, ok := s2.First()
		assert.False(t, ok)
		assert.Equal(t, int64(0), v)
	}

	{
		v, ok := s3.First()
		assert.True(t, ok)
		assert.Equal(t, int64(2), v)
	}
}

func TestInt64sGet(t *testing.T) {
	s := Int64s{1, 2, 3}
	s2 := Int64s{}
	s3 := Int64s{2, 4, 5}

	{
		v, ok := s.Get(0)
		assert.True(t, ok)
		assert.Equal(t, int64(1), v)
	}

	{
		v, ok := s2.Get(6)
		assert.False(t, ok)
		assert.Equal(t, int64(0), v)
	}

	{
		v, ok := s3.Get(2)
		assert.True(t, ok)
		assert.Equal(t, int64(5), v)
	}
}

func TestInt64sIntersect(t *testing.T) {
	s := Int64s{1, 2, 3}
	s2 := Int64s{2, 3, 4, 5}
	s3 := Int64s{}

	assert.Equal(t, Int64s{2, 3}, s.Intersect(s2))
	assert.Equal(t, s3, s3.Intersect(s))
	assert.Equal(t, s3, s.Intersect(s3))
}

func TestInt64sLast(t *testing.T) {
	s := Int64s{1, 2, 3}
	s2 := Int64s{}
	s3 := Int64s{2, 4, 5}

	{
		v, ok := s.Last()
		assert.True(t, ok)
		assert.Equal(t, int64(3), v)
	}

	{
		v, ok := s2.Last()
		assert.False(t, ok)
		assert.Equal(t, int64(0), v)
	}

	{
		v, ok := s3.Last()
		assert.True(t, ok)
		assert.Equal(t, int64(5), v)
	}
}

func TestInt64sTake(t *testing.T) {
	s := Int64s{1, 2, 3}
	s2 := Int64s{}

	assert.Equal(t, s, s.Take(3))
	assert.Equal(t, Int64s{1, 2}, s.Take(2))
	assert.Equal(t, s, s.Take(50))
	assert.Empty(t, s.Take(0))
	assert.Empty(t, s2.Take(10))
}
