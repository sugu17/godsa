package list

import (
	"fmt"
	"log"
)

func init() {
	log.SetFlags(log.Lshortfile)
}

type singlyLinkedNode[T any] struct {
	next  *singlyLinkedNode[T]
	Value T
}

func (sln *singlyLinkedNode[T]) Next() *singlyLinkedNode[T] {
	return sln.next
}

type SinglyLinkedList[T any] struct {
	head *singlyLinkedNode[T]
}

func (sll *SinglyLinkedList[T]) Front() *singlyLinkedNode[T] {
	return sll.head
}

func (sll *SinglyLinkedList[T]) Back() *singlyLinkedNode[T] {
	node := sll.Front()
	for ; node != nil && node.Next() != nil; node = node.Next() {
	}
	return node
}

func (sll *SinglyLinkedList[T]) PushFront(value T) *singlyLinkedNode[T] {
	head := sll.Front()
	if head == nil {
		sll.head = &singlyLinkedNode[T]{
			next:  nil,
			Value: value,
		}
		return sll.head
	}
	prev := head
	sll.head = &singlyLinkedNode[T]{
		next:  prev,
		Value: value,
	}
	return sll.head
}

func (sll *SinglyLinkedList[T]) PushBack(value T) *singlyLinkedNode[T] {
	lastNode := sll.Back()
	newNode := &singlyLinkedNode[T]{
		next:  nil,
		Value: value,
	}
	if lastNode == nil {
		sll.head = newNode
		return newNode
	}
	lastNode.next = newNode
	return newNode
}

func (sll *SinglyLinkedList[T]) Remove(index int) (T, bool) {
	if index == 0 {
		prevVal := sll.head.Value
		sll.head = sll.head.next
		return prevVal, true
	}
	var defValue T
	curr := 0
	for node := sll.Front(); node.next != nil; node = node.next {
		if index == curr+1 {
			prevVal := node.next.Value
			nodeAfterTarget := node.next.next
			node.next = nodeAfterTarget
			return prevVal, true
		}
		curr++
	}
	return defValue, false

}

func (sll *SinglyLinkedList[T]) IndexOf(target T, test func(val T, target T) bool) (int, bool) {
	index := 0
	for node := sll.Front(); node != nil; node = node.Next() {
		if test(node.Value, target) {
			return index, true
		}
		index++
	}
	return -1, false
}

func (sll SinglyLinkedList[T]) String() string {
	output := "["
	for node := sll.Front(); node != nil; node = node.Next() {
		verb := "%v, "
		if node.Next() == nil {
			verb = "%v"
		}
		output += fmt.Sprintf(verb, node.Value)
	}
	output += "]"
	return output
}

func NewSinglyLinkedList[T any]() SinglyLinkedList[T] {
	return SinglyLinkedList[T]{
		head: nil,
	}
}
