package list

import (
	"fmt"
	"log"
)

func init() {
	log.SetFlags(log.Lshortfile)
}

type doublyLinkedNode[T any] struct {
	next  *doublyLinkedNode[T]
	Value T
	prev  *doublyLinkedNode[T]
}

func (dln *doublyLinkedNode[T]) Next() *doublyLinkedNode[T] {
	return dln.next
}

func (dln *doublyLinkedNode[T]) Prev() *doublyLinkedNode[T] {
	return dln.prev
}

type DoublyLinkedList[T any] struct {
	head *doublyLinkedNode[T]
	tail *doublyLinkedNode[T]
}

func (dll *DoublyLinkedList[T]) Front() *doublyLinkedNode[T] {
	if dll.head == nil {
		return nil
	}
	return dll.head
}

func (dll *DoublyLinkedList[T]) Back() *doublyLinkedNode[T] {
	if dll.tail == nil {
		return nil
	}
	return dll.tail
}

func (dll *DoublyLinkedList[T]) PushFront(value T) *doublyLinkedNode[T] {
	head := dll.Front()
	newNode := &doublyLinkedNode[T]{
		next:  nil,
		Value: value,
		prev:  nil,
	}
	if head == nil {
		dll.head = newNode
		dll.tail = newNode
		return newNode
	}
	newNode.next = head
	head.prev = newNode
	dll.head = newNode
	return newNode
}

func (dll DoublyLinkedList[T]) String() string {
	output := "["
	for node := dll.Front(); node != nil; node = node.Next() {
		verb := "%v, "
		if node.Next() == nil {
			verb = "%v"
		}
		output += fmt.Sprintf(verb, node.Value)
	}
	output += "]"
	return output
}

func NewDoublyLinkedList[T any]() DoublyLinkedList[T] {
	return DoublyLinkedList[T]{
		head: nil,
		tail: nil,
	}
}
