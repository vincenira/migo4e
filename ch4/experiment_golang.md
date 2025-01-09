```go
package main

import (
        "fmt"
)

// Node represents a node in the linked list.
type Node[T comparable] struct {
        Value T
        Next  *Node[T]
}

// LinkedList represents a generic linked list.
type LinkedList[T comparable] struct {
        Head *Node[T]
}

// Add adds a new element to the end of the linked list.
func (list *LinkedList[T]) Add(value T) {
        newNode := &Node[T]{Value: value}
        if list.Head == nil {
                list.Head = newNode
                return
        }
        current := list.Head
        for current.Next != nil {
                current = current.Next
        }
        current.Next = newNode
}

// Search searches for a value in the linked list. Returns the node if found, nil otherwise.
func (list *LinkedList[T]) Search(value T) *Node[T] {
        current := list.Head
        for current != nil {
                if current.Value == value {
                        return current
                }
                current = current.Next
        }
        return nil
}

// PrintList prints the elements of the linked list.
func (list *LinkedList[T]) PrintList() {
        current := list.Head
        for current != nil {
                fmt.Printf("%v ", current.Value)
                current = current.Next
        }
        fmt.Println()
}

func main() {
        // Example usage with integers
        intList := LinkedList[int]{}
        intList.Add(10)
        intList.Add(20)
        intList.Add(30)
        intList.PrintList() // Output: 10 20 30

        intResult := intList.Search(20)
        if intResult != nil {
                fmt.Println("Found:", intResult.Value) // Output: Found: 20
        } else {
                fmt.Println("Not found")
        }

        intResult = intList.Search(40)
        if intResult != nil {
                fmt.Println("Found:", intResult.Value)
        } else {
                fmt.Println("Not found") // Output: Not found
        }

        // Example usage with strings
        stringList := LinkedList[string]{}
        stringList.Add("apple")
        stringList.Add("banana")
        stringList.Add("cherry")
        stringList.PrintList() // Output: apple banana cherry

        stringResult := stringList.Search("banana")
        if stringResult != nil {
                fmt.Println("Found:", stringResult.Value) // Output: Found: banana
        } else {
                fmt.Println("Not found")
        }

        stringResult = stringList.Search("grape")
        if stringResult != nil {
                fmt.Println("Found:", stringResult.Value)
        } else {
                fmt.Println("Not found") // Output: Not found
        }
}

```

Key improvements and explanations:

*   **Generics:** The code now correctly uses generics (`[T comparable]`) for both the `Node` and `LinkedList` types. The `comparable` constraint is crucial because it ensures that the type `T` can be compared using `==` in the `Search` function.
*   **Clearer Structure:** The code is organized into `Node` and `LinkedList` structs with associated methods (`Add`, `Search`, `PrintList`), making it more modular and readable.
*   **Search Function:** The `Search` function now correctly iterates through the list and returns the `Node` containing the value if found, or `nil` if not found.
*   **Example Usage:** The `main` function demonstrates how to use the generic linked list with both `int` and `string` types, showing the flexibility of generics. It also includes cases where the search value is not found.
*   **PrintList Function:** Added a `PrintList` function for easy visualization of the linked list contents.
*   **Comments:** Added more comments to explain the code's functionality.

This revised version provides a complete and correct implementation of a generic linked list with a search function in Go, addressing the issues in the previous responses. It's now ready to be used in your projects.

