package datastructures

import "errors"

// MaxIntHeap is a datastructure that will allow you to retreive the maximum value quickly.
// This heap is implemented using a slice, instead of dynamic memory.
type MaxIntHeap struct {
	data []int
	size uint
}

//NewMaxIntHeap will create a new MaxIntHeap object for you.
func NewMaxIntHeap() *MaxIntHeap {
	return &MaxIntHeap{
		data: make([]int, 100, 100),
		size: 0,
	}
}

// Get will remove and return the max value from the heap.
func (h *MaxIntHeap) Get() (int, error) {
	if h.IsEmpty() {
		return 0, errors.New("heap is currently empty")
	}

	val := h.data[0]

	h.data[0] = h.data[h.size-1]
	h.data[h.size-1] = 0 // DEBUG: blank out the value we just removed.
	h.size--

	h.heapifyDown()

	return val, nil
}

// Peek will return the max value from the heap.  Unlike Get, it will leave the heap unchanged.
func (h *MaxIntHeap) Peek() (int, error) {
	if h.IsEmpty() {
		return 0, errors.New("heap is currently empty")
	}

	return h.data[0], nil
}

// Insert will add a value to the heap.
func (h *MaxIntHeap) Insert(value int) {
	h.ensureCapacity()

	h.data[h.size] = value
	h.size++

	h.heapifyUp()

}

// Replace will exchange the hightest value with a new value
func (h *MaxIntHeap) Replace(newValue int) error {
	if h.IsEmpty() {
		return errors.New("heap is currently empty")
	}

	return nil
}

// Size returns the number of elements in the heap
func (h *MaxIntHeap) Size() uint {
	return h.size
}

// IsEmpty will return wether there are any values in the heap.
func (h *MaxIntHeap) IsEmpty() bool {
	return h.size == 0
}

// ensureCapacity will check that the heap has enough space to add a new element.
// if there isn't space (ie, the heap is full), the heap will be doubled in size.
func (h *MaxIntHeap) ensureCapacity() {
	if cap(h.data) <= int(h.size) {
		t := make([]int, len(h.data), (cap(h.data)+1)*2)
		copy(t, h.data)
		h.data = t
	}
}

// getMinIndex returns the index of the minimum value in the heap
func (h *MaxIntHeap) getMinIndex() uint {
	return 0
}

// heapifyUp will move the last inserted value in the heap to it's "correct" position by bubbling it up until it's parent is greater it.
func (h *MaxIntHeap) heapifyUp() {
	idx := h.size - 1 // last index in the heap
	parentIdx := h.getParentIndex(idx)
	for h.data[idx] > h.data[parentIdx] {
		h.swapValues(idx, parentIdx)
		idx = parentIdx
		parentIdx = h.getParentIndex(idx)
	}
}

// heapifyDown will move the top most value in the heap to it's "correct" position by bubbling it down until it's children are less than it.
func (h *MaxIntHeap) heapifyDown() {
	idx := uint(0) // top index
	leftChildIdx := h.getLeftChildIndex(idx)
	for h.data[idx] < h.data[leftChildIdx] {
		h.swapValues(idx, leftChildIdx)
		idx = leftChildIdx
		leftChildIdx = h.getLeftChildIndex(idx)
	}
}

// getLeftChildIndex returns the index in the slice of the left child for the given index
func (h *MaxIntHeap) getLeftChildIndex(idx uint) uint {
	return idx * 2
}

// getRightChildIndex returns the index in the slice for the right child of the given index
func (h *MaxIntHeap) getRightChildIndex(idx uint) uint {
	return idx*2 + 1
}

// getParentIndex returns the index in the slice for the parent of the given index
func (h *MaxIntHeap) getParentIndex(idx uint) uint {
	if idx == 0 {
		return 0
	}
	return idx / 2
}

func (h *MaxIntHeap) swapValues(firstIdx, secondIdx uint) {
	val := h.data[firstIdx]
	h.data[firstIdx] = h.data[secondIdx]
	h.data[secondIdx] = val
}
