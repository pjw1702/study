package main

import (
	"fmt"
	"net"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Printf("Usage: %s <port>\n", os.Args[0])
	}

	port := os.Args[1]

	// Go 언어에서는 네트워크 소켓을 생성하는 명시적인 함수 호출 및 소켓을 생성하고 설정하는 과정이 필요하지 않다
	// net.Listen() 함수는 내부적으로 소켓을 생성하고, 주어진 포트에서 연결을 수신하기 위한 리스닝 소켓을 반환한다
	// 소켓을 생성한 후 주소를 할당하는 바인딩 과정 또한 명시적으로 처리할 필요가 없다
	// net.Listen() 함수를 호출할 때, 첫 번째 인자로 사용하는 네트워크 타입과 포트 번호를 전달하면 자동으로 주소가 할당된다
	// 즉, Go 언어의 net.Listen() 함수는 주어진 네트워크 타입과 포트 번호로 소켓을 생성하고, 해당 주소에 대한 바인딩도 자동으로 수행한다
	// 굳이 직접 소켓을 직접 생성하고 싶다면, syscall 패키지를 통해 시스템 콜을 하여 수행할 수 있다
	listen, err := net.Listen("tcp", ":"+port) // 소켓 생성 후 바인딩 및 TCP 연결을 수신
	if err != nil {
		fmt.Println("Error listening:", err.Error())
		return
	}
	defer listen.Close()

	fmt.Println("Listening on port", port)

	for {
		conn, err := listen.Accept() // 연결을 허가
		if err != nil {
			fmt.Println("Error accepting connection:", err.Error())
			return
		}

		go handleConnection(conn)
	}
}

func handleConnection(conn net.Conn) {
	defer conn.Close()

	message := "Hello World!"

	_, err := conn.Write([]byte(message)) // 데이터 전달(Hello World! 메세지)
	if err != nil {
		fmt.Println("Error sending:", err.Error())
		return
	}
}
