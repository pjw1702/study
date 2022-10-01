package main

import "fmt"

func main() {
	str1 := "Hello"
	str2 := "World"

	// 참고: Go에서는 곱하기나 빼기 등의 연산은 지원되지 않고 오직 더하기 연산만 지원한다
	str3 := str1 + " " + str2
	fmt.Println(str3)

	str1 += " " + str2
	fmt.Println(str1)
}
