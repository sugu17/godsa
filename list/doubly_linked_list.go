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
	head   *doublyLinkedNode[T]
	tail   *doublyLinkedNode[T]
	length int
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
		dll.length += 1
		return newNode
	}
	newNode.next = head
	head.prev = newNode
	dll.head = newNode
	dll.length += 1
	return newNode
}

func (dll *DoublyLinkedList[T]) PushBack(value T) *doublyLinkedNode[T] {
	tail := dll.Back()
	newNode := &doublyLinkedNode[T]{
		next:  nil,
		prev:  nil,
		Value: value,
	}
	if tail == nil {
		dll.head = newNode
		dll.tail = newNode
		dll.length += 1
		return newNode
	}
	tail.next = newNode
	newNode.prev = tail
	dll.tail = newNode
	dll.length += 1
	return newNode
}

func (dll *DoublyLinkedList[T]) IndexOf(test func(val T) bool) (int, bool) {
	index := 0
	for node := dll.Front(); node != nil; node = node.Next() {
		if test(node.Value) == true {
			return index, true
		}
		index++
	}
	return -1, false
}

func (dll *DoublyLinkedList[T]) Remove(index int) (T, bool) {
	var defaultValue T
	if !(index >= 0 && index < dll.length) {
		return defaultValue, false
	}
	if index == 0 {
		nodeNextToHead := dll.head.next
		nodeNextToHead.prev = nil
		nodeVal := dll.head.Value
		dll.head = nodeNextToHead
		dll.length -= 1
		return nodeVal, true
	}
	if index == dll.length-1 {
		nodeBeforeTail := dll.tail.prev
		nodeBeforeTail.next = nil
		nodeVal := dll.tail.Value
		dll.tail = nodeBeforeTail
		dll.length -= 1
		return nodeVal, true
	}
	cur := 1
	for node := dll.Front().Next(); node.Next() != nil; {
		if index == cur {
			nodeBeforeTarget := node.Prev()
			nodeAfterTarget := node.Next()
			nodeVal := node.Value
			nodeBeforeTarget.next = nodeAfterTarget
			nodeAfterTarget.prev = nodeBeforeTarget
			node = nil
			dll.length -= 1
			return nodeVal, true
		}
		cur++
	}
	return defaultValue, false
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
		head:   nil,
		tail:   nil,
		length: 0,
	}
}
