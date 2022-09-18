package main

import "fmt"

func main() {

	for i := 0; i < 10; i++ {
		fmt.Print(i, ", ")
	}
	/** i 변수를 초기화 하지 않을 시 다음과 같은 경우 for문의 조건문 내에서 초기화한 i는 코드 블럭을 벗어났으므로 컴파일 에러가 발생한다 */
	// fmt.Print(i)
}

/** 다음과 같이 후처리를 설정하지 않고 변수 선언만 해 놓으면 무한히 실행된다 */

/*
func main() {
	i := 0
	for ; i < 10; {
		fmt.Println(i, ", ")
		// i++		// 후처리를 실행 문에서 직접 설정할 수도 있다
	}
	fmt.Println()
	fmt.Print(i)
}
*/
