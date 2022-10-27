package main

import (
	"fmt"
	"unsafe" // sizeof() 함수와 같은 타입의 크기를 가져오기 위한 패키지
)

type Stringer interface {
	String() string
}

type Student struct {
	Name string
}

func (s *Student) String() string {
	return s.Name
}

type User struct {
	Name string
	Age  int
}

func (u User) String() string {
	return u.Name
}

func main() {
	var stringer1 Stringer
	fmt.Printf("type: %T size: %d\n", stringer1, unsafe.Sizeof(stringer1))

	student := &Student{"AAA"}
	stringer1 = student // Stringer 인터페이스에 대입
	fmt.Printf("type: %T size: %d\n", stringer1, unsafe.Sizeof(stringer1))

	var stringer2 Stringer
	fmt.Printf("type: %T size: %d\n", stringer2, unsafe.Sizeof(stringer2))

	user := User{"BBB", 20}
	stringer2 = user // Stringer 인터페이스에 대입
	fmt.Printf("type: %T size: %d\n", stringer2, unsafe.Sizeof(stringer2))

}

// 실행 결과
// type: <nil> size: 16
// type: *main.Student size: 16
// type: <nil> size: 16
// type: main.User size: 16
