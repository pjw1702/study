package main

import (
	"fmt"
)

func main() {
	nums := [...]int{10, 20, 30, 40, 50} // [5]int{10, 20, 30, 40, 50}와 같다

	nums[2] = 300

	/** len()함수는 go에 내장된 배열의 길이를 반환하는 함수이다 */
	for i := 0; i < len(nums); i++ {
		fmt.Println(nums[i])
	}
}
