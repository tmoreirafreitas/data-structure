package generics

import (
	"errors"
	"reflect"
)

type NodeList struct {
	Data any
	prev *NodeList
	next *NodeList
}

// Next returns the next list element or nil.
func (n *NodeList) Next() *NodeList {
	return n.next
}

// Prev returns the previus list element or nil
func (n *NodeList) Prev() *NodeList {
	return n.prev
}

// List represents a doubly linked list. The zero value for List is an empty list ready to use.
type List[T any] struct {
	length uint
	front  *NodeList
	rear   *NodeList
}

// Front returns the first element of list l or nil if the list is empty.
func (l *List[T]) Front() *NodeList {
	return l.front
}

// Rear returns the last element of list l or nil if the list is empty.
func (l *List[T]) Rear() *NodeList {
	return l.rear
}

// Count returns the number of elements of list. The complexity is O(1).
func (l *List[T]) Count() uint {
	return l.length
}

// InsertStart inserts a new node with value element at the start of list.
func (l *List[T]) InsertStart(element *T) error {
	if element == nil {
		return errors.New("element cannot be nil")
	}

	var node *NodeList = new(NodeList)
	node.Data = element
	node.prev = nil

	if l.front == nil {
		l.rear = node
	} else {
		l.front.prev = node
		node.next = l.front
	}

	l.front = node
	l.length += 1

	return nil
}

// IsEmpty check if list is empty
func (l *List[T]) IsEmpty() bool {
	return l.front == nil
}

// Insert inserts by default a new node with value element at the back of list
func (l *List[T]) Insert(element *T) error {
	if element == nil {
		return errors.New("element cannot be nil")
	}

	node := new(NodeList)
	node.Data = element
	node.next = nil

	if l.rear == nil {
		l.front = node
	} else {
		l.rear.next = node
		node.prev = l.rear.next
	}
	l.rear = node
	l.length += 1

	return nil
}

// InsertAt inserts the element at a specific position in the list. Return error if index out of range
func (l *List[T]) InsertAt(element *T, position uint) error {
	var err error = nil
	if element == nil {
		err = errors.New("element cannot be nil")
	}

	if position > l.length-1 {
		err = errors.New("Index out of range.")
	}

	currentNode := l.front
	var index uint = 0

	node := new(NodeList)
	node.Data = element

	if position == 0 {
		err = l.InsertStart(element)
	} else {
		for currentNode != nil {
			if index == position {
				prev := currentNode.next
				node.prev = prev
				prev.next = node

				node.next = currentNode
				currentNode.prev = node
				l.length += 1

				break
			}

			index += 1
			currentNode = currentNode.next
		}

		if index != position {
			err = l.Insert(element)
		}
	}

	return err
}

// Contains checks if there is an element with the value of the field informed
func (l List[T]) Contains(field string, value any) bool {
	currentNode := l.front
	finding := false
	for currentNode != nil {
		data := currentNode.Data.(*T)
		e := reflect.ValueOf(data).Elem()

		for i := 0; i < e.NumField(); i++ {
			varName := e.Type().Field(i).Name
			varType := e.Type().Field(i).Type
			varValue := e.Field(i).Interface()

			if varName == field && varType == reflect.ValueOf(value).Type() && varValue == value {
				finding = true
			}
		}

		currentNode = currentNode.next
	}

	return finding
}

// GetByIndex returns element based on the given index
func (l List[T]) GetByIndex(index uint) (*T, error) {

	if index > l.length-1 {
		return nil, errors.New("Index out of range.")
	}

	if l.front == nil {
		return nil, nil
	}

	var data *T = nil
	currentNode := l.front

	var position uint = 0
	for currentNode != nil {
		if position == index {
			data = currentNode.Data.(*T)
			break
		}
		currentNode = currentNode.next
		position += 1
	}

	return data, nil
}

// GetIndex returns the index of the searched element.
// If index is -1 it is because the element is not in the list
func (l List[T]) GetIndex(element T) int {
	currentNode := l.front
	index := 0
	findingNode := -1
	for currentNode != nil {
		data := currentNode.Data.(*T)
		if reflect.DeepEqual(*data, element) {
			findingNode = index
			break
		}
		index += 1
		currentNode = currentNode.next
	}

	return findingNode
}

// Remove removes element from list if element is an item of list. The element must not be nil.
func (l *List[T]) Remove(element *T) {
	var prev *NodeList = nil
	var next *NodeList = nil
	nodeToRemove := l.find(*element)

	if nodeToRemove != nil {
		prev = nodeToRemove.prev
		next = nodeToRemove.next

		if prev == nil && next != nil {
			l.front = next
		} else if prev != nil && next == nil {
			l.rear = prev
		}

		if prev != nil {
			prev.next = next
		}

		if next != nil {
			next.prev = prev
		}

		nodeToRemove = nil
		l.length -= 1
	}
}

func (l *List[T]) find(element T) *NodeList {
	currentNode := l.front
	var findingNode *NodeList = nil
	for currentNode != nil {
		data := currentNode.Data.(*T)
		if reflect.DeepEqual(*data, element) {
			findingNode = currentNode
			break
		}
		currentNode = currentNode.next
	}

	return findingNode
}
