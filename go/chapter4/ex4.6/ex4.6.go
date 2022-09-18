package main

import "fmt"

var g int = 10 // 전역 변수: g변수의 범위는 패키지 전역에서 영향을 미침

func main() {
	var m int = 20

	{
		var s int = 50
		fmt.Println(m, s, g)
	}

	m = s + 20 // s 변수의 범위(중괄호 블록)를 벗어났으므로 범위 밖에서의 s의 초기화를 다시 해야 함
}
