// Implement the delete() and search() functionality using generics for the linked list found in structures.go

package main

import (
	"fmt"
)

type node[T any] struct {
	Data T
	next *node[T]
}

type list[T any] struct {
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

func (l *list[T]) search(data T) {
	// need to define a comparable function and finish the delete function
	for l.start != nil {

		if data == l.start.Data {
			fmt.Println(l.start.Data)
		}
		l.start = l.start.next
	}
}

func (l *list[T]) PrintMe() {
	for l.start != nil {
		fmt.Println("*", l.start.Data)
		l.start = l.start.next
	}
}

func main() {
	var myList list[int]
	fmt.Println(myList)
	myList.add(12)
	myList.add(9)
	myList.add(3)
	myList.add(9)

	// Print all elements
	cur := myList.start
	for {
		fmt.Println("*", cur)
		if cur == nil {
			break
		}
		cur = cur.next
	}

	myList.PrintMe()
}