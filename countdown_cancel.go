package main

import (
	"bufio"
	"fmt"
	"os"
	"time"
)

func main() {

	cancel := make(chan int)

	go func() {

		consolereader := bufio.NewReader(os.Stdin)
		fmt.Println("Enter 'exit' to exit: ")

		input, err := consolereader.ReadString('\n') // this will prompt the user for input

		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		if input == "e" {
			cancel <- 1
		}

		//os.Stdin.Read(make([]byte, 1))

		//cancel <- 1
	}()

	fmt.Println("Countdown started... exit to cansel")
	tick := make(<-chan time.Time)
	tick = time.Tick(1 * time.Second)

	for countdown := 1000; countdown > 0; countdown-- {
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
