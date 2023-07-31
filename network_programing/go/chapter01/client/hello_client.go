package main

import (
	"fmt"
	"net"
	"os"
)

func main() {
	if len(os.Args) != 3 {
		fmt.Printf("Usage: %s <IP> <port>\n", os.Args[0])
		return
	}

	ip := os.Args[1]
	port := os.Args[2]

	conn, err := net.Dial("tcp", ip+":"+port)
	if err != nil {
		fmt.Println("Erorr connecting: ", err.Error())
		return
	}
	defer conn.Close()

	message := make([]byte, 30)
	n, err := conn.Read(message)
	if err != nil {
		fmt.Println("Error reading: ", err.Error())
		return
	}

	fmt.Printf("Message from server: %s\n", string(message[:n]))
}
