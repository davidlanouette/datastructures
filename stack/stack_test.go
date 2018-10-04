package stack

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStackEmptyStack(t *testing.T) {
	s := NewStack()

	assert.Equal(t, 0, s.Size())

	e, err := s.Pop()
	if err != nil {
		assert.Error(t, err)
	}
	assert.Nil(t, e)
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

func TestStackFewElements(t *testing.T) {
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

func TestStackManyElements(t *testing.T) {
	s := NewStack()

	for i := 0; i < 1000000; i++ {
		s.Push(i)
	}

	assert.Equal(t, 1000000, s.Size())
}

func TestStackPeekWithOneElement(t *testing.T) {
	s := NewStack()

	s.Push("ABC")
	e1 := s.Peek()
	assert.Equal(t, "ABC", e1)

	e2, err := s.Pop()
	if err != nil {
		assert.Error(t, err)
	}
	assert.Equal(t, "ABC", e2)
}

func TestStackPeekWithFewElements(t *testing.T) {
	s := NewStack()

	s.Push("abc")
	s.Push("def")
	e1 := s.Peek()
	assert.Equal(t, "def", e1)

	e1, err := s.Pop()
	if err != nil {
		assert.Error(t, err)
	}
	assert.Equal(t, "def", e1)

	s.Push("ghi")
	s.Push("jkl")
	s.Push("mno")
	e1 = s.Peek()
	assert.Equal(t, "mno", e1)

	e1, err = s.Pop()
	if err != nil {
		assert.Error(t, err)
	}
	assert.Equal(t, "mno", e1)
}

func TestStackWithDifferentTypes(t *testing.T) {
	s := NewStack()

	s.Push(1)
	s.Push("abc")
	s.Push(3.1415)

	e, err := s.Pop()
	if err != nil {
		assert.Error(t, err)
	}
	assert.Equal(t, 3.1415, e)

	e, err = s.Pop()
	if err != nil {
		assert.Error(t, err)
	}
	assert.Equal(t, "abc", e)

	e, err = s.Pop()
	if err != nil {
		assert.Error(t, err)
	}
	assert.Equal(t, 1, e)

	// there should be no more elements
	e, err = s.Pop()
	if err != nil {
		assert.Error(t, err)
	}
	assert.Nil(t, e)
}
