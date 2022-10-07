package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"time"
)

var stdin = bufio.NewReader(os.Stdin) // Reader 객체 생성: input으로 부터 값을 읽어올 수 있는 Reader 객체 생성(input: 표준 입력)

func InputValue() (int, error) {
	var n int
	_, err := fmt.Scanln(&n) // 한 줄을 입력받음
	if err != nil {
		stdin.ReadString('\n') // 표준 입력 문자를 개행 문자가 나올 때 까지 읽어서 비움(입력 버퍼에 남아있는 문자 비움)
	}
	return n, err

}

func main() {
	rand.Seed(time.Now().UnixNano())

	r := rand.Intn(100)
	cnt := 1 // 시도한 횟수 설정

	for { // break 문을 만날 때까지 무한 루프 설정
		fmt.Println("숫자를 입력하세요:")
		n, err := InputValue()
		if err != nil {
			fmt.Println("숫자만 입력하세요!")
		} else {
			if n > r {
				fmt.Println("입력하신 숫자가 더 큽니다.")
			} else if n < r {
				fmt.Println("입력하신 숫자가 더 작습니다.")
			} else {
				fmt.Println("숫자를 맞췄습니다. 축하합니다. 시도한 횟수:", cnt)
				break
			}
			cnt++
		}
	}
}
