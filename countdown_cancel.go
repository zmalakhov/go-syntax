package main

import (
	"fmt"
	"time"
)

func main() {
	cancel := make(chan int)

	go func() {
		//os.Stdin.Read(make([]byte, 1))

		var comm string

		fmt.Scan(&comm)
		if comm == "e" {
			cancel <- 1
		}

		//cancel <- 1
	}()

	fmt.Println("Countdown started... exit to cansel")
	tick := make(<-chan time.Time)
	tick = time.Tick(1 * time.Second)

	for countdown := 10; countdown > 0; countdown-- {
		select {
		case moment := <-tick:
			fmt.Println(moment.Format("15:04:05"), countdown)
		case <-cancel:
			fmt.Println("Launch canseled!")
			return
		}

	}
	fmt.Println("The rocket starts!")

}
