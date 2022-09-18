/*Stack 자료구조를 원리로 한 멀티 반환 함수*/
package main

import (
	"fmt"
)

func Divide(a, b int) (int, bool) {
	if b == 0 {
		return 0, false
	}

	return a / b, true
}

func main() {
	c, success := Divide(9, 3) // 호출
	fmt.Println(c, success)
	d, success := Divide(9, 0) // 호출
	fmt.Println(d, success)
}
