package main

import "fmt"

func main() {
	str := "Hello 월드"
	// 길이가 고정된 배열은 변환이 되지 않고, 동적 배열인 슬라이스를 사용해야 변환이 가능하다
	arr := []rune(str) // 슬라이스: 동적 배열, rune 타입: int32의 별칭타입, rune 타입 배열: 한 메모리 공간당 4바트를 할당

	// len()에 변수가 아닌 배열을 인자값으로 전달하면, 배열의 길이를 반환한다
	for i := 0; i < len(arr); i++ {
		fmt.Printf("타입: %T 값: %d 문자 값: %c\n", str[i], str[i], str[i])
	}
}
