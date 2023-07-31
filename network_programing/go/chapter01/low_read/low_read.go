package main

import (
	"fmt"
	"os"
)

func main() {
	filePath := "data.txt"
	file, err := os.Open(filePath) // Open()함수: 내부적으로 O_RDONLY 플래그를 사용하여 읽기 전용으로 연다
	if err != nil {
		fmt.Println("Fail to open The file of data.txt!: ", err.Error())
		return
	}
	defer file.Close()

	message_buffer := make([]byte, 100)
	letters_size, err := file.Read(message_buffer)
	if err != nil {
		fmt.Println("Fail to read The file of data.txt!: ", err.Error())
		return
	}

	fmt.Printf("file descripter: %d \n", file.Fd())
	fmt.Printf("Read %d bytes: %s \n", letters_size, message_buffer[:letters_size])
}
