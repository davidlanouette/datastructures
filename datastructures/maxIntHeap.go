package datastructures

import (
	"errors"
	"math"
	"strconv"
	"strings"
)

// ErrHeapIsEmpty indicates that you tried to get data out of an empty heap.
var ErrHeapIsEmpty = errors.New("heap is currently empty")

// MaxIntHeap is a datastructure that will allow you to retreive the maximum value quickly.
// This heap is implemented using a slice, instead of dynamic memory.
type MaxIntHeap struct {
	data []int
	size uint
}

//NewMaxIntHeap will create a new MaxIntHeap object for you.
func NewMaxIntHeap() *MaxIntHeap {
	return &MaxIntHeap{
		data: make([]int, 10),
		size: 0,
	}
}

// Get will remove and return the max value from the heap.
func (h *MaxIntHeap) Get() (int, error) {
	if h.IsEmpty() {
		return 0, ErrHeapIsEmpty
	}
	// log.Printf("[Get] before: %v", h.data)
	// log.Print(h.printTree())

	// grab the top most element
	val := h.data[0]

	// swap the last element in the heap with the top element, and decrease the size of the heap by 1.
	h.data[0] = h.data[h.size-1]
	h.size--
	h.data[h.size] = 0 // not strictly needed, but more clearly shows that the element is now blank
	// log.Printf("[Get] middle: %v", h.data)

	h.heapifyDown()

	// log.Printf("[Get] after: %v", h.data)
	// log.Print(h.printTree())

	// log.Printf("[Get] return: %v", val)
	return val, nil
}

// Peek will return the max value from the heap.  Unlike Get, it will leave the heap unchanged.
func (h *MaxIntHeap) Peek() (int, error) {
	if h.IsEmpty() {
		return 0, ErrHeapIsEmpty
	}

	return h.data[0], nil
}

// Insert will add a value to the heap.
func (h *MaxIntHeap) Insert(value int) {
	// log.Printf("[Insert] before: %v", h.data)
	h.ensureCapacity()

	// log.Printf("[ensureCapacity] size: %v, data: %v\n", h.size, h.data)

	h.data[h.size] = value
	h.size++
	// log.Printf("[Insert] middle: %v", h.data)

	h.heapifyUp()

	// log.Printf("[Insert] after: %v", h.data)
}

// Replace will exchange the hightest value with a new value
func (h *MaxIntHeap) Replace(newValue int) error {
	if h.IsEmpty() {
		return ErrHeapIsEmpty
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
	// log.Printf("[ensureCapacity] Before - cap: %v, size: %v", cap(h.data), h.size)
	if cap(h.data) < int(h.size)+1 {
		newSize := (cap(h.data) + 1) * 2
		t := make([]int, newSize)
		copy(t, h.data)
		h.data = t
		// log.Printf("[ensureCapacity] data: %v\n", h.data)
	}
	// log.Printf("[ensureCapacity] After - cap: %v, size: %v", cap(h.data), h.size)
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
	rightChildIdx := h.getRightChildIndex(idx)
	// log.Printf("[heapifyDown] start idx: %d, left (idx: %d, val: %d), right(idx: %d, val: %d), data: %v",
	//	idx, leftChildIdx, h.data[leftChildIdx], rightChildIdx, h.data[rightChildIdx], h.data)
	// log.Print(h.printTree())

	for {
		swapValue := h.data[leftChildIdx]
		swapIdx := leftChildIdx
		if h.data[rightChildIdx] > swapValue {
			swapValue = h.data[rightChildIdx]
			swapIdx = rightChildIdx
		}

		// have we gotten to the end of the array?
		if swapIdx >= h.size {
			break
		}

		if h.data[idx] < swapValue {
			h.swapValues(idx, swapIdx)
		} else {
			// we are done.
			break
		}

		idx = swapIdx
		leftChildIdx = h.getLeftChildIndex(idx)
		rightChildIdx = h.getRightChildIndex(idx)

		// log.Printf("[heapifyDown] idx: %d, left (idx: %d, val: %d), right(idx: %d, val: %d), data: %v",
		//	idx, leftChildIdx, h.data[leftChildIdx], rightChildIdx, h.data[rightChildIdx], h.data)
		// log.Print(h.printTree())

		if idx > h.Size() || leftChildIdx > h.Size() || rightChildIdx > h.Size() {
			break
		}
	}
}

// getLeftChildIndex returns the index in the slice of the left child for the given index
func (h *MaxIntHeap) getLeftChildIndex(idx uint) uint {
	return (idx * 2) + 1
}

// getRightChildIndex returns the index in the slice for the right child of the given index
func (h *MaxIntHeap) getRightChildIndex(idx uint) uint {
	return (idx * 2) + 2
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

// printTree will generate a string representative of the heap, formatted as a 2D tree
func (h *MaxIntHeap) printTree() string {
	var buf strings.Builder

	buf.WriteString("Tree:")

	lvl := 0
	nextLevelStart := 0

	// calculate the "cell size" for each number
	// this is the number of count of numbers on the last row [#levels (log(h.size)+1), times 2] times the number of sapces in each cell [five, in our case]
	numLevels := int(math.Log2(float64(h.size))) + 1
	cellWidth := int(math.Pow(float64(numLevels), float64(2))) * 8

	for i, v := range h.data {
		if uint(i) >= h.size {
			break
		}

		if i == nextLevelStart {
			lvl++
			nextLevelStart = int(math.Pow(2, float64(lvl))) - 1

			cellWidth /= 2
			buf.WriteString("\n")
		}

		buf.WriteString(centerJustify(strconv.Itoa(v), cellWidth))
	}
	buf.WriteString("\n")

	return buf.String()
}

func centerJustify(val string, size int) string {
	buf := ""
	spaces := (size - len(val)) / 2

	for i := 0; i < spaces; i++ {
		buf += " "
	}

	buf += val

	for len(buf) < size {
		buf += " "
	}

	return buf
}
