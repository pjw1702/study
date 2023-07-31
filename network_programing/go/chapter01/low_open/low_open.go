package main

import (
	"fmt"
	"os"
)

func main() {
	filePath := "data.txt"
	file, err := os.OpenFile(filePath, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0644) // O_TRUNC: 파일을 열 때, 이미 존재하는 경우 해당 파일을 길이를 0으로 잘라냄(O_TRUNC 플래그는 파일을 열 때 O_WRONLY 또는 O_RDWR 플래그와 함께 사용되어야 한다. 읽기 전용(O_RDONLY)으로 파일을 열 때는 O_TRUNC 플래그를 사용할 수 없다)
	if err != nil {
		fmt.Println("The file of data.txt open fail!: ", err.Error())
	}
	defer file.Close()

	message_buffer := []byte("Let's go! \n")
	if _, err := file.Write(message_buffer); err != nil {
		fmt.Println("Fail to write to The file of data.txt!: ", err.Error())
	}

	fmt.Printf("file descripter: %d \n", file.Fd())
}
