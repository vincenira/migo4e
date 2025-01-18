/*
Create a slice of structures using a structure that you created and sort the elements of the
slice using a field from the structure.
*/

package main

import (
	"fmt"
	"sort"
)

type Contact struct {
	Lname  string
	Fname  string
	Number string
}

type Contacts []*Contact

func (c Contacts) Len() int      { return len(c) }
func (c Contacts) Swap(i, j int) { c[i], c[j] = c[j], c[i] }

type ByLname struct{ Contacts }

func (c ByLname) Less(i, j int) bool { return c.Contacts[i].Lname < c.Contacts[j].Lname }

type ByFname struct{ Contacts }

func (c ByFname) Less(i, j int) bool { return c.Contacts[i].Fname < c.Contacts[j].Fname }

type ByNumber struct{ Contacts }

func (n ByNumber) Less(i, j int) bool { return n.Contacts[i].Number < n.Contacts[j].Number }
func printContacts(cs []*Contact) {
	for _, c := range cs {
		fmt.Printf("%-4s %-2s (%s)\n", c.Lname, c.Fname, c.Number)
	}
}

func main() {
	s := []*Contact{
		{"tete", "u", "012-130-402920"},
		{"titi", "i", "012-131-402920"},
		{"tutu", "a", "012-131-412920"},
		{"tata", "o", "112-131-412920"},
		{"toto", "e", "012-131-412921"},
	}
	fmt.Println("Sorted by LastName")
	sort.Sort(ByLname{s})
	printContacts(s)
	fmt.Println("")
	fmt.Println("")
	fmt.Println("Sorted by FirstName")
	sort.Sort(ByFname{s})
	printContacts(s)
	fmt.Println("")
	fmt.Println("")
	fmt.Println("Sorted by Number")
	sort.Sort(ByNumber{s})
	printContacts(s)
}
