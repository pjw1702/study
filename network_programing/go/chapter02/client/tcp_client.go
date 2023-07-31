package main

import (
	"fmt"
	"log"
	"net"
	"os"
)

func main() {
	if len(os.Args) != 3 {
		fmt.Printf("Usage: %s <IP> <port>\n", os.Args[0])
		return
	}

	conn, err := net.Dial("tcp", fmt.Sprintf("%s:%s", os.Args[1], os.Args[2]))
	if err != nil {
		log.Println("Error connect to server!: ", err.Error())
		return
	}
	defer conn.Close()

	message_buffer := make([]byte, 30)
	strLen := 0

	// 서버에서는 데이터를 한 번에 전달하였지만, 클라이언트에서는 1바이트씩 쪼개서 수신: 전송되는 데이터의 경계(Boundary)가 존재하지 않음
	for {
		readLen, err := conn.Read(message_buffer[strLen : strLen+1]) // 읽은 바이트 수 반환
		if err != nil {
			log.Println("Error read to message!: ", err.Error())
		}

		if readLen == 0 {
			fmt.Println("Letters of message is not exist.")
			break
		}

		strLen += readLen
	}

	fmt.Printf("Message from server: %s\n", message_buffer)
	fmt.Printf("Function read call count: %d\n", strLen)
}
