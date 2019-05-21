package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	//go process()
	end()
}

func process() {
	for x := 1; ; x++ {
		fmt.Println(x)
	}
}

func end() {
	// var comm string

	// fmt.Scan(&comm)
	// if comm == "exit" {
	// 	fmt.Println(comm)
	// }

	comm := "exit"
	// s := []byte(comm)

	// if s == os.Stdin.Read(make([]byte, 4)) {
	// 	fmt.Println(s)
	// }
	consolereader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter 'exit' to exit: ")

	input, err := consolereader.ReadString('\n') // this will prompt the user for input

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	if input == comm {
		fmt.Println(input)
	}

}
