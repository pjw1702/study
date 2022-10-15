// 슬라이스를 그대로 복사해서 새로운 슬라이스를 만들면 두 슬라이스가 서로 영향을 주지 않는다

package main

import "fmt"

func main() {
	/*
		// 다음과 같은 경우에는 슬라이스가 서로 영향을 준다
		slice1 := []int{1, 2, 3, 4, 5}
		slice2 := slice1[:]

		slice2[1] = 100
		fmt.Println(slice1)
	*/

	/*
		// case 1. make()함수로 초기화 하여 for 문을 이용하여 복사
		slice1 := []int{1, 2, 3, 4, 5}
		slice2 := make([]int, len(slice1))

		// slice1을 복사하여 새로운 배열을 할당하여 slice2가 포인터로서 가리키게 된다
		for i, v := range slice1 {
			slice2[i] = v
		}
	*/

	/*
		// case 2. append()함수를 이용하여 빈 슬라이스를 생성한 뒤 slice1을 전체 복사
		slice1 := []int{1, 2, 3, 4, 5}
		slice2 := append([]int{}, slice1...)
	*/

	// case 3. make()함수로 초기화 하여 copy()함수를 이용하여 복사
	slice1 := []int{1, 2, 3, 4, 5}
	slice2 := make([]int, len(slice1))
	// copy(dest,src)
	copy(slice2, slice1) // slice1의 내용을 slice2로 전체 복사

	slice2[1] = 100

	fmt.Println("slice1", slice1)
	fmt.Println("slice2", slice2)
}
