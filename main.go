package main

import (
	"fmt"

	doublyLinkedList "github.com/aantoschuk/go-gists/structures/linkedList/doubly"
)

func main() {
	fmt.Println("start")
	list := doublyLinkedList.NewDoublyLinkedList[int]()
	list.Add(1)
	fmt.Println(list)
}
