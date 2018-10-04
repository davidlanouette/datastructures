package stack

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStackEmptyStack(t *testing.T) {
	s := NewStack()

	assert.Equal(t, 0, s.Size())
}

func TestStackOneElement(t *testing.T) {
	s := NewStack()

	assert.Equal(t, 0, s.Size())

	s.Push(1)
	assert.Equal(t, 1, s.Size())

	item, err := s.Pop()
	if err != nil {
		assert.Error(t, err)
	}
	assert.Equal(t, 0, s.Size())
	assert.Equal(t, 1, item)
}

func TestStackManyElements(t *testing.T) {
	s := NewStack()

	assert.Equal(t, 0, s.Size())

	s.Push("abc")
	s.Push("def")
	assert.Equal(t, 2, s.Size())
	item, err := s.Pop()
	if err != nil {
		assert.Error(t, err)
	}
	assert.Equal(t, 1, s.Size())
	assert.Equal(t, "def", item)

	s.Push("ghi")
	s.Push("jkl")
	s.Push("mno")

	assert.Equal(t, 4, s.Size())

	item, err = s.Pop()
	if err != nil {
		assert.Error(t, err)
	}
	assert.Equal(t, 3, s.Size())
	assert.Equal(t, "mno", item)
}
