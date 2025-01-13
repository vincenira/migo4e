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

func (l *list[T]) Fadd(data T) {
	n := node[T]{
		Data: data,
		next: nil,
		prev: nil,
	}

	if l.start == nil {
		l.start = &n
		return
	}

	if l.start.next == nil {
		l.start.next = &n
		l.start.prev = l.start
		return
	}

	temp := l.start
	l.start = l.start.next

	l.Fadd(data)
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
	myList.Fadd(12)
	myList.Fadd(9)
	myList.Fadd(3)
	myList.Fadd(9)
	myList.Fadd(1)
	myList.Fadd(17)
	myList.PrintMe()
}
