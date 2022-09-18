/*출력 값에 이름 지정*/
package main

import "fmt"

func Divide(a, b int) (result int, success bool) {
	if b == 0 {
		result = 0      // return 값 1 이름
		success = false // return 값 2의 이름
		return          // result의 값과 success의 값을 순서대로 반환 함
	}

	result = a / b
	success = true
	return
}

func main() {
	c, success := Divide(9, 3)
	fmt.Println(c, success)
	d, success := Divide(9, 3)
	fmt.Println(d, success)

}
