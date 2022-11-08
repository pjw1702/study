// 함수 타입(함수 포인터)의 사용: 의존성 주입
// 콜백(CallBack) 함수로서 사용되는 것을 의존성 주입이라고 한다
package main

import (
	"fmt"
	"os"
)

type Writer func(string)
type WriterInterface interface {
	Write(string)
}

// writeHello() 함수를 외부에서 프로그램을 수행하기 위한 로직으로서 주입하여 호출한다: 의존성 주입이라고 한다
func writeHello(writer Writer) {
	writer("Hello World")
}

// 어떠한 프로그램을 수행하기 위한 로직으로서 준비된 함수이지 어떤 특별한 시점에서 수행하는지는 명시되지 않았다
// 의존성 주입을 위한 함수이다
func writeHello2(writer WriterInterface) {
	writer.Write("Hello World")
}

func main() {
	// test.txt 라는 파일을 생성
	f, err := os.Create("test.txt")
	if err != nil {
		fmt.Println("Failed to create a file")
		return
	}
	defer f.Close()

	// 생성한 test.txt라는 파일에 "Hello World" 를 쓴다
	writeHello(func(msg string) {
		fmt.Fprintln(f, msg)
	})
}
