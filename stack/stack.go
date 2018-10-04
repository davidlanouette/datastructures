package stack

import "fmt"

// Stack a last-in-first-out stack implementation
type Stack struct {
	head *stackElement
	size int
}

type stackElement struct {
	element     interface{}
	nextElement *stackElement
}

// Should the traversal ethod actually print the contents of the stack?
const printTraversal = false

// NewStack Create a new stack instance
func NewStack() *Stack {
	return &Stack{}
}

// Push add a new value to the stack
func (s *Stack) Push(val interface{}) {
	s.traverse("--- before push")
	element := &stackElement{val, s.head}
	s.head = element
	s.size++
	s.traverse("--- after push")
	return
}

// Pop removes the most recently added value from the stack, and removes it from the top of the stack.
func (s *Stack) Pop() (interface{}, error) {
	s.traverse("--- before pop")

	var element = s.head

	// short circuit for cases of empty stack
	if element == nil {
		return nil, nil
	}

	s.head = element.nextElement
	s.size--

	s.traverse("--- after pop")
	return element.element, nil
}

// Peek returns the most recently added value from the stack, but leaves it on the stack
func (s *Stack) Peek() interface{} {
	return s.head.element
}

// Size returns the number of elements in the stack
func (s *Stack) Size() int {
	return s.size
}

// traverse A utility method to traverse the stack and print out the values.
func (s *Stack) traverse(message string) {
	if printTraversal {
		fmt.Println(message)

		var e = s.head

		for e != nil {
			fmt.Printf("element = %v\n", e.element)
			e = e.nextElement
		}
	}
}
