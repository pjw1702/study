package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func SumAtoB(a, b int) {
	sum := 0
	for i := a; i <= b; i++ {
		sum += i
	}
	fmt.Printf("%d 부터 %d까지 합계는 %d입니다\n", a, b, sum)
	// 고루틴 실행 완료: 완료될 때마다 고루틴이 하나씩 줄어듦
	wg.Done()
}

func main() {
	// 10개의 고루틴 작업을 기다림
	wg.Add(10)
	// 10 개의 고루틴 생성
	for i := 0; i < 10; i++ {
		go SumAtoB(1, 1000000000)
	}
	// 10개의 고루틴이 모두 실행 완료될 때까지 기다림
	wg.Wait()
}
