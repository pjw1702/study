/** 스택 메모리와 힙 메모리 */

package main

import "fmt"

type User struct {
	Name string
	Age  int
}

/**
C/C++ 는, 함수가 끝나는 경우 u 변수는 사라지게 되는데,
스택의 데이터가 pop에 의해 삭제됨에 따라 없어진 변수의 주소(유효하지 않는 주소)를 반환하는 Dangling이 발생한다
*/
func NewUser(name string, age int) *User {
	var u = User{name, age}
	return &u
}

/**
함수내의 지역변수는 스택에 쌓이게 되는데, Go에서는 컴파일러의 Escape Analysising이 발동하여, 함수가 끝나는 시점에 밖으로 탈출하는 인스턴스를 분석한다.
탈출하는 인스턴스를 분석하여 힙 영역에 저장하고, 저장된 인스턴스의 데이터는 쓰임이 다 하면 사라지게 된다
*/
func main() {
	userPointer := NewUser("AAA", 23)
	fmt.Println(userPointer)
}
