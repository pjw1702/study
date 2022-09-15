/* if문의 조건문 판별을 위해 호출된 함수는 값 비교 목적 외에 조작이나 연산 등에 사용되어서는 안된다 */

package main

import "fmt"

var cnt int = 0

func IncreaseAndReturn() int {
	fmt.Println("IncreaseAndReturn()", cnt)
	cnt++
	return cnt
}

func main() {
	if false && IncreaseAndReturn() < 5 { // 쇼트 서킷: 좌변이 false 이므로 IncreaseAndReturn() 함수의 조건문은 검사조차 안한다(호출 X)
		fmt.Println("1 증가")
	}

}
