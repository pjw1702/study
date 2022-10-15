package main

import "fmt"

func main() {
	array := [100]int{1: 1, 2: 2, 99: 100}
	slice1 := array[1:10]

	// 특정 배열을 슬라이싱을 한 슬라이스를 슬라이싱한 경우 두 번째 슬라이스는 첫 번째 슬라이스를 가리키는 것이 아니라 첫 번째 슬라이스가 가리키는 배열을 가리키게 된다
	// 특정 배열을 슬라이싱한 슬라이스도 길이는 m-n 이지만, cap은 배열의 전체 길이 - 시작 인덱스 이므로 따지고 보면 당연한 결과이다
	slice2 := slice1[2:99]
	fmt.Println(slice1)
	fmt.Println(slice2)
}
