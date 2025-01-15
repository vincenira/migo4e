// Implement a doubly-linked list using generics starting with the code found in structures.go
package main

import (
	"fmt"
)

type node[T comparable] struct {
	// store the value
	Data T
	// store the address of the next linked list
	next *node[T]
	// store the addres of the previous
	prev *node[T]
}

type list[T comparable] struct {
	start *node[T]
	end   *node[T]
}

func (l *list[T]) addAtEnd(data T) {
	n := node[T]{
		Data: data,
		next: nil,
		prev: nil,
	}

	if l.start == nil {
		l.start = &n
		l.end = &n
		return
	}

	if l.start.next == nil {
		temp := l.start
		l.start.next = &n
		l.start.next.prev = temp
		l.end = l.start.next
		return
	}

	temp := l.start
	l.start = l.start.next
	l.addAtEnd(data)
	l.start = temp
}

func (l *list[T]) addAtBegin(data T) {
	n := node[T]{
		Data: data,
		next: nil,
		prev: nil,
	}

	if l.start == nil {
		l.start = &n
		l.end = &n
		return
	}

	if l.end.prev == nil {
		temp := l.end
		l.end.prev = &n
		l.end.prev.next = temp
		l.start = l.end.prev

		return
	}
	temp := l.end
	l.end = l.end.prev
	l.addAtBegin(data)
	l.end = temp
}

func (l *list[T]) delete(data T) {
	current := l.start
	previous := l.start
	found := false
	for !found && current != nil {
		if current.Data == data {
			found = true
		} else {
			previous = current
			current = current.next
		}
	}
	if found {
		if l.start == nil {
			l.start = l.start.next
		} else {
			previous.next = current.next
		}
	}
}

func (l *list[T]) search(data T) bool {
	for node := l.start; node != nil; node = node.next {
		if node.Data == data {
			return true
		}
	}
	return false
}

func (l *list[T]) PrintMe() {
	for node := l.start; node != nil; node = node.next {
		fmt.Printf("* %d next: %p prev: %p\n", node.Data, node.next, node.prev)
	}
}

func (l *list[T]) BPrintMe() {
	for node := l.end; node != nil; node = node.prev {
		fmt.Printf("* %d next: %p prev: %p\n", node.Data, node.next, node.prev)
	}
}

func main() {
	var myList list[int]
	fmt.Println(myList)
	myList.addAtEnd(12)
	myList.addAtEnd(9)
	myList.addAtEnd(3)
	myList.addAtEnd(9)
	myList.addAtEnd(1)
	myList.addAtEnd(17)
	myList.addAtEnd(20)
	myList.addAtEnd(21)
	myList.addAtBegin(40)
	myList.addAtBegin(14)
	myList.addAtBegin(23)
	myList.PrintMe()
	fmt.Println("==========")
	myList.BPrintMe()
}
