package list

import (
	"fmt"
)

type singlyLinkedNode[T any] struct {
	next *singlyLinkedNode[T]
	data T
}

type SinglyLinkedList[T any] struct {
	head *singlyLinkedNode[T]
}

func (sll *SinglyLinkedList[T]) Front() *singlyLinkedNode[T] {
	return sll.head
}

func (sll *SinglyLinkedList[T]) Back() *singlyLinkedNode[T] {
	node := sll.Front()
	if node == nil {
		return nil
	}
	for {
		if node.next == nil {
			break
		}
		node = node.next
	}
	return node
}

func (sll *SinglyLinkedList[T]) PushBack(data T) *singlyLinkedNode[T] {
	lastNode := sll.Back()
	newNode := &singlyLinkedNode[T]{
		next: nil,
		data: data,
	}
	if lastNode == nil {
		sll.head = newNode
		return newNode
	}
	lastNode.next = newNode
	return newNode
}

func (sll SinglyLinkedList[T]) String() string {
	node := sll.Front()
	if node == nil {
		return "[]"
	}
	if node.next == nil {
		return fmt.Sprintf("[%v]", node.data)
	}
	output := "["
	for {
		if node.next == nil {
			output += fmt.Sprintf("%v", node.data)
			break
		}
		output += fmt.Sprintf("%v, ", node.data)
		node = node.next
	}
	output += "]"
	return output
}

func NewSinglyLinkedList[T any]() SinglyLinkedList[T] {
	return SinglyLinkedList[T]{
		head: nil,
	}
}
