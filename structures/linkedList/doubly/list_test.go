package doublyLinkedList

import (
	"fmt"
	"testing"

	"github.com/google/go-cmp/cmp"
)

// empty list test
func TestDoublyLinkedList_EmptyInit(t *testing.T) {
	list := NewDoublyLinkedList[int]()

	if !list.IsEmpty() {
		t.Fatal("expected empty list")
	}

	if list.Size != 0 {
		t.Fatal("expected size 0")
	}

	if list.Head != nil || list.Tail != nil {
		t.Fatal("expected nil head and tail")
	}

	if diff := cmp.Diff([]int{}, list.ToSlice()); diff != "" {
		t.Error(diff)
	}
}

func TestDoublyLinkedList_Prepand(t *testing.T) {
	value := 1

	list := NewDoublyLinkedList[int]()
	list.Prepend(value)

	if list.Size != 1 {
		t.Fatalf("expected list size to be 1, but got %d", list.Size)
	}

	if list.Head.Value != value {
		t.Fatalf("expected head value to be %d, but got %d", value, list.Head.Value)
	}

	if list.Tail.Value != value {
		t.Fatalf("expected tail value to be %d, but got %d", value, list.Tail.Value)
	}
}

// insert at
func TestDoublyLinkedList_InsertAt(t *testing.T) {
	list := NewDoublyLinkedList[int]()
	list.Add(1)
	list.Add(2)
	list.Add(3)

	list.InsertAt(1, 5)
	list.InsertAt(0, 0)
	list.InsertAt(list.Size, 10)

	r := list.ToSlice()
	expected := []int{0, 1, 5, 2, 3, 10}
	if diff := cmp.Diff(r, expected); diff != "" {
		t.Fatal(diff)
	}

}

// add to list
func TestDoublyLinkedList_Add(t *testing.T) {
	list := NewDoublyLinkedList[int]()

	list.Add(1)
	list.Add(2)
	list.Add(3)

	if list.Size != 3 {
		t.Fatalf("expected size 3, got %d", list.Size)
	}

	if diff := cmp.Diff([]int{1, 2, 3}, list.ToSlice()); diff != "" {
		t.Error(diff)
	}

	if list.Head.Value != 1 {
		t.Fatal("wrong head value")
	}

	if list.Tail.Value != 3 {
		t.Fatal("wrong tail value")
	}
}

// pop front test
func TestDoublyLinkedList_PopFront(t *testing.T) {
	list := NewDoublyLinkedList[int]()

	list.Add(1)
	list.Add(2)
	list.Add(3)

	n := list.PopFront()
	fmt.Println(list.Size)
	if n == nil || n.Value != 1 {
		t.Fatal("wrong popped value")
	}

	if diff := cmp.Diff([]int{2, 3}, list.ToSlice()); diff != "" {
		t.Error(diff)
	}

	if list.Head.Value != 2 {
		t.Fatal("wrong head after pop")
	}
}

// pop back test
func TestDoublyLinkedList_PopBack(t *testing.T) {
	list := NewDoublyLinkedList[int]()

	list.Add(1)
	list.Add(2)
	list.Add(3)

	n := list.PopBack()
	if n == nil || n.Value != 3 {
		t.Fatal("wrong popped value")
	}

	if diff := cmp.Diff([]int{1, 2}, list.ToSlice()); diff != "" {
		t.Error(diff)
	}

	if list.Tail.Value != 2 {
		t.Fatal("wrong tail after pop")
	}
}

// pop back list with single item
func TestDoublyLinkedList_PopBack_Single(t *testing.T) {
	list := NewDoublyLinkedList[int]()

	list.Add(1)

	n := list.PopBack()
	if n == nil || n.Value != 1 {
		t.Fatal("wrong popped value")
	}

	if !list.IsEmpty() {
		t.Fatal("list should be empty")
	}

	if list.Head != nil || list.Tail != nil {
		t.Fatal("head/tail should be nil after pop")
	}
}

// testing peek functions
func TestDoublyLinkedList_Peek(t *testing.T) {
	list := NewDoublyLinkedList[int]()

	// peek on empty list

	if _, ok := list.PeekBack(); ok {
		t.Fatal("peekBack: list should be empty, but got non-zero value")
	}

	if _, ok := list.PeekFront(); ok {
		t.Fatal("peekFront: list should be empty, but got non-zero value")
	}

	// peek on list with values

	list.Add(1)
	list.Add(2)

	value, ok := list.PeekBack()
	if _, ok := list.PeekBack(); !ok || value != 2 {
		t.Fatalf("peekBack: list should return value. expected 2 and true, but got: %d, %v", value, ok)
	}

	value, ok = list.PeekFront()
	if !ok || value != 1 {
		t.Fatalf("peekFront: list should return value. expected 1 and true, but got: %d, %v", value, ok)
	}
}

