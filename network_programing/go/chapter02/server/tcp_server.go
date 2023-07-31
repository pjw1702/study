package main

import (
	"fmt"
	"net"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Printf("Usage: %s <port>\n", os.Args[0])
		return
	}

	port := os.Args[1]

	listen, err := net.Listen("tcp", ":"+port)
	if err != nil {
		fmt.Println("Failed to create and listen to socket!: ", err.Error())
		return
	}
	defer listen.Close()

	fmt.Println("Listening on port", port)

	for {
		conn, err := listen.Accept()
		if err != nil {
			fmt.Println("Error accepting connection: ", err.Error())
		}

		go handleConnection(conn)
	}
}

func handleConnection(conn net.Conn) {
	defer conn.Close()

	message := "Hello World!"

	_, err := conn.Write([]byte(message))
	if err != nil {
		fmt.Println("Error send to message: ", err.Error())
		return
	}
}
