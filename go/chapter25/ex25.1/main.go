package main

import (
	"fmt"
	"sync"
	"time"
)

// 채널을 사용하면 고루틴이 간의 데이터를 주고 받을 때 뮤텍스를 사용해서 락을 걸 필요가 없다
func square(wg *sync.WaitGroup, ch chan int) { // 채널 인스턴스는 그 자체만으로 call by reference 타입이므로 포인터 타입으로 인자를 넘겨줄 필요가 없다
	n := <-ch

	time.Sleep(time.Second)
	fmt.Println("Square:", n*n)
	wg.Done()
}

func main() {
	var wg sync.WaitGroup
	ch := make(chan int) // 채널 인스턴스: 힙 메모리에 할당되는 채널을 참조하는 call by reference 타입

	// square 고루틴은 함수나 메서드 호출이 아닌 main 고루틴 밖에서 실행되는 '스레드'이므로 square라는 스레드가 실행되는 동안 채널을 통해 9라는 데이터가 전송되면 Done()에 의해 종료된다
	// main 고루틴과 square 고루틴(서로 다른 스레드) 사이에서 채널을 통해 9라는 데이터를 주고 받을 수 있다
	wg.Add(1)
	go square(&wg, ch)
	ch <- 9
	wg.Wait()
}
