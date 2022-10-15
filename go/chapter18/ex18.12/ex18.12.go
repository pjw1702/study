// 슬라이스 정렬
package main

import (
	"fmt"
	"sort"
)

/*
// int 타입 슬라이스의 정렬
func main() {
	slice := []int{5, 2, 6, 3, 1, 4}
	sort.Ints(slice)		// int 타입 슬라이스의 정렬을 위한 함수
	fmt.Println(slice)
}
*/

// 구조체 타입 슬라이스의 정렬
// primitive 타입(기본 타입)의 정렬 함수는 Ints, Float64s, Strings와 같은 이름의 함수로 제공이 되지만 구조체 타입은 메소드와 인터페이스를 생성하여 sort.Sort() 함수의 인자로서 넘겨줘야 한다

type Student struct {
	Name string
	Age  int
}

type Students []Student

// 정렬 인터페이스를 생성하기 위한 메소드 생성
func (s Students) Len() int           { return len(s) }
func (s Students) Less(i, j int) bool { return s[i].Age < s[j].Age }
func (s Students) Swap(i, j int)      { s[i], s[j] = s[j], s[i] }

func main() {
	s := []Student{
		{"화랑", 31},
		{"백두산", 52},
		{"류", 42},
		{"켄", 38},
		{"송하나", 18},
	}

	// sort.Sort(s)				// 컴파일 에러 발생: sort.Sorts()함수는 슬라이스가 아닌 인터페이스를 인자로서 받는다. 인터페이스 내에서도 Len(), Less(), Swap() 메소드를 가지고 있는 타입의 인터페이스를 인자로서 받는다

	sort.Sort(Students(s)) // []Students 타입의 구조체 슬라이스 변수 s를 Students 타입으로 변환
	fmt.Println(s)
}
