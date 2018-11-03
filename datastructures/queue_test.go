package datastructures

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestQueueEmptyQueue(t *testing.T) {
	q := NewQueue()

	assert.Equal(t, 0, q.Size())

	e, err := q.Remove()
	assert.Equal(t, err, EmptyQueueError)
	assert.Nil(t, e)
}

func TestQueueOneElement(t *testing.T) {
	q := NewQueue()

	assert.Equal(t, 0, q.Size())

	q.Add(1)
	assert.Equal(t, 1, q.Size())

	item, err := q.Remove()
	assert.Nil(t, err)
	assert.Equal(t, 0, q.Size())
	assert.Equal(t, 1, item)
}
func TestQueueFewElements(t *testing.T) {
	q := NewQueue()

	assert.Equal(t, 0, q.Size())

	q.Add("abc")
	q.Add("def")

	assert.Equal(t, 2, q.Size())
	item, err := q.Remove()
	assert.Nil(t, err)
	assert.Equal(t, 1, q.Size())
	assert.Equal(t, "abc", item)

	q.Add("ghi")
	q.Add("jkl")
	q.Add("mno")

	assert.Equal(t, 4, q.Size())

	item, err = q.Remove()
	assert.Nil(t, err)
	assert.Equal(t, 3, q.Size())
	assert.Equal(t, "def", item)

	item, err = q.Remove()
	assert.Nil(t, err)
	assert.Equal(t, 2, q.Size())
	assert.Equal(t, "ghi", item)

	item, err = q.Remove()
	assert.Nil(t, err)
	assert.Equal(t, 1, q.Size())
	assert.Equal(t, "jkl", item)

	item, err = q.Remove()
	assert.Nil(t, err)
	assert.Equal(t, 0, q.Size())
	assert.Equal(t, "mno", item)

	item, err = q.Remove()
	assert.Equal(t, EmptyQueueError, err)
	assert.Equal(t, 0, q.Size())
	assert.Nil(t, item)
}

func TestQueueManyElements(t *testing.T) {
	q := NewQueue()

	for i := 0; i < 1000000; i++ {
		q.Add(i)
	}

	assert.Equal(t, 1000000, q.Size())

	for i := 0; i < 1000000; i++ {
		_, err := q.Remove()
		assert.Nil(t, err)
	}
	assert.Equal(t, 0, q.Size())
}

func TestQueuePeekWithOneElement(t *testing.T) {
	q := NewQueue()

	q.Add("ABC")
	assert.Equal(t, 1, q.Size())
	e1, err := q.Peek()
	assert.Nil(t, err)
	assert.Equal(t, 1, q.Size())
	assert.Equal(t, "ABC", e1)

	e2, err := q.Remove()
	assert.Nil(t, err)
	assert.Equal(t, "ABC", e2)
	assert.Equal(t, 0, q.Size())
}

func TestQueuePeekWithFewElements(t *testing.T) {
	q := NewQueue()

	q.Add("abc")
	q.Add("def")

	e1, err := q.Peek()
	assert.Nil(t, err)
	assert.Equal(t, 2, q.Size())
	assert.Equal(t, "abc", e1)
	assert.Equal(t, 2, q.Size())

	e1, err = q.Remove()
	assert.Nil(t, err)
	assert.Equal(t, "abc", e1)
	assert.Equal(t, 1, q.Size())

	q.Add("ghi")
	q.Add("jkl")
	q.Add("mno")
	assert.Equal(t, 4, q.Size())
	e1, err = q.Peek()
	assert.Nil(t, err)
	assert.Equal(t, "def", e1)
	assert.Equal(t, 4, q.Size())

	e1, err = q.Remove()
	assert.Nil(t, err)
	assert.Equal(t, "def", e1)
	assert.Equal(t, 3, q.Size())
}

func TestQueueWithDifferentTypes(t *testing.T) {
	q := NewQueue()

	q.Add(1)
	q.Add("abc")
	q.Add(3.1415)

	e, err := q.Remove()
	assert.Nil(t, err)
	assert.Equal(t, 1, e)
	assert.Equal(t, 2, q.Size())

	e, err = q.Remove()
	assert.Nil(t, err)
	assert.Equal(t, "abc", e)
	assert.Equal(t, 1, q.Size())

	e, err = q.Remove()
	assert.Nil(t, err)
	assert.Equal(t, 3.1415, e)
	assert.Equal(t, 0, q.Size())

	// there should be no more elements
	e, err = q.Remove()
	assert.Equal(t, err, EmptyQueueError)
	assert.Nil(t, e)
	assert.Equal(t, 0, q.Size())
}
