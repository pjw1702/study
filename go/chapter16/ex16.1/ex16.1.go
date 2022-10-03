package main

import (
	"fmt"
	"math/rand" // 패키지 import: main이 아닌 외부 패키지를 가져옴
)

func main() {
	fmt.Println(rand.Int()) // math의 rand 패키지에 정의 된 함수를 사용
}
