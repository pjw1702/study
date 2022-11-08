package main

import "fmt"

type Stringer interface {
	String() string
}

type Student struct {
}

// 다음과 같은 메서드를 선언하지 않고 타입 변환을 시도할 시, 컴파일 에러가 발생한다
// 다음 메서드를 선언하지 않으면 stringer() 메서드가 Student 타입을 가리키고 있지 않으므로 변환 불가
func (s *Student) String() string {
	return "Student"
}

func main() {
	var stringer Stringer

	s := stringer.(*Student) // stringer를 Student 포인터 타입으로 변환한다
	fmt.Println(s)
}
