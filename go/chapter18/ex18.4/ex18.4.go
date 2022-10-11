package main

import "fmt"

func changeArray(array2 [5]int) {
	array2[2] = 200
}

func changeSlice(slice2 []int) {
	slice2[2] = 200
}

func main() {
	array := [5]int{1, 2, 3, 4, 5}
	slice := []int{1, 2, 3, 4, 5}

	changeArray(array)
	changeSlice(slice)

	// C/C++은 함수의 매개변수에 인자로서 배열을 전달할 경우 배열의 메모리 주소값을 전달하므로 함수가 종료되어도 배열의 값에 변경된 값이 저장된다
	// Go에서 배열은 함수 인자로서 넘겨줘도 R-value로 동작이지만 슬라이스는 포인터이다
	fmt.Println(array)
	fmt.Println(slice)
}
