package main

/*
* - если программа пытается открыть файл и возникает ошибка,
* то лучше сообщить пользователю об ошибке, а не просто выходить
* из программы.
* Так будет корректнее.
* log.Fatal(err)
 */

import (
	"fmt"
	"log"
	"os"
)

func main() {
	file, err := os.Open("fileread.go1")
	if err != nil {
		//return
		log.Fatal(err)
	}
	defer file.Close()

	// getting size of file
	stat, err := file.Stat()
	if err != nil {
		//return
		log.Fatal(err)
	}

	// reading file
	bs := make([]byte, stat.Size())
	_, err = file.Read(bs)
	if err != nil {
		//return
		log.Fatal(err)
	}

	fmt.Println(string(bs))
}
