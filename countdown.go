package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Countdown started...")
	tick := make(<-chan time.Time)
	tick = time.Tick(1 * time.Second)
	for countdown := 10; countdown > 0; countdown-- {
		moment := <-tick
		fmt.Println(moment.Format("15:04:05"), countdown)

	}
	fmt.Println("The rocket starts!")

}
