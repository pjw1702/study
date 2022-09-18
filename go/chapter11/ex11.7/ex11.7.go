package main

import (
	"fmt"
)

/** 플래그 변수를 활용하는 경우 */

func main() {
	a := 1
	b := 1
	found := false

	for ; a <= 9; a++ {
		for b = 1; b <= 9; b++ {
			if a*b == 45 {
				found = true // 플래그 변수: 특정 조건 만족 시, 속해있는 for문의 블럭 뿐만 아니라 중첩된 for문의 블럭 전체를 탈출하고 싶을 때 이용한다
				break
			}
		}
		if found { // 플래그 변수를 이용하여 중첩된 for문 블럭까지 탈출한다
			break
		}
	}
	fmt.Printf("%d * %d = %d\n", a, b, a*b)
}

/** 레이블을 활용하는 경우 */

/*
func main() {
	a := 1
	b := 1

OuterFor: // Label(레이블)
	for ; a <= 9; a++ {
		for b = 1; b <= 9; b++ {
			if a*b == 45 {
				break OuterFor // 레이블로 지정해 놓은 코드 블럭을 탈출한다
			}
		}
	}
	fmt.Printf("%d * %d = %d\n", a, b, a*b)

}
*/
