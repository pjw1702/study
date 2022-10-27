package main

type Stringer interface {
	String() string
}

type Student struct {
}

func main() {
	var stringer Stringer
	// 다음과 같이 타입 변환을 시도할 시,
	stringer.(*Student) // stringer를 Student 포인터 타입으로 변환한다
}
