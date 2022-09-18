package main

import (
	"fmt"
	// "time" // time 함수 패키지
)

func main() {

	/*
		i := 1
		for {
			time.Sleep(time.Second) // 스레드를 일시 정지하는 함수(2초동안)
			fmt.Println(i)
			i++
		}
	*/

	for i := 0; i < 10; i++ {
		if i == 3 {
			continue
		}
		if i == 6 {
			break
		}
		fmt.Println("6 *", i, "=", 6*i)
	}
}
