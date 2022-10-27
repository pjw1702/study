// 인터페이스는 인스턴스 메모리 주소 + 타입 정보를 멤버로 갖는 구조체 형식이다
// 인터페이스가 가지는 메모리 주소는 정의한 메서드의 리시버 타입(구조체)의 메모리 주소이고, 타입 정보는 리시버 타입의 정보이다
// 인터페이스는 low-level 함수 포인터라고 볼 수 있다
// 제공하는 메서드의 이름이 갑작스레 변경되는 경우도 있다

package main

import "fmt"

type Database interface {
	Get()
	Set()
}

type CDatabase struct {
}

// 인터페이스에서 제공하는 메서드 명과 특정 패키지에서 제공하는 메서드 명이 일치하지 않을 경우
func (c CDatabase) GetData() {
}

func (c CDatabase) SetData() {
}

// 어댑터 생성
type Wrapper struct {
	cdb CDatabase
}

// 어댑터를 리시버로 받는 메서드 선언
func (w Wrapper) Get() {
	w.cdb.GetData()
}

func (w Wrapper) Set() {
	w.cdb.SetData()
}

func main() {
	// Get(), Set()메서드 호출 시 cdb 필드의 GetData(),SetData() 메서드를 wrapping 한다(바인딩)
	var cdatabase CDatabase
	var database Database         // 인터페이스 초기화
	database = Wrapper{cdatabase} // (CDatabase라는 패키지내에서 제공한다고 가정)CDatabase 타입의 메서드 주소 값을 인터페이스로 전달
	database.Get()
	database.Set()

	fmt.Println(database)
}
