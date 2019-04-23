package main

import (
	"fmt"
)

func main() {

	// xs := []float64{1,2,3}
	// l, t, a := average(xs...)
	// fmt.Println(l, t, a)

	// defer fmt.Println("Second")
	// fmt.Println("First")

	// add := func(x, y int) int {
	// 	return x + y
	// }
	// fmt.Println(add(1,1))

	// x := 0
	// increment := func() int {
	// 	x++
	// 	return x
	// }

	// fmt.Println(increment())
	// fmt.Println(increment())

	// fmt.Println(factorial(5))
	// fmt.Println(factorial(6))


	//===============
	defer func() {
		str := recover()
		fmt.Println(str)
	}()

	b := 0
	fmt.Println(4 / b)
	fmt.Println("ok")


}

func factorial(x int) int {
	if x == 0 {
		return 1
	}
	return x*factorial(x-1)
}


func average(xs ...float64) (lenght int, total float64, average float64) {
	
	total = 0.
	for _, v := range xs {
		total += v
	}
	lenght = len(xs)
	average = total / float64(lenght)
	return 
	
}







