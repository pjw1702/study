package main

import (
	"fmt"
	"time"
)

func PrintHangul() {
	hanguls := []rune{'가', '나', '다', '라', '마', '바', '사'}

	for _, v := range hanguls {
		// 0.3 초마다 한 번씩 출력
		time.Sleep(300 * time.Millisecond) // 1sec = 1000 millisecond
		fmt.Printf(" %c ", v)
	}

}

func PrintNumbers() {
	for i := 1; i <= 5; i++ {
		// 0.4 초마다 한 번씩 출력
		time.Sleep(400 * time.Millisecond)
		fmt.Printf("%d", i)
	}
}

func main() {
	go PrintHangul()
	go PrintNumbers()

	// main() 함수는 3초 동안 sleep
	// 프로그램은 항상 main()에서 시작해서 main()으로 끝난다: 쓰레드가 끝나기 전 까지는 메인 함수가 종료되면 안되므로 sleep을 줘야 함
	time.Sleep(3 * time.Second)
}
