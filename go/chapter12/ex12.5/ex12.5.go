package main

import (
	"fmt"
)

func main() {
	a := [5]int{1, 2, 3, 4, 5}
	b := [5]int{500, 400, 300, 200, 100}
	// b := [5]int64{500, 400, 300, 200, 100} 로 선언한 경우, b = a와 같은 대입 연산은 크기가 다르므로 컴파일 에러가 발생한다

	for i, v := range a {
		fmt.Printf("a[%d] = %d\n", i, v)
	}

	fmt.Println()

	for i, v := range b {
		fmt.Printf("b[%d] = %d\n", i, v)
	}

	/** 변수의 값을 복사하는 경우, 복사하는 변수의 타입의 크기만큼 복사된다(메모리 공간의 크기(타입)가 같아야 한다) */
	/** 요소를 하나씩 복사하는 것이 아닌 메모리의 크기 만큼 통째로 한 번에 복사한다 */
	b = a

	fmt.Println()

	for i, v := range b {
		fmt.Printf("b[%d] = %d\n", i, v)
	}
}
