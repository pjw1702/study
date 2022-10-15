/*
malloc vs 슬라이스

C의 malloc 함수는 high-level 포인터로서 배열을 할당하지만 Go의 슬라이스는 low-level 포인터로서 배열을 할당한다

따라서 슬라이스로 인한 배열을 할당할 때의 좌변의 변수는 포인터 변수가 아닌 일반 변수(low-level 포인터 변수)이다

그러므로 함수의 매개변수에 인자로서 슬라이스를 전달할 경우, 포인터 변수가 아닌 일반 변수를 전달하므로 전혀 다른 공간에서 인자 값과 동일한 low-level 포인터 변수와 배열이 복사되어 가리키게 된다
*/

package main

import "fmt"

/*
// case 1. 포인터로 넘기는 경우
func addNum(slice *[]int) {
	*slice = append(*slice, 4)
}

func main() {
	slice := []int{1, 2, 3}
	addNum(&slice)

	fmt.Println(slice)
}
*/

// 값을 넘기는 경우
// 슬라이스는 값을 받아서 값을 반환하는 함수이다 그러므로 case 2의 방법이 슬라이스를 더 슬라이스 답게 쓰는 방법이다
// 포인터와 비교해서 메모리 효율성에 대해 고려해볼 수도 있지만, 성능 차이가 그렇게 많이 나지는 않는다
func addNum(slice []int) []int {
	slice = append(slice, 4)
	return slice // 새로운 슬라이스를 반환
}

func main() {
	slice := []int{1, 2, 3}
	slice = addNum(slice)

	fmt.Println(slice)
}
