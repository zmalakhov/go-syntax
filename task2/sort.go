package main

import (
	"fmt"
	"sort"
)

type Phones [] int

// Person for list
type PhoneBook struct {
	Name string
	phones Phones
}

// List for sorting
type List []PhoneBook

// Len is the number of elements in the collection
func (l List) Len() int {
	return len(l)
}

// Less reports whether the element with
// index i should sort before the element with index j.
func (l List) Less(i, j int) bool {
	return l[i].Name < l[j].Name
}

// Swap swaps the elements with indexes i and j
func (l List) Swap(i, j int) {
	l[i], l[j] = l[j], l[i]
}

func main() {
	phoneBook := []PhoneBook{
		{"Oliver", Phones{9514785694, 4785147963}},
		{"James", Phones{2547895123}},
		{"Alfie", Phones{6587458921}},
		{"Emily", Phones{5478598712}},
		{"Heather", Phones{5478123598}},
		{"Isabella", Phones{3125478131}},
	}
	fmt.Println(phoneBook)
	sort.Sort(List(phoneBook))
	fmt.Println(phoneBook)
}
