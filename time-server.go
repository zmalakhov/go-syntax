package main

import (
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"time"
)

func main() {

	cancel := make(chan int)

	go func() {

		// consolereader := bufio.NewReader(os.Stdin)
		// fmt.Println("Enter 'exit' to exit: ")

		// input, err := consolereader.ReadString('\n') // this will prompt the user for input

		// if err != nil {
		// 	fmt.Println(err)
		// 	os.Exit(1)
		// }
		// if input == "e" {
		// 	cancel <- 1
		// }

		os.Stdin.Read(make([]byte, 1))

		cancel <- 1
	}()

	listener, err := net.Listen("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Print(err)
			continue
		}
		go handleConn(conn)

		select {
		case <-cancel:
			fmt.Println("server canseled!")
		}
	}
}
func handleConn(c net.Conn) {
	defer c.Close()
	for {
		_, err := io.WriteString(c, time.Now().Format("15:04:05\n\r"))
		if err != nil {
			return
		}
		time.Sleep(1 * time.Second)
	}
}
