/*재귀 호출*/
/*
	재귀 호출을 할때마다 스택(Stack)이 쌓이는데, 일정 크기의 스택 사이즈를 넘어서면 프로그램이 강제 종료된다(고정 길이 스택)
	Go언어는 자동 증가되는 Stack을 쓴다(고정 길이 스택과는 반대 기술): 메모리가 고갈되지 않는 한 탈출 조건을
	명시해 두지 않으면 계속해서 스택이 쌓임
*/
package main

import "fmt"

func printNo(n int) {
	if n == 0 { // 재귀 호출 탈출 조건(가장 중요한 부분)
		return
	}

	fmt.Println(n)
	printNo(n - 1)          // 재귀 호출
	fmt.Println("After", n) // 재귀 호출 이후 출력
}

func main() {
	printNo(3) // 함수 호출
}
