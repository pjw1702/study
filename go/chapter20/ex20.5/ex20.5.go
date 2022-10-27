// 빈 인터페이스(interface{})는 모든 타입이 가능
// 모든 타입을 매개변수로 선언 또는 인자로 전달 가능하다
// fmt.Print 함수 또한 빈 인터페이스(interface{})를 인자로 받는다
package main

import "fmt"

func PrintVal(v interface{}) {
	switch t := v.(type) {
	case int:
		fmt.Printf("v is int %d\n", int(t))
	case float64:
		fmt.Printf("v is int %d\n", float64(t))
	case string:
		fmt.Printf("v is int %d\n", string(t))
	default:
		fmt.Printf("Not supported type %T:%v\n", t, t)
	}
}

type Student struct {
	Age int
}

func main() {
	PrintVal(10)
	PrintVal(3.14)
	PrintVal("Hello")
	PrintVal(Student{15})
}
