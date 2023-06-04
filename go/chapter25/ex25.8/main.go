package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

func PrintEverySecond(ctx context.Context) {
	tick := time.Tick(time.Second)
	for {
		select {
		case <-ctx.Done(): // Done(): cancel() 함수 호출로 인해 작업이 끝날 때, cancel() 함수가 보낸 종료 시그널을 ctx.Done()의 채널에 받아서 종료(종료 지시를 받음)
			wg.Done()
			return
		case <-tick:
			fmt.Println("tick")
		}
	}
}

func main() {
	wg.Add(1)
	// Background(): 기본 컨텍스트 할당(컨텍스트는 항상 초기에 백그라운드에 할당받은 기본 컨텍스트 위에 덮어 씌워서 실행되는 구조로 설계되어 있다)
	// WithCancel(): 할당 받은 기본 컨텍스트에 취소 기능을 추가해 주는 함수
	ctx, cancel := context.WithCancel(context.Background())
	go PrintEverySecond(ctx)
	time.Sleep(5 * time.Second)
	cancel() //

	wg.Wait()
}
