package types

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func makeDefaultBytes() Bytes {
	return Bytes{1, 2, 3, 4, 5}
}

func TestBytesReset(t *testing.T) {
	s := &Bytes{1, 2, 3, 4, 5}
	s2 := s.Copy()
	s.Reset()

	assert.Empty(t, s)
	assert.NotEmpty(t, s2)

	s2.Reset()
	assert.Empty(t, s2)
}

func TestBytesCopy(t *testing.T) {
	s := makeDefaultBytes()
	s2 := s.Copy()

	assert.Equal(t, s, s2)
	s2 = append(s2, 6, 7, 8)
	assert.NotEqual(t, s, s2)
}

func TestBytesDiff(t *testing.T) {
	s := makeDefaultBytes()
	s2 := Bytes{}

	assert.Equal(t, s, s2.Diff(s))
	assert.Equal(t, s, s.Diff(s2))

	s3 := Bytes{1, 3, 5}
	assert.Equal(t, Bytes{2, 4}, s.Diff(s3))
	assert.Equal(t, Bytes{2, 4}, s3.Diff(s))
}

func TestBytesContains(t *testing.T) {
	s := makeDefaultBytes()
	assert.True(t, s.Contains(1))
	assert.True(t, s.Contains(1, 2, 3))
	assert.True(t, s.Contains(1, 4, 5))
	assert.False(t, s.Contains(0))
	assert.False(t, s.Contains(0, 5, 1))
	assert.True(t, s.Contains(makeDefaultBytes()...))
}

func TestBytesContainsOneOf(t *testing.T) {
	s := makeDefaultBytes()
	assert.True(t, s.ContainsOneOf(1))
	assert.True(t, s.ContainsOneOf(1, 2, 3))
	assert.True(t, s.ContainsOneOf(1, 4, 5))
	assert.True(t, s.ContainsOneOf(0, 5, 1))
	assert.False(t, s.ContainsOneOf(0))
	assert.False(t, s.ContainsOneOf(0, 6))

}

func TestBytesEmpty(t *testing.T) {
	s := Bytes{}
	assert.True(t, s.Empty())
	s = makeDefaultBytes()
	assert.False(t, s.Empty())
}

func TestBytesEqual(t *testing.T) {
	s := makeDefaultBytes()
	s2 := makeDefaultBytes()
	s3 := Bytes{}
	s4 := Bytes{6, 7, 8, 9, 10}
	assert.True(t, s.Equal(s2))
	assert.False(t, s.Equal(s3))
	assert.False(t, s3.Equal(s2))
	assert.False(t, s4.Equal(s2))
	assert.False(t, s2.Equal(s4))
}

func TestBytesFirst(t *testing.T) {
	s := Bytes{1, 2, 3}
	s2 := Bytes{}
	s3 := Bytes{2, 4, 5}

	{
		v, ok := s.First()
		assert.True(t, ok)
		assert.Equal(t, byte(1), v)
	}

	{
		v, ok := s2.First()
		assert.False(t, ok)
		assert.Equal(t, byte(0), v)
	}

	{
		v, ok := s3.First()
		assert.True(t, ok)
		assert.Equal(t, byte(2), v)
	}
}

func TestBytesGet(t *testing.T) {
	s := Bytes{1, 2, 3}
	s2 := Bytes{}
	s3 := Bytes{2, 4, 5}

	{
		v, ok := s.Get(0)
		assert.True(t, ok)
		assert.Equal(t, byte(1), v)
	}

	{
		v, ok := s2.Get(6)
		assert.False(t, ok)
		assert.Equal(t, byte(0), v)
	}

	{
		v, ok := s3.Get(2)
		assert.True(t, ok)
		assert.Equal(t, byte(5), v)
	}
}

func TestBytesIntersect(t *testing.T) {
	s := Bytes{1, 2, 3}
	s2 := Bytes{2, 3, 4, 5}
	s3 := Bytes{}

	assert.Equal(t, Bytes{2, 3}, s.Intersect(s2))
	assert.Equal(t, s3, s3.Intersect(s))
	assert.Equal(t, s3, s.Intersect(s3))
}

func TestBytesLast(t *testing.T) {
	s := Bytes{1, 2, 3}
	s2 := Bytes{}
	s3 := Bytes{2, 4, 5}

	{
		v, ok := s.Last()
		assert.True(t, ok)
		assert.Equal(t, byte(3), v)
	}

	{
		v, ok := s2.Last()
		assert.False(t, ok)
		assert.Equal(t, byte(0), v)
	}

	{
		v, ok := s3.Last()
		assert.True(t, ok)
		assert.Equal(t, byte(5), v)
	}
}

func TestBytesTake(t *testing.T) {
	s := Bytes{1, 2, 3}
	s2 := Bytes{}

	assert.Equal(t, s, s.Take(3))
	assert.Equal(t, Bytes{1, 2}, s.Take(2))
	assert.Equal(t, s, s.Take(50))
	assert.Empty(t, s.Take(0))
	assert.Empty(t, s2.Take(10))
}
