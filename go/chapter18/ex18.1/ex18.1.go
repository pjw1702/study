//슬라이스 동작 원리
//문자열과 같은 방식으로, low-level 포인터로서 동작한다
/*
type SliceHeader struct {
	Data uintptr		// 실제 배열을 가리키는 포인터
	Len int				// 요소 개수 (실제 배열의 길이 중 현재 사용하고 있는 배열의 길이)
	Cap int				// 실제 배열의 길이 (현재 사용하고 있는 배열의 최대 길이)
}
*/

package main

import "fmt"

/*
func main() {
	var slice []int // 슬라이스 선언

	if len(slice) == 0 {
		fmt.Println("slice is empty", slice)
	}

	// 런타임 에러 발생
	// slice[1] = 10 에서, var slice []int을 선언만 했을 뿐 초기화 하진 않았다. 한 마디로 슬라이스의 길이가 0인데 1의 인덱스 값을 가지는 배열 요소가 없어서 발생하는 에러이다
	slice[1] = 10
	fmt.Println(slice)
}
*/

func main() {
	// slice := []int{1, 2, 3} // 슬라이스 초기화
	slice := make([]int, 3) // make 함수로 슬라이스 초기화

	if len(slice) == 0 {
		fmt.Println("slice is empty", slice)
	}

	// 런타임 에러 발생
	// slice[1] = 10 에서, var slice []int을 선언만 했을 뿐 초기화 하진 않았다. 한 마디로 슬라이스의 길이가 0인데 1의 인덱스 값을 가지는 배열 요소가 없어서 발생하는 에러이다
	slice[1] = 10
	fmt.Println(slice)
}