func TestDoublyLinkedList_Delete(t *testing.T) {
	list := NewDoublyLinkedList[int]()
	list.Add(1)
	list.Add(2)
	list.Add(3)
	list.Add(4)

	// the result should be [1,3,4]
	list.Delete(1)

	slice := list.ToSlice()

	if diff := cmp.Diff(slice, []int{1, 3, 4}); diff != "" {
		t.Fatal(diff)
	}

	// delete current head
	list.Delete(0)

	slice = list.ToSlice()
	if diff := cmp.Diff(slice, []int{3, 4}); diff != "" {
		t.Fatal(diff)
	}

	// delete current tail
	list.Delete(list.Size - 1)
	slice = list.ToSlice()

	if diff := cmp.Diff(slice, []int{3}); diff != "" {
		t.Fatal(diff)
	}

}

func TestDoublyLinkedList_Get(t *testing.T) {
	list := NewDoublyLinkedList[int]()

	// get on empty list

	if _, ok := list.Get(2); ok {
		t.Fatal("should expected to return false for empty list, but got true.")
	}

	// negative position
	if _, ok := list.Get(-1); ok {
		t.Fatal("should expected to return false for empty list, but got true.")
	}

	// get with on list with values

	list.Add(1)
	list.Add(2)
	list.Add(3)

	if value, ok := list.Get(0); !ok || value != 1 {
		t.Fatal("Get with proper position should return a value and true")
	}

	if value, ok := list.Get(1); !ok || value != 2 {
		t.Fatal("Get with proper position should return a value that expected but got wrong one and true")
	}

	if value, ok := list.Get(2); !ok || value != 3 {
		t.Fatal("Get with proper position should return a value and true")
	}

}

func TestDoublyLinkedList_Clear(t *testing.T) {
	l := NewDoublyLinkedList[int]()
	l.Add(1)
	l.Add(2)
	l.Add(3)

	slice := l.ToSlice()
	if diff := cmp.Diff(slice, []int{1, 2, 3}); diff != "" {
		t.Fatalf("list is not properly created: %v", diff)
	}

	l.Clear()
	slice = l.ToSlice()

	if diff := cmp.Diff(slice, []int{}); diff != "" {
		t.Fatalf("should be empty slice but list contains some values: %v", diff)
	}
}

func TestDoublyLinkedList_Update(t *testing.T) {
	l := NewDoublyLinkedList[int]()
	l.Add(2)
	l.Add(3)
	l.Add(4)

	l.Update(0, 1)
	l.Update(1, 2)
	l.Update(2, 3)
	slice := l.ToSlice()

	if diff := cmp.Diff(slice, []int{1, 2, 3}); diff != "" {
		t.Fatalf("values in the list is not updated: %v", diff)
	}

}

func TestDoublyLinkedList_Reverse(t *testing.T) {
	l := NewDoublyLinkedList[int]()
	l.Add(1)
	l.Add(2)
	l.Add(3)

	l.Reverse()
	slice := l.ToSlice()

	if diff := cmp.Diff(slice, []int{3, 2, 1}); diff != "" {
		t.Fatalf("slice should be in desc order after list reverse: %v", diff)
	}

}

func TestDoublyLinkedList_Extend(t *testing.T) {
	l := NewDoublyLinkedList[int]()
	l.Add(1)
	fill := []int{2, 3, 4}

	l.Extend(fill)
	slice := l.ToSlice()

	if diff := cmp.Diff(slice, []int{1, 2, 3, 4}); diff != "" {
		t.Fatalf("list should be filled with elements of the provided slice, but not: %v", diff)
	}
}

func TestDoublyLinkedList_Fill(t *testing.T) {
	l := NewDoublyLinkedList[int]()
	fill := []int{1, 2, 3}

	l.Fill(fill)
	slice := l.ToSlice()

	if diff := cmp.Diff(slice, fill); diff != "" {
		t.Fatalf("list should be filled with elements of the provided slice, but not: %v", diff)
	}
}

func TestDoublyLinkedList_Contains(t *testing.T) {
	intEqual := func(a, b int) bool {
		return a == b
	}

	l := NewDoublyLinkedList[int]()
	l.Equal = intEqual
	l.Add(1)
	l.Add(2)
	l.Add(3)
	contains := l.Contains(2)

	if !contains {
		t.Fatal("Contains should return true for searched value, bot got false")
	}

}
