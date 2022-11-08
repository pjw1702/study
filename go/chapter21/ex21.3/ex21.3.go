// C 언어의 함수 포인터와 같다
// 인터페이스는 Low-Level 함수포인터로서 작동하고, 슬라이스는 Low-Level 포인터로서 작동하는 동적 메모리 할당이라고 생각할 수 있다
package main

import "fmt"

func add(a, b int) int {
	return a + b
}

func mul(a, b int) int {
	return a * b
}

func getOperator(op string) func(int, int) int {
	if op == "+" {
		return add // add() 함수의 주소를 반환
	} else if op == "*" {
		return mul // mul() 함수의 주소를 반환
	} else {
		return nil
	}
}

// 함수 타입의 가독성이 떨어지므로 별칭 타입으로 정의하여 사용하면 유용하다
// type OpFn func (int, int) int

func main() {
	// 변수 operator는 func(int, int) int 함수 타입을 가지는 변수이다
	var operator func(int, int) int
	// 함수 타입 변수 operator를 getOperator("+") 함수로 초기화
	operator = getOperator("+")

	// getOperator("+")를 호출하고 add(3, 4)의 결과 값을 반환
	var result = operator(3, 4)
	fmt.Println(result)
}
