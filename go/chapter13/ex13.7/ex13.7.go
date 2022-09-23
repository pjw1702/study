package main

import (
	"fmt"
	"unsafe"
)

/*
type User struct {
	A int8
	B int
	C int8
	D int
	E int8
}
*/

/** 8바이트 보다 작은 메모리의 배치를 나란히하여 메모리 낭비를 줄일 수 있다: 메모리 사용이 작은 프로그램일 경우 유용하다, 큰 프로그램은 사실상 의미없다 */
type User struct {
	A int8
	C int8
	E int8
	B int
	D int
}

func main() {
	user := User{1, 2, 3, 4, 5}
	fmt.Println(unsafe.Sizeof(user))
}
