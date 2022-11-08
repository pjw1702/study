package main

import (
	"fmt"
	"os"
)

func main() {
	f, err := os.Create("test.txt")
	if err != nil {
		fmt.Println("Failed to create a file", err)
		return
	}

	// defer 명령문으로 작성되었으므로 실행이 지연되어 함수가 끝나기 직전에 수행한다
	defer fmt.Println("반드시 호출됩니다")
	defer f.Close() // 함수가 종료 전에 파일 핸들러 반환
	defer fmt.Println("파일을 닫았습니다")

	fmt.Println("파일에 Hello World를 씁니다")
	// Standard I/O가 아닌 File I/O 스트림으로 지정
	// 출력 스트림을 파일로 지정
	fmt.Fprintln(f, "Hello World")
}
