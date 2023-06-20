package set

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNew(t *testing.T) {
	s := New[int]()
	assert.Equal(t, 0, s.Length())
	assert.False(t, s.Contains(0))
	assert.Equal(t, 0, len(s.GetAll()))
}

func TestAdd(t *testing.T) {
	s := New[int]()
	assert.Equal(t, 0, s.Length())
	assert.False(t, s.Contains(0))
	assert.Equal(t, 0, len(s.GetAll()))

	s.Add(0)
	assert.Equal(t, 1, s.Length())
	assert.True(t, s.Contains(0))
	assert.False(t, s.Contains(1))
	assert.Equal(t, 1, len(s.GetAll()))

	s.Add(1)
	assert.Equal(t, 2, s.Length())
	assert.True(t, s.Contains(0))
	assert.True(t, s.Contains(1))
	assert.False(t, s.Contains(2))
	assert.Equal(t, 2, len(s.GetAll()))
}

func TestDelete(t *testing.T) {
	s := New[int]()
	s.Delete(0)
	assert.Equal(t, 0, s.Length())
	assert.False(t, s.Contains(0))
	assert.Equal(t, 0, len(s.GetAll()))

	s.Add(0)
	s.Delete(0)
	assert.Equal(t, 0, s.Length())
	assert.False(t, s.Contains(0))
	assert.False(t, s.Contains(1))
	assert.Equal(t, 0, len(s.GetAll()))

	s.Add(0)
	s.Add(1)
	s.Delete(1)
	assert.Equal(t, 1, s.Length())
	assert.True(t, s.Contains(0))
	assert.False(t, s.Contains(1))
	assert.False(t, s.Contains(2))
	assert.Equal(t, 1, len(s.GetAll()))
}
