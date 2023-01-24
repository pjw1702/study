package main

import (
	"fmt"
)

func divide(a, b int) {
	if b == 0 {
		panic("b는 0일 수 없습니다")
	}
	fmt.Printf("%d / %d = %d\n", a, b, a/b)
}

func main() {
	divide(9, 3)
	divide(9, 0)
}

/* 출력결과
9 / 3 = 3
panic: b는 0일 수 없습니다

goroutine 1 [running]:
main.divide(0x9?, 0x3?)
        C:/Users/박정우/Documents/workspace/github/study/go/chapter23/ex23.5/ex23.5.go:9 +0x105
main.main()
        C:/Users/박정우/Documents/workspace/github/study/go/chapter23/ex23.5/ex23.5.go:16 +0x31
*/

/* 패닉(Panic)의 장점
에러가 발생한 시점에서 프로그램이 죽고, 어느 부분에서 에러가 발생했는지 알 수 있다
*/
