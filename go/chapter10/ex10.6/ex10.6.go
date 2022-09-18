package main

import "fmt"

/** 나이를 반환하는 함수 */
func getMyAge() int {
	return 22
}

func main() {

	switch age := getMyAge(); age {
	case 10:
		fmt.Println("Teenage")
	case 33:
		fmt.Println("Pair 3")
	default:
		fmt.Println("My age 15", age)
	}

	/** case문 내의 지역변수 age는 탈출 시 삭제되었으므로 main함수 내의 지역변수 age를 초기화 하지 않아 컴파일 에러가 발생한다 */
	// fmt.Println("age is", age)
}
