package main

import (
	"fmt"
	"sync"
	"time"
)

func square(wg *sync.WaitGroup, ch chan int) {
	for n := range ch {
		fmt.Println("Square: ", n*n)
		time.Sleep(time.Second)
	}
	wg.Done()
}

func main() {
	var wg sync.WaitGroup
	ch := make(chan int)

	wg.Add(1)
	go square(&wg, ch)

	for i := 0; i < 10; i++ {
		ch <- i * 2
	}
	close(ch) // 무한 루프로 닫혀있지 않은 채널에 데이터를 저장하면 좀비 고루틴이 되어 데드락이 발행하므로 채널을 닫아주는 작업을 수행해 주어야 한다
	wg.Wait()
}
