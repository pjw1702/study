package main

import (
	"fmt"
)

func f() {
	fmt.Println("f() 함수 시작")
	defer func() {
		// 패닉이 있을 시에
		if r := recover(); r != nil {
			// r = panic()의 인자 값인 "제수는 0일 수 없습니다"
			fmt.Println("panic 복구 - ", r)
		}
	}() // 호출

	g()
	fmt.Println("f() 함수 끝")
}

func g() {
	fmt.Printf("9/3 = %d\n", h(9, 3))
	fmt.Printf("9/0 = %d\n", h(9, 0))
}

func h(a, b int) int {
	if b == 0 {
		panic("제수는 0일 수 없습니다")
	}
	return a / b
}

func main() {
	f()
	fmt.Println("프로그램이 계속 실행 됨")
}

/* 패닉의 전파: 패닉(충격)은 콜 스택 역순으로 전파

* main() -> f() -> g() -> h():패닉!

* h() -> g(): 복구 확인 -> f(): 복구 확인 -> main(): 복구 확인
	* 복구가 안되면 프로그램 종료

* 복구는 recover()
	ㄴ 반환 값: panic 함수에 넣은 인자 값

	* func recover() interface{}

	* 패닉 객체를 반환

	* Defer와 함께 사용
*/

/* 출력 결과

f() 함수 시작
9/3 = 3
panic 복구 - 제수는 0일 수 없습니다
프로그램이 계속 실행 됨
*/

/* 복구 처리를 하지 않았을 경우에 출력 결과

f() 함수 시작
9/3 = 3
panic: 제수는 0일 수 없습니다

goroutine 1 [running]:
main.h(...)
        C:/Users/박정우/Documents/workspace/github/study/go/chapter23/ex23.6/ex23.6.go:28
main.g()
        C:/Users/박정우/Documents/workspace/github/study/go/chapter23/ex23.6/ex23.6.go:23 +0x79
main.f()
        C:/Users/박정우/Documents/workspace/github/study/go/chapter23/ex23.6/ex23.6.go:17 +0x5b
main.main()
        C:/Users/박정우/Documents/workspace/github/study/go/chapter23/ex23.6/ex23.6.go:34 +0x19
*/
