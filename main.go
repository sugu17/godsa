package main

import (
	"fmt"
	"godsa/list"
)

func main() {
	list := list.NewDoublyLinkedList[int]()
	list.PushBack(54)
	list.PushBack(89)
	list.PushBack(43)
	fmt.Println(list)
	fmt.Println("Value at the head", list.Front().Value())
	fmt.Println("Value at the end", list.Back().Value())
}
