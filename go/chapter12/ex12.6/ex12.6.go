package main

import (
	"fmt"
)

func main() {
	a := [2][5]int{
		{1, 2, 3, 4, 5},
		{5, 6, 7, 8, 9}, // go에서는 중괄호가 다른 줄에서 닫히면 줄바꿈 기능으로 인해 마지막 요소라도 ,를 입력해 주어야 한다
	}

	/*
		a := [2][5]int{
			{1, 2, 3, 4, 5},
			{5, 6, 7, 8, 9} }	// 같은 줄에서 중괄호가 닫혔으므로 컴파일 에러가 발생하지 않는다
	*/

	for _, arr := range a {
		for _, v := range arr {
			fmt.Print(v, " ")
		}
		fmt.Println()
	}
}
