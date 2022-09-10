package main

import (
	"fmt"
)

func Add(a int, b int) int {
	return a + b
}

func main() {
	c := Add(3, 6)
	fmt.Println(c)
}

/*
// C/C++과는 달리 함수의 순서를 바꾸어도 정상 작동 한다
package main

import (
	"fmt"
)

func main() {
	c := Add(3, 6)
	fmt.Println(c)
}

func Add(a int, b int) int {
	return a + b
}
*/
