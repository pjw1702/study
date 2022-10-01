package main

import (
	"fmt"
	"reflect"
	"unsafe"
)

func main() {
	str1 := "Hello 월드"
	str2 := str1

	// unsafe.Pointer(&str1): str1을 row 포인터 변수로 변환(*str1)
	// *reflect.StringHeader: StringHeader 구조체는 reflect 패키지 내에 내장되어 있는데, 변환된 string 타입의 row 포인터 변수 str1을 StringHeader 구조체로 타입 변환을 한다
	stringHeader1 := (*reflect.StringHeader)(unsafe.Pointer(&str1)) // StringHeader 타입의 포인터 변수
	stringHeader2 := (*reflect.StringHeader)(unsafe.Pointer(&str2)) // StringHeader 타입의 포인터 변수

	fmt.Println(stringHeader1)
	fmt.Println(stringHeader2)

}
