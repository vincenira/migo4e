// Implement the delete() and search() functionality using generics for the linked list found in structures.go

package main

import (
	"fmt"
)

type node[E comparable] struct {
	Data E
	next *node[E]
}

type list[E comparable] struct {
	start *node[E]
}

func (l *list[E]) add(data E) {
	n := node[E]{
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

func (l *list[E]) search(data E) bool {
	for l.start != nil {
		if l.start.Data == data {
			return true
		}
		l.start = l.start.next
	}
	return false
}

func (l *list[E]) PrintMe() {
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
	value := myList.search(12)
	myList.PrintMe()
	valueF := myList.search(9)
	fmt.Println(value)
	fmt.Println(valueF)
}
