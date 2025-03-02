package main

import "fmt"

/*
How do you test this? It is good to separate your "domain" code from the outside world (side-effects).
The fmt.Println is a side effect (printing to stdout), and the string we send in is our domain.
*/

func helloK() string {
	var1 := "kashikuro"
	return var1
}

func main() {
	fmt.Println(helloK())
}
