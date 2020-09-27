package main

/**
* Syntax Go. Homework 2
* Zaur Malakhov, dated Apr 24, 2019
*/

import (
	"fmt"
)

func main()  {
	//01
	fmt.Println("01. Написать функцию, которая определяет, четное ли число")
	n := 1
	fmt.Println(n, " - ", isEvenNumber(n))
	n = 4
	fmt.Println(n, " - ", isEvenNumber(n))

	//02
	fmt.Println()
	fmt.Println("02. Написать функцию, которая определяет, делится ли число без остатка на 3")
	n = 9
	fmt.Println(n, " - ", isDividedBy3(n))
	n = 19
	fmt.Println(n, " - ", isDividedBy3(n))

	//03
	fmt.Println()
	fmt.Println("03. Написать функцию, которая последовательно выводит на экран 100 первых чисел Фибоначчи,")
	fmt.Println("начиная от 0.")
	fmt.Println(selectFiboNumbers(0))
	fmt.Println(selectFiboNumbers(1))
	fmt.Println(selectFiboNumbers(2))
	fmt.Println(selectFiboNumbers(3))
	fmt.Println(selectFiboNumbers(4))
	fmt.Println(selectFiboNumbers(13))
	fmt.Println(selectFiboNumbers(100))


	// 04
	fmt.Println()
	fmt.Println("04. Заполнить массив из 100 элементов различными простыми числами")
	fmt.Println(selectPrimeNumbers(100))
	//fmt.Println(selectPrimeNumbers(1000))
}

func isEvenNumber(n int) bool {
	if n%2 == 0 {
		return true
	}
	return false
}

func isDividedBy3(n int) bool {
	if n%3 == 0 {
		return true
	}
	return false
}

func selectFiboNumbers(n int) ([]uint) {
	var fiboNumbers [] uint
	if n <= 0 {
		return fiboNumbers
	}
	for i := 0; i < n; i++ {
		if len(fiboNumbers) < 2 {
			fiboNumbers = append(fiboNumbers, uint(i))
		} else {
			fiboNumbers = append(fiboNumbers, (fiboNumbers[i-1] + fiboNumbers[i-2]))
		}
	} 
	return fiboNumbers
}

func selectPrimeNumbers(num int) (prime [] int) {
	var m [1000000] int

	n := 0
	//v := 0
	for i := 0; i < len(m); i++ {
		n = i + 2
		// v = m[i]

		for j := n*n-2; j < len(m); j +=n {
			m[j] = 1
		}

		if m[i] == 0 {
			prime = append(prime,n)
		}

		if len(prime) == num {
			break
		}
	}

	
	return
}









