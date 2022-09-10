package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	/* var stdin = bufio.NewReader() */
	stdin := bufio.NewReader(os.Stdin)

	var a int
	var b int

	n, err := fmt.Scanln(&a, &b)
	/*
		첫 문자를 읽었을 때 에러가 발견이 되면 그 다음 문자는 굳지 읽지 않아도 변수의 데이터 자체가 에러이므로
		다음 변수의 데이터를 읽기위해 현재 입력 버퍼의 문자들을 모두 읽어서(꺼내서) 비워줘야 한다.
		입력 버퍼를 표준 입력에서 읽어오는 함수 ReadString()를 이용하여 '\n'이 나올 때까지 모두 비운다
	*/
	if err != nil { // 아무것도 아닌(올바르지 않은) 메모리 주소값을 가지는 경우
		fmt.Println(err)
		stdin.ReadString('\n')
	} else {
		fmt.Println(n, a, b)
	}
	/* 입력 버퍼를 모두 지우고 새롭게 a, b 변수에 입력을 할 수 있다 */
	n, err = fmt.Scanln(&a, &b)
	if err != nil { // 아무것도 아닌(올바르지 않은) 메모리 주소값을 가지는 경우
		fmt.Println(err)
		stdin.ReadString('\n')
	} else {
		fmt.Println(n, a, b)
	}
}
