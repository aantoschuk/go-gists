package doublyLinkedList

import (
	"fmt"
	"strings"
)

type DoublyLinkedList[T any] struct {
	Head  *Node[T]
	Tail  *Node[T]
	Size  int
	Equal func(T, T) bool
}

// Allocates empty DoublyLinkedList struct
func NewDoublyLinkedList[T any]() *DoublyLinkedList[T] {
	return &DoublyLinkedList[T]{
		Head: nil,
		Tail: nil,
		Size: 0,
	}
}

// IsEmpty functions returns true if list has nodes
func (l *DoublyLinkedList[T]) IsEmpty() bool {
	return l.Size == 0
}

// Contains function check if the specified value is in the list
func (l *DoublyLinkedList[T]) Contains(v T) bool {
	for n := l.Head; n != nil; n = n.Next {
		if l.Equal(n.Value, v) {
			return true
		}
	}
	return false
}

func (l *DoublyLinkedList[T]) unlink(curr *Node[T]) {
	prev := curr.Prev
	next := curr.Next

	if prev != nil {
		prev.Next = next
	} else {
		l.Head = next
	}

	if next != nil {
		next.Prev = prev
	} else {
		l.Tail = prev
	}

	l.Size--
}

// Prepend adds a new node to the start of the list
func (l *DoublyLinkedList[T]) Prepend(v T) {
	node := &Node[T]{Value: v}

	oldHead := l.Head

	node.Next = oldHead

	if oldHead != nil {
		oldHead.Prev = node
	} else {
		l.Tail = node
	}

	l.Head = node
	l.Size++
}

// InsertAt functions inserts a new node at specified position
func (l *DoublyLinkedList[T]) InsertAt(pos int, v T) bool {
	if pos > l.Size || pos < 0 {
		return false
	}

	if pos == 0 {
		l.Prepend(v)
		return true
	}
	if pos == l.Size {
		l.Add(v)
		return true
	}

	curr := l.Head
	for range pos {
		curr = curr.Next
	}

	node := &Node[T]{Value: v, Next: nil, Prev: nil}
	node.Prev = curr.Prev
	node.Next = curr
	curr.Prev = node
	node.Prev.Next = node

	return true
}

// Add adds item to the tail of the list
func (l *DoublyLinkedList[T]) Add(val T) {
	node := &Node[T]{Value: val, Next: nil, Prev: nil}

	if l.IsEmpty() {
		l.Head = node
		l.Tail = node
		l.Size++
		return
	}

	node.Prev = l.Tail
	l.Tail.Next = node
	l.Tail = node
	l.Size++
}

// PopFront removes first element from the List and updates head
func (l *DoublyLinkedList[T]) PopFront() *Node[T] {
	if l.Head == nil {
		return nil
	}
	n := l.Head
	l.unlink(n)
	return n
}

// PopBack removes last element from the list and updates it's tail
func (l *DoublyLinkedList[T]) PopBack() *Node[T] {
	if l.Tail == nil {
		return nil
	}
	n := l.Tail
	l.unlink(n)
	return n
}

// PeekFront return first value from the list and bool variable,
// which will be false if list is empty
func (l *DoublyLinkedList[T]) PeekFront() (T, bool) {
	var value T
	if l.IsEmpty() {
		return value, false
	}
	value = l.Head.Value

	return value, true
}

// PeekBack return last value from the list and bool variable,
// which will be false if list is empty
func (l *DoublyLinkedList[T]) PeekBack() (T, bool) {
	var value T
	if l.IsEmpty() {
		return value, false
	}
	value = l.Tail.Value

	return value, true
}

// Get function returns value of the list at the specific node.
// false if specified position is not proper or not exists in the list.
func (l *DoublyLinkedList[T]) Get(pos int) (T, bool) {
	var zero T

	if pos < 0 || pos >= l.Size {
		return zero, false
	}

	curr := l.Head
	for range pos {
		curr = curr.Next
	}

	return curr.Value, true
}

// Delete node at specified position
func (l *DoublyLinkedList[T]) Delete(pos int) bool {
	if pos < 0 || pos >= l.Size {
		return false
	}

	curr := l.Head
	for range pos {
		curr = curr.Next
	}

	l.unlink(curr)
	return true
}

// Update functions changes value at the specified position
func (l *DoublyLinkedList[T]) Update(pos int, v T) bool {
	if pos < 0 || pos >= l.Size {
		return false
	}

	curr := l.Head
	for range pos {
		curr = curr.Next
	}
	curr.Value = v
	return true
}

// Reserse reverses list
func (l *DoublyLinkedList[T]) Reverse() {
	if l.IsEmpty() {
		return
	}

	curr := l.Head
	for curr != nil {
		temp := curr.Next

		curr.Next = curr.Prev
		curr.Prev = temp

		curr = temp
	}

	l.Head, l.Tail = l.Tail, l.Head
}

// Clear functions clears whole list
func (l *DoublyLinkedList[T]) Clear() {
	curr := l.Head
	for curr != nil {
		next := curr.Next
		curr.Next = nil
		curr.Prev = nil
		curr = next
	}
	l.Head = nil
	l.Tail = nil
	l.Size = 0
}

// ToSlice converts list into slice containing values of each node.
func (l *DoublyLinkedList[T]) ToSlice() []T {
	out := make([]T, 0)
	for n := l.Head; n != nil; n = n.Next {
		out = append(out, n.Value)
	}
	return out
}

// Extend funtion extends list by adding elements of the slice
func (l *DoublyLinkedList[T]) Extend(s []T) {
	for _, v := range s {
		l.Add(v)
	}
}

// Fill feels list with elemets of the provided slice
func (l *DoublyLinkedList[T]) Fill(s []T) {
	l.Clear()
	for _, v := range s {
		l.Add(v)
	}
}

func (list *DoublyLinkedList[T]) String() string {
	if list.IsEmpty() {
		return "empty list"
	}
	sb := strings.Builder{}
	n := list.Head
	for n != nil {
		fmt.Fprint(&sb, n.Value)
		if n.Next != nil {
			fmt.Fprint(&sb, " -> ")
		}
		n = n.Next
	}
	return sb.String()
}
