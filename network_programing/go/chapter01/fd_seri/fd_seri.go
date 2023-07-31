package main

import (
	"fmt"
	"os"
	"syscall"
)

/** 소켓 타입
SOCK_STREAM: TCP
SOCK_DGRAM: UDP
*/

func main() {
	// Create a socket
	fd1, err := syscall.Socket(syscall.AF_INET, syscall.SOCK_STREAM, 0)
	if err != nil {
		fmt.Println("Fail to create socket!: ", err.Error())
		return
	}

	// Open file
	filePath := "test.txt"
	fd2, err := os.OpenFile(filePath, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0666)
	if err != nil {
		fmt.Println("Fail to open the file of test.txt!: ", err.Error())
		return
	}

	// Create another socket
	fd3, err := syscall.Socket(syscall.AF_INET, syscall.SOCK_DGRAM, 0)
	if err != nil {
		fmt.Println("Fail to create socket!: ", err.Error())
		return
	}

	fmt.Printf("file descripter 1: %d\n", fd1)
	fmt.Printf("file descripter 2: %d\n", fd2.Fd())
	fmt.Printf("file descripter 3: %d\n", fd3)

	// Close the file descripters
	syscall.Close(fd1)
	fd2.Close()
	syscall.Close(fd3)
}
