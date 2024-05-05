package main

import (
	"fmt"
	"godsa/list"
)

func main() {
	list := list.NewSinglyLinkedList[int]()
	fmt.Println(list)
}
