package main

import (
	"fmt"
	"unsafe" // 안전하지 않은 함수를 제공해주는 패키지
)

type User struct {
	Age   int32   // 4바이트
	Score float64 // 8바이트
}

func main() {

	user := User{23, 77.2}
	fmt.Println(unsafe.Sizeof(user)) // Sizeof(): 변수의 메모리 공간의 크기를 반환해주는 함수
}
