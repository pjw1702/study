// 함수 리터럴
package main

import "fmt"

type OpFn func(int, int) int

func getOperator(op string) OpFn {
	if op == "+" {
		// 함수를 별도로 선언하지 않고도 바로 반환 값으로서 사용할 수 있다
		return func(a, b int) int {
			return a + b
		}
	} else if op == "*" {
		return func(a, b int) int {
			return a * b
		}
	} else {
		return nil
	}
}

func main() {
	// 변수 operator는 func(int, int) int 함수 타입을 가지는 변수이다
	var operator func(int, int) int
	// 함수 타입 변수 operator를 getOperator("+") 함수로 초기화
	operator = getOperator("+")

	// getOperator("+")를 호출하고 add(3, 4)의 결과 값을 반환
	var result = operator(3, 4)
	fmt.Println(result)
}
