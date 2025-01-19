/*
Use the empty interface and a function that allows you to differentiate between two dif-
ferent structures that you have created.
*/
package main

import (
	"fmt"
	"reflect"
)

type Animal struct {
	Name string
}

type Fish struct {
	Name  string
	Color string
}

func Differentiate(a, b interface{}) {
	Atype := reflect.TypeOf(a)
	Btype := reflect.TypeOf(b)
	if Atype == Btype {
		fmt.Println("the two structures are the same")
	} else {
		fmt.Println("The two structures are different")
	}

	fmt.Printf("the first argument type is %s and second argument type is %s\n", Atype.Name(), Btype.Name())
}

func main() {
	a := Animal{"girafe"}
	b := Fish{"whale", "grey"}
	fmt.Println("let's test if two structures are different or not")
	fmt.Println("exploring reflect in golang")
	fmt.Printf("\n\n")
	Differentiate(a, b)
	c := Animal{"sheep"}
	d := Animal{"dog"}
	fmt.Printf("\n\n")
	Differentiate(c, d)
	f := Fish{"whale", "grey"}
	e := Fish{"whale", "grey"}
	fmt.Printf("\n\n")
	Differentiate(f, e)
}
