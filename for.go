package main

import (
	"fmt"
)

func main_() {
	//бесконечный цикл
	// for {
	// 	fmt.Print("F")
	// }


	// цикл с условием
	// i := 0
	// for i < 10 {
	// 	fmt.Print("F")
	// 	i++
	// }

	// fmt.Println("")
	// // классический цикл
	
	// for i := 0; i < 10; i++ {
	// 	if i==5 {
	// 		break
	// 	}
	// 	fmt.Print(i)
	// }
	// fmt.Println("")
	// for i := 0; i < 10; i++ {
	// 	if i==5 {
	// 		continue
	// 	}
	// 	fmt.Print(i)
	// }

	// // обход по сущности (коллекции)
	// str := "Привет go"
	// for i, ch := range str {
	// 	fmt.Println("Code", ch, "has number", i)
	// }

	// // пустой идентификатор
	// str = "Привет go"
	// for _, ch := range str {
	// 	fmt.Println("Code", ch, "has number")
	// }


	/*str := "Привет go"
	for i, ch := range str {
		fmt.Println("Code", ch, "has number", i)
		if ch >255 {
			fmt.Println("Unicode")
		} else {
			fmt.Println("Ascii")
		}
	}*/


	a :=1
	switch a {
		case 1: 
			fmt.Println("One")
		case 2: 
			fmt.Println("Two")
		default: 
			fmt.Println("What?")
	}



}