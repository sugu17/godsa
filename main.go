package main

import (
	"fmt"
	"godsa/list"
)

func main() {
	list := list.NewDoublyLinkedList[int]()
	list.PushFront(10)
	list.PushFront(89)
	list.PushFront(95)
	fmt.Println(list)
}
