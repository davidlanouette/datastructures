package datastructures

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGet(t *testing.T) {

	h := NewMaxIntHeap()

	h.Insert(4)
	h.Insert(2)

	val, err := h.Get()
	assert.NoError(t, err)
	assert.Equal(t, 4, val)

	h.Insert(33)
	h.Insert(32)
	h.Insert(12)
	h.Insert(-1)
	val, err = h.Get()
	assert.NoError(t, err)
	assert.Equal(t, 33, val)
}
