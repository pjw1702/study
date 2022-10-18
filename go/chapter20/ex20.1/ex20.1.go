package main

import "fmt"

type Stringer interface {
	String() string
}

type Student struct {
	Name string
	Age  int
}

func (s Student) String() string {
	return fmt.Sprintf("안녕! 나는 %d살 %s라고 해", s.Age, s.Name) // string 타입의 문자열 반환
}

func main() {
	student := Student{"철수", 12}
	var stringer Stringer

	stringer = student // Go는 원래 강타입이라서 좌변과 우변의 타입이 같아야만 연산이 가능하지만, stringer의 인터페이스에 포함되어 잇는 String() 메서드가 Student 타입의 구조체를 리시버로서 관계를 맺으므로 연산이 가능하다
	// 특정 인터페이스 타입의 변수는 해당 인터페이스가 포함하고 있는 메서드만 호출 가능하다
	fmt.Printf("%s\n", stringer.String())
}
