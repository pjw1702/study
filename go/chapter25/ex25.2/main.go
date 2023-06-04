package main

import (
	"fmt"
	"time"
)

// func square() {
// 	// 무한 루프
// 	for {
// 		time.Sleep(2 *time.Second)
// 		fmt.Println("sleep")
// 	}
// }

// func main() {
// 	ch := make(chan int)	// square 고루틴이 채널에 저장된 9를 빼는 로직없이 무한루프로 실행되고 있으므로 main 고루틴에서 생성한 채널에 9를 저장했지만 빠지지가 않으므로 해당 채널 인스턴스 ch는 무한히 대기한다

// 	go square()
// 	ch <- 9
// 	fmt.Println("Never print")
// }

func square() {
	// 무한 루프
	for {
		time.Sleep(2 * time.Second)
		fmt.Println("sleep")
	}
}

func main() {
	ch := make(chan int, 2) // square 고루틴이 채널에 저장된 9를 빼는 로직없이 무한루프로 실행되고 있지만 main 고루틴에서 생성한 여유 데이터 공간이 있는 채널에 9를 저장했으므로 빠지지가 않아도 해당 채널 인스턴스 ch는 대기하지 않는다

	go square()
	ch <- 9
	fmt.Println("Never print")
}
