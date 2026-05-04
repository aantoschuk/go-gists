package doublyLinkedList

type Node[T any] struct {
	Value T
	Prev  *Node[T]
	Next  *Node[T]
}
