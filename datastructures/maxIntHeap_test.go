package datastructures

import (
	"log"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGet(t *testing.T) {

	// t.SkipNow()

	h := NewMaxIntHeap()

	h.Insert(4)
	assert.Equal(t, uint(1), h.Size())

	h.Insert(2)
	assert.Equal(t, uint(2), h.Size())

	val, err := h.Get()
	assert.NoError(t, err)
	assert.Equal(t, 4, val)
	assert.Equal(t, uint(1), h.Size())

	h.Insert(1)
	assert.Equal(t, uint(2), h.Size())

	h.Insert(33)
	assert.Equal(t, uint(3), h.Size())

	h.Insert(32)
	assert.Equal(t, uint(4), h.Size())

	h.Insert(12)
	assert.Equal(t, uint(5), h.Size())

	h.Insert(-1)
	assert.Equal(t, uint(6), h.Size())

	val, err = h.Get()
	assert.NoError(t, err)
	assert.Equal(t, 33, val)
	assert.Equal(t, uint(5), h.Size())

	val, err = h.Get()
	assert.NoError(t, err)
	assert.Equal(t, 32, val)
	assert.Equal(t, uint(4), h.Size())

	val, err = h.Get()
	assert.NoError(t, err)
	assert.Equal(t, 12, val)
	assert.Equal(t, uint(3), h.Size())

	val, err = h.Get()
	assert.NoError(t, err)
	assert.Equal(t, 2, val)
	assert.Equal(t, uint(2), h.Size())

	val, err = h.Get()
	assert.NoError(t, err)
	assert.Equal(t, 1, val)
	assert.Equal(t, uint(1), h.Size())

	val, err = h.Get()
	assert.NoError(t, err)
	assert.Equal(t, -1, val)
	assert.Equal(t, uint(0), h.Size())

	val, err = h.Get()
	assert.Error(t, err)
}

func TestGetOnSmallHeap(t *testing.T) {
	h := NewMaxIntHeap()

	h.Insert(4)
	assert.Equal(t, uint(1), h.Size())

	val, err := h.Get()
	assert.NoError(t, err)
	assert.Equal(t, 4, val)
	assert.Equal(t, uint(0), h.Size())
}

func TestGetOnEmptyHeap(t *testing.T) {
	h := NewMaxIntHeap()

	_, err := h.Get()
	assert.Error(t, err)
	assert.Equal(t, uint(0), h.Size())
}

func TestPrintTree(t *testing.T) {
	t.SkipNow()
	h := NewMaxIntHeap()

	h.Insert(4)
	h.Insert(33)
	h.Insert(6)
	h.Insert(100)
	h.Insert(2)
	h.Insert(3)
	h.Insert(40)
	h.Insert(22)
	h.Insert(88)
	h.Insert(32)
	h.Insert(90)
	h.Insert(72)
	h.Insert(73)
	h.Insert(74)
	h.Insert(75)
	h.Insert(66)
	h.Insert(6)
	h.Insert(99)

	log.Printf("%v\n", h.data)
	log.Print(h.printTree())

}

func TestPrintTree2(t *testing.T) {
	t.SkipNow()
	h := NewMaxIntHeap()

	for i := 0; i < 200; i += 2 {
		h.Insert(i)
	}

	log.Printf("%v\n", h.data)
	log.Print(h.printTree())

}
