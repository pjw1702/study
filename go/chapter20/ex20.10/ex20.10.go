package main

import "fmt"

type Reader interface {
	Read()
}

type Closer interface {
	Close()
}

type File struct {
}

func (f *File) Read() {
	fmt.Println("Read()")
}

// 다음 메서드를 정의하지 않으면 File 타입에 Close() 메소드가 존재하지 않으므로 타입 변환하고 실행(ReadFile 함수) 시 런타임 에러 발생
func (f *File) Close() {
	fmt.Println("Close()")
}

func ReadFile(reader Reader) {
	// c, ok := reader.(Closer) // Closer 인터페이스 타입으로 타입 변환 (서로 다른 인터페이스 타입으로 변환 가능)
	// Close 메소드가 없다면 <nil>, false가 출력된다
	// fmt.Println(c, ok)
	// if ok {
	//	c.Close()
	// }
	// 위의 실행 문을 관용적 표현으로서 아래와 같이 줄여서 작성할 수 있다
	if c, ok := reader.(Closer); ok {
		fmt.Println(c, ok)
		c.Close()
	}
	// c.Close()
}

func main() {
	file := &File{}
	ReadFile(file)
}
