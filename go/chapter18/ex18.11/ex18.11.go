// 요소 삽입
package main

import "fmt"

func main() {
	slice := []int{1, 2, 3, 4, 5, 6}
	idx := 2

	// case 1 append()함수를 사용하여 길이를 늘리고 for 문을 사용하여 전체 복사
	// slice = append(slice, 0)

	// for i := len(slice)-2; i >= idx; i-- {
	// 	slice[i+1] = slice[i]
	// }

	// case 2 이중 append() 함수를 사용하여 100으로 시작 + slice의 요소를 길이를 가지는 슬라이스를 idx-1의 길이를 가지는 슬라이스에 append하여 100을 삽입한 길이까지 같이 늘어난다
	// 내부에서 append()함수를 이용하여 임시로 슬라이스를 생성하는 과정에서 메모리를 추가적으로 사용하므로 효율성이 상대적으로 떨어지는 방법이다
	// slice = append(slice[:idx], append([]int{100}, slice[idx:]...)...)

	// case 3
	// 불필요한 메모리 할당이 없다
	slice = append(slice, 0)
	// slice[idx:] 를 긁어서 slice[idx+1:]에 복붙 한다고 생각할 수 있다
	copy(slice[idx+1:], slice[idx:])

	slice[idx] = 100
	fmt.Println(slice)
}
