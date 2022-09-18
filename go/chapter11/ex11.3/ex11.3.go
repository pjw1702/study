package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	stdin := bufio.NewReader(os.Stdin)

	for {

		fmt.Println("숫자를 입력하세요")

		var number int
		_, err := fmt.Scanln(&number) // '_' 변수: 반환한 return 값을 쓰지 않을 때

		if err != nil {
			fmt.Println("숫자로 입력하세요")

			// 키보드 버퍼를 지웁니다.
			stdin.ReadString('\n') // 개행 문자가 나올 때까지 입력한 문자 버퍼를 읽어서 지움
			continue
		}
		fmt.Println("입력하신 숫자는 %d입니다. \n", number)

		if number%2 == 0 {
			break // 입력한 숫자가 짝수인지 검사를 합니다.
		}
	}
	fmt.Println("for문이 종료되었습니다.")
}
