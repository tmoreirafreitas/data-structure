package generics

import "errors"

type nodeStack struct {
	data any
	next *nodeStack
}

type Stack[T any] struct {
	length uint
	top    *nodeStack
}

// IsEmpty check if stack is empty
func (s *Stack[T]) IsEmpty() bool {
	return s.top == nil
}

// Count returns the number of element of stack. The complexity is O(1).
func (s *Stack[T]) Count() uint {
	return s.length
}

// Push a new value onto the stack. Return error if element is nil.
func (s *Stack[T]) Push(element *T) error {
	if element == nil {
		return errors.New("Item cannot be nil")
	}

	newElement := new(nodeStack)
	newElement.data = element
	newElement.next = s.top
	s.top = newElement
	s.length += 1
	return nil
}

// Remove and return top element of stack. Return error if stack is empty.
func (s *Stack[T]) Pop() (*T, error) {
	if !s.IsEmpty() {
		tempPtr := s.top
		data := tempPtr.data.(*T)
		s.top = s.top.next
		tempPtr = nil
		return data, nil
	} else {
		return nil, errors.New("Stack is empty!")
	}
}
