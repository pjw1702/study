package main

import "fmt"

func main() {
	var a int = 500
	var p *int

	p = &a

	fmt.Printf("p의 값: %p\n", p) // %p: 포인터 주소
	fmt.Printf("p가 가리키는 메모리의 값: %d\n", *p)

	*p = 100 // p가 가리키는 메모리 주소 공간에 100을 대입하여 a의 값을 변경한다
	fmt.Printf("a의 값: %d", a)
}
