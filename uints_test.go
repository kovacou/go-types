package types

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func makeDefaultUints() Uints {
	return Uints{1, 2, 3, 4, 5}
}

func TestUintsReset(t *testing.T) {
	s := &Uints{1, 2, 3, 4, 5}
	s2 := s.Copy()
	s.Reset()

	assert.Empty(t, s)
	assert.NotEmpty(t, s2)

	s2.Reset()
	assert.Empty(t, s2)
}

func TestUintsCopy(t *testing.T) {
	s := makeDefaultUints()
	s2 := s.Copy()

	assert.Equal(t, s, s2)
	s2 = append(s2, 6, 7, 8)
	assert.NotEqual(t, s, s2)
}

func TestUintsDiff(t *testing.T) {
	s := makeDefaultUints()
	s2 := Uints{}

	assert.Equal(t, s, s2.Diff(s))
	assert.Equal(t, s, s.Diff(s2))

	s3 := Uints{1, 3, 5}
	assert.Equal(t, Uints{2, 4}, s.Diff(s3))
	assert.Equal(t, Uints{2, 4}, s3.Diff(s))
}

func TestUintsContains(t *testing.T) {
	s := makeDefaultUints()
	assert.True(t, s.Contains(1))
	assert.True(t, s.Contains(1, 2, 3))
	assert.True(t, s.Contains(1, 4, 5))
	assert.False(t, s.Contains(0))
	assert.False(t, s.Contains(0, 5, 1))
	assert.True(t, s.Contains(makeDefaultUints()...))
}

func TestUintsContainsOneOf(t *testing.T) {
	s := makeDefaultUints()
	assert.True(t, s.ContainsOneOf(1))
	assert.True(t, s.ContainsOneOf(1, 2, 3))
	assert.True(t, s.ContainsOneOf(1, 4, 5))
	assert.True(t, s.ContainsOneOf(0, 5, 1))
	assert.False(t, s.ContainsOneOf(0))
	assert.False(t, s.ContainsOneOf(0, 6))

}

func TestUintsEmpty(t *testing.T) {
	s := Uints{}
	assert.True(t, s.Empty())
	s = makeDefaultUints()
	assert.False(t, s.Empty())
}

func TestUintsEqual(t *testing.T) {
	s := makeDefaultUints()
	s2 := makeDefaultUints()
	s3 := Uints{}
	s4 := Uints{6, 7, 8, 9, 10}
	assert.True(t, s.Equal(s2))
	assert.False(t, s.Equal(s3))
	assert.False(t, s3.Equal(s2))
	assert.False(t, s4.Equal(s2))
	assert.False(t, s2.Equal(s4))
}

func TestUintsFirst(t *testing.T) {
	s := Uints{1, 2, 3}
	s2 := Uints{}
	s3 := Uints{2, 4, 5}

	{
		v, ok := s.First()
		assert.True(t, ok)
		assert.Equal(t, uint(1), v)
	}

	{
		v, ok := s2.First()
		assert.False(t, ok)
		assert.Equal(t, uint(0), v)
	}

	{
		v, ok := s3.First()
		assert.True(t, ok)
		assert.Equal(t, uint(2), v)
	}
}

func TestUintsGet(t *testing.T) {
	s := Uints{1, 2, 3}
	s2 := Uints{}
	s3 := Uints{2, 4, 5}

	{
		v, ok := s.Get(0)
		assert.True(t, ok)
		assert.Equal(t, uint(1), v)
	}

	{
		v, ok := s2.Get(6)
		assert.False(t, ok)
		assert.Equal(t, uint(0), v)
	}

	{
		v, ok := s3.Get(2)
		assert.True(t, ok)
		assert.Equal(t, uint(5), v)
	}
}

func TestUintsIntersect(t *testing.T) {
	s := Uints{1, 2, 3}
	s2 := Uints{2, 3, 4, 5}
	s3 := Uints{}

	assert.Equal(t, Uints{2, 3}, s.Intersect(s2))
	assert.Equal(t, s3, s3.Intersect(s))
	assert.Equal(t, s3, s.Intersect(s3))
}

func TestUintsLast(t *testing.T) {
	s := Uints{1, 2, 3}
	s2 := Uints{}
	s3 := Uints{2, 4, 5}

	{
		v, ok := s.Last()
		assert.True(t, ok)
		assert.Equal(t, uint(3), v)
	}

	{
		v, ok := s2.Last()
		assert.False(t, ok)
		assert.Equal(t, uint(0), v)
	}

	{
		v, ok := s3.Last()
		assert.True(t, ok)
		assert.Equal(t, uint(5), v)
	}
}

func TestUintsTake(t *testing.T) {
	s := Uints{1, 2, 3}
	s2 := Uints{}

	assert.Equal(t, s, s.Take(3))
	assert.Equal(t, Uints{1, 2}, s.Take(2))
	assert.Equal(t, s, s.Take(50))
	assert.Empty(t, s.Take(0))
	assert.Empty(t, s2.Take(10))
}
