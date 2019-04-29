package main

import (
	"fmt"
)

func main(){
	// var r [10]int
	// for i := 0; i < len(r); i++ {
	// 	r[i] = i+1
	// }
	// fmt.Println(r)

	// arr := [5]int{1,2,3,4,5}
	// for i, v := range arr {
	// 	fmt.Printf("arr [%d] = %d\n", i, v)
	// }

	// //x := make([]int , 5)
	// x := arr[1:2]
	// for i, v := range x {
	// 	fmt.Printf("x [%d] = %d\n", i, v)
	// }

	// копирование срезов
	// s1 := []int{1,2,3,4,5}
	// s2 := make([]int, 3)

	// fmt.Println(s1, s2)
	// copy(s2, s1)
	// fmt.Println(s1, s2)

	// scores := make([]int, 0, 5)
	// c := cap(scores)
	// fmt.Println(c)

	// for i := 0; i <25; i++ {
	// 	scores = append(scores,i)
	// 	fmt.Println(scores)


	// 	if cap(scores) != c {
	// 		c = cap(scores)
	// 		fmt.Println(c)
	// 	}
	// }

	ab := make(map[string][]string)
	ab["Alex"] = []string{"1234567"}
	ab["Bob"] =  []string{"24578128"}
	ab["Bob"] = append(ab["Bob"], "92647814")

	fmt.Println(ab)

	for name, nums := range ab {
		fmt.Println("Абонент: ", name)
		for i, num := range nums {
			fmt.Printf("\t %v) %v \n", i+1, num)
		}
	}

	//структуры
	type position struct {
		X int
		Y int
	}

	type Circle struct {
		Point position
		Radius int
	}

	FirstCircle := Circle{Point: position{18,21}, Radius: 43}
	SecondCircle := Circle{
		Point: position{35,21}, 
		Radius: 41,
	}

	fmt.Println(FirstCircle, SecondCircle)

	SecondCircle.Point.Y = 78

	fmt.Println(FirstCircle, SecondCircle)

}