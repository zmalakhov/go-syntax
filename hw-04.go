package main

/**
* Syntax Go. Homework 4
* Zaur Malakhov, dated May 03, 2019
*/

import (
	"fmt"
	"math"
	"sort"
)

type Figure interface {
	area() float64
}

type Circle struct {
	cx, cy, cr float64
}
func (circle Circle) area() float64 {
	return math.Pi * circle.cr * circle.cr
}

type Rectangle struct {
	x1, y1, x2, y2 float64
}
func (rectangle Rectangle) area() float64 {
	l := distance(rectangle.x1, rectangle.y1, rectangle.x1, rectangle.y2)
	w := distance(rectangle.x1, rectangle.y1, rectangle.x2, rectangle.y1)
	return l * w
}

func distance(x1, y1, x2, y2 float64) float64 {
	a := x2 - x1
	b := y2 - y1
	return math.Sqrt(a*a + b*b)
}

func allArea(figures ...Figure) float64 {
	var area float64
	for _, figure := range figures {
		if figure == nil {
			continue
		}
		area += figure.area()
	}
	return area
}

//=================================

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
	//01
	rectangle := Rectangle{0,0,10,10}
	fmt.Println(rectangle.area())

	circle := Circle{0,1,5}
	fmt.Println(circle.area())

	fmt.Println(allArea(rectangle, circle))

	//02
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







