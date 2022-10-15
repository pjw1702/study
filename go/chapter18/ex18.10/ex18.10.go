// 요소 삭제
package main

import "fmt"

func main() {
	slice := []int{1, 2, 3, 4, 5, 6}
	idx := 2

	// case 1
	// for i := idx+1; i < len(slice1); i++ {
	// 	slice[i-1] = slice[i]
	// }
	// slice = slice[:len(slice)-1]

	// case 2 copy()함수 사용
	// copy(slice[idx:], slice[idx+1:])
	// slice = slice[:len(slice)-1]

	// case 3 append()함수 사용
	slice = append(slice[:idx], slice[idx+1:]...)

	fmt.Println("slice", slice)
}
