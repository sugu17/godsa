package main

import (
	"fmt"
	"godsa/list"
)

func main() {
	list := list.NewSinglyLinkedList[int]()
	list.PushBack(3)
	list.PushBack(6)
	list.PushBack(10)
	list.PushFront(19)
	list.Remove(3)
}
