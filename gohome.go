package main

import (
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
	s := []byte(comm)

	if s == os.Stdin.Read(make([]byte, 4)) {
		fmt.Println(s)
	}

}
