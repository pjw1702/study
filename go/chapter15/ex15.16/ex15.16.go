// string 타입과 slice가 서로 다른 공간의 메모리를 할당 받았음을 내부적으로 확인하기 위함이다

package main

import (
	"fmt"
	"reflect"
	"unsafe"
)

func main() {
	var str string = "Hello World"
	var slice []byte = []byte(str)

	stringHeader := (*reflect.StringHeader)(unsafe.Pointer(&str))
	sliceHeader := (*reflect.SliceHeader)(unsafe.Pointer(&slice))

	fmt.Printf("str:\t%x\n", stringHeader.Data)
	fmt.Printf("slice:\t%x\n", sliceHeader.Data)
}
