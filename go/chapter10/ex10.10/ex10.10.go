package main

import "fmt"

func main() {
	a := 3

	switch a {
	case 1:
		fmt.Println("a == 1")
		break // 값이 일치할 때, Go 언어는 case문에서 break를 선언하지 않아도 자동으로 탈출한다
	case 2:
		fmt.Println("a == 2")
	case 3:
		fmt.Println("a == 3")
		fallthrough // Go 언어는 case문에서 break를 선언하지 않아도 자동으로 탈출하므로, 일치하는 case 뒤(밑)에 또 다른 case도 이어서 실행하고 싶을 때 사용한다
	case 4:
		fmt.Println("a == 4")
	default:
		fmt.Println("a != 1, 2, 3")
	}
}
