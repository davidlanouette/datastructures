package datastructures

// Heap is an interface that represents a min or max heap
type Heap interface {
	Get() (interface{}, error)
	Peek() (interface{}, error)
	Insert(value interface{})
	Replace(newValue interface{}) error
	Size() int
	IsEmpty() bool
}
