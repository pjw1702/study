package main

import "fmt"

// 슬라이스가 아닌 ...키워드를 통해 가변 인자를 받을 수 있다
func sum(nums ...int) int {
	sum := 0

	fmt.Printf("nums 타입: %T\n", nums)

	for _, v := range nums {
		sum += v
	}
	return sum
}

func main() {
	fmt.Println(sum(1, 2, 3, 4, 5))
	fmt.Println(sum(10, 20))
	fmt.Println(sum())
	// 다음과 같이 fmt 패키지는 소숫점과 문자열 등 어떠한 타입도 입력 가능하다
	// Go Docs에서 참고할 수 있다
	fmt.Println(1, 2, 3, 4, 5, 3.14, "Hello")
}

// 다음과 같은 프로그램이다
// func sum(nums []int) int {
//	sum := 0
//
//	fmt.Printf("nums 타입: %T\n", nums)
//
//	for _, v := range nums {
//		sum += v
//	}
//	return sum
// }
//
// func main() {
//	fmt.Println(sum([]int{1, 2, 3, 4, 5}))
// 	fmt.Println(sum([]int{10, 20}))
//	fmt.Println(sum([]int{}))
// }

// 모든 타입은 빈인터페이스로

// func Print(args ...interface{}) string {			// 모든 타입을 받는 가변 인수
//	for _, arg := range arg {			// 모든 인수 순회
//		// 입력 받은 인수의 타입에 따른 변수 타입 결정
//		switch f := arg.(type) {		// 인수의 타입에 따른 동작
//		case bool:
//			val := arg.(bool)		// 인터페이스 변환
//			...
//		case float64:
//			val := arg.(float64)
//			...
//		case int:
//			val := arg.(int)
//			...
//		}
//		// 다른 타입들도 위와 같이 반복
//	}
// }
