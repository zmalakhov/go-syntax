package main

import (
	"fmt"
	"time"
	"io/ioutil"
)

func main() {
	t := time.Now()
	fmt.Println(t)
	fmt.Println(t.Format("02-01-2006"))
	fmt.Println(t.Format("02-01-2006 15:04:05"))
	time.Sleep(2000 * time.Millisecond)
	end := time.Now()
	fmt.Println(end.Sub(t))
	fmt.Println(end.Format("02-01-2006 15:04:05"))
	fmt.Println(t.Format(time.RFC822))

	arr1 := []byte("test")
	arr2 := []byte("тест")
	fmt.Println(arr1, arr2)

	fmt.Println("============")

	bs, err := ioutil.ReadFile("time.go")
	if err != nil {
		return
	}
	fmt.Println(string(bs))

}