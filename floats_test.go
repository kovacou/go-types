package types

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func makeDefaultFloats() Floats {
	return Floats{1, 2, 3, 4, 5}
}

func TestFloatsReset(t *testing.T) {
	s := &Floats{1, 2, 3, 4, 5}
	s2 := s.Copy()
	s.Reset()

	assert.Empty(t, s)
	assert.NotEmpty(t, s2)

	s2.Reset()
	assert.Empty(t, s2)
}

func TestFloatsCopy(t *testing.T) {
	s := makeDefaultFloats()
	s2 := s.Copy()

	assert.Equal(t, s, s2)
	s2 = append(s2, 6, 7, 8)
	assert.NotEqual(t, s, s2)
}

func TestFloatsDiff(t *testing.T) {
	s := makeDefaultFloats()
	s2 := Floats{}

	assert.Equal(t, s, s2.Diff(s))
	assert.Equal(t, s, s.Diff(s2))

	s3 := Floats{1, 3, 5}
	assert.Equal(t, Floats{2, 4}, s.Diff(s3))
	assert.Equal(t, Floats{2, 4}, s3.Diff(s))
}

func TestFloatsContains(t *testing.T) {
	s := makeDefaultFloats()
	assert.True(t, s.Contains(1))
	assert.True(t, s.Contains(1, 2, 3))
	assert.True(t, s.Contains(1, 4, 5))
	assert.False(t, s.Contains(0))
	assert.False(t, s.Contains(0, 5, 1))
	assert.True(t, s.Contains(makeDefaultFloats()...))
}

func TestFloatsContainsOneOf(t *testing.T) {
	s := makeDefaultFloats()
	assert.True(t, s.ContainsOneOf(1))
	assert.True(t, s.ContainsOneOf(1, 2, 3))
	assert.True(t, s.ContainsOneOf(1, 4, 5))
	assert.True(t, s.ContainsOneOf(0, 5, 1))
	assert.False(t, s.ContainsOneOf(0))
	assert.False(t, s.ContainsOneOf(0, 6))

}

func TestFloatsEmpty(t *testing.T) {
	s := Floats{}
	assert.True(t, s.Empty())
	s = makeDefaultFloats()
	assert.False(t, s.Empty())
}

func TestFloatsEqual(t *testing.T) {
	s := makeDefaultFloats()
	s2 := makeDefaultFloats()
	s3 := Floats{}
	s4 := Floats{6, 7, 8, 9, 10}
	assert.True(t, s.Equal(s2))
	assert.False(t, s.Equal(s3))
	assert.False(t, s3.Equal(s2))
	assert.False(t, s4.Equal(s2))
	assert.False(t, s2.Equal(s4))
}

func TestFloatsFirst(t *testing.T) {
	s := Floats{1, 2, 3}
	s2 := Floats{}
	s3 := Floats{2, 4, 5}

	{
		v, ok := s.First()
		assert.True(t, ok)
		assert.Equal(t, float64(1), v)
	}

	{
		v, ok := s2.First()
		assert.False(t, ok)
		assert.Equal(t, float64(0), v)
	}

	{
		v, ok := s3.First()
		assert.True(t, ok)
		assert.Equal(t, float64(2), v)
	}
}

func TestFloatsGet(t *testing.T) {
	s := Floats{1, 2, 3}
	s2 := Floats{}
	s3 := Floats{2, 4, 5}

	{
		v, ok := s.Get(0)
		assert.True(t, ok)
		assert.Equal(t, float64(1), v)
	}

	{
		v, ok := s2.Get(6)
		assert.False(t, ok)
		assert.Equal(t, float64(0), v)
	}

	{
		v, ok := s3.Get(2)
		assert.True(t, ok)
		assert.Equal(t, float64(5), v)
	}
}

func TestFloatsIntersect(t *testing.T) {
	s := Floats{1, 2, 3}
	s2 := Floats{2, 3, 4, 5}
	s3 := Floats{}

	assert.Equal(t, Floats{2, 3}, s.Intersect(s2))
	assert.Equal(t, s3, s3.Intersect(s))
	assert.Equal(t, s3, s.Intersect(s3))
}

func TestFloatsLast(t *testing.T) {
	s := Floats{1, 2, 3}
	s2 := Floats{}
	s3 := Floats{2, 4, 5}

	{
		v, ok := s.Last()
		assert.True(t, ok)
		assert.Equal(t, float64(3), v)
	}

	{
		v, ok := s2.Last()
		assert.False(t, ok)
		assert.Equal(t, float64(0), v)
	}

	{
		v, ok := s3.Last()
		assert.True(t, ok)
		assert.Equal(t, float64(5), v)
	}
}

func TestFloatsTake(t *testing.T) {
	s := Floats{1, 2, 3}
	s2 := Floats{}

	assert.Equal(t, s, s.Take(3))
	assert.Equal(t, Floats{1, 2}, s.Take(2))
	assert.Equal(t, s, s.Take(50))
	assert.Empty(t, s.Take(0))
	assert.Empty(t, s2.Take(10))
}
