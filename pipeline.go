package main

/**
* Syntax Go. Homework 7
* Zaur Malakhov, dated May 21, 2019
 */

import (
	"fmt"
	"time"
)

func main() {
	naturals := make(chan int)
	squares := make(chan int)
	// генерация
	go func() {
		for x := 0; x < 10; x++ {
			naturals <- x

		}
		close(naturals)
	}()
	// возведение в квадрат
	go func() {
		for {
			x, ok := <-naturals
			if !ok {
				break // канал закрыт и пуст
			}
			squares <- x * x
		}
		close(squares)
	}()
	// печать
	for {

		for x := range squares {
			//fmt.Println(<-squares)
			fmt.Println(x)
			time.Sleep(1 * time.Second)
		}
		break
	}
}
