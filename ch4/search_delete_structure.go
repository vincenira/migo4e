// Implement the delete() and search() functionality using generics for the linked list found in structures.go

package main

import (
	"fmt"
)

type node[T comparable] struct {
	Data T
	next *node[T]
}

type list[T comparable] struct {
	start *node[T]
}

func (l *list[T]) add(data T) {
	n := node[T]{
		Data: data,
		next: nil,
	}

	if l.start == nil {
		l.start = &n
		return
	}

	if l.start.next == nil {
		l.start.next = &n
		return
	}

	temp := l.start
	l.start = l.start.next
	l.add(data)
	l.start = temp
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
		fmt.Println("*", node.Data)
	}
}

func main() {
	var myList list[int]
	fmt.Println(myList)
	myList.add(12)
	myList.add(9)
	myList.add(3)
	myList.add(9)
	myList.PrintMe()
	value := myList.search(12)
	valueF := myList.search(8)
	myList.delete(9)
	myList.PrintMe()
	fmt.Println(value)
	fmt.Println(valueF)
}
