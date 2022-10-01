// string 타입의 변수의 문자열 값은 항상 보호 받아야 하므로 슬라이스로 새로운 공간에 같은 값으로 재 할당을 해주어야 한다

package main

import "fmt"

func main() {
	var str string = "Hello World"
	var slice []byte = []byte(str)

	slice[2] = 'a' // 새로 할당된 메모리의 배열에 저장되므로 값이 바뀐다

	// str과 slice는 서로 다른 공간에 메모리가 할당된다
	fmt.Println(str)
	fmt.Printf("%s\n", slice)
}
