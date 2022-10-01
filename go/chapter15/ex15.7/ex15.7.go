package main

import (
	"fmt"
)

func main() {
	str := "Hello 월드"

	// len()은 문자 갯수가 아니라 바이트 길이를  반환한다
	// len(str)일 경우, H(1byte) + e(1byte) + l(1byte) + l(1byte) + o(1byte) + ''(1byte) + 월(3byte) + 드(byte)= 12를 반환한다
	// 배열이 아닌 문자열 변수를 순회(for문 등)하는 경우, len()함수의 사용은 적합하지 않다
	for i := 0; i < len(str); i++ {
		fmt.Printf("타입: %T 값: %d 문자 값: %c\n", str[i], str[i], str[i]) // %T: 타입, %c: 각 문자에 대응되는 숫자 값
	}
}
