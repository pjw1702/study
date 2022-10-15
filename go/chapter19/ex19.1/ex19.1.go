package main

import "fmt"

type account struct {
	balance int
}

type myInt int

func (m myInt) AddMethod(a int) myInt {
	rst := int(m) + a // 변수 rst는 int 타입으로 초기화 됨
	return myInt(rst) // int 타입 변수 rst를 myInt 타입으로 타입 변환
}

func addFunc(m myInt, a int) myInt {
	rst := int(m) + a
	return myInt(rst)
}

func withdrawFunc(a *account, amount int) {
	a.balance -= amount
}

// 메소드의 리시버 타입은 반드시 패키지 내에서 선언한 타입(로컬 타입)만 선언 가능하다
// 아래의 int 타입 리시버는 로컬 타입이 아니므로 선언 시 컴파일 에러가 발생한다
// func (c int) Del(b int) int {
//	return c +b
// }

// 메소드는 리시버에 포함되어 있는 기능이라고 생각할 수 있다
// 리시버는 기능적으로는 함수의 첫 번째 매개변수와 똑같은 기능을 하지만, 단지 의미상의 차이가 있을 뿐이다(리시버의 변수의 타입에 따라 어느 타입에 속해있는 메소드인지 지정할 수 있다)
func (a *account) withdrawMethod(amount int) {
	a.balance -= amount
}

func main() {
	// 포인터 구조체 초기화 (변수가 아닌 구조체 자체를 초기화)
	// a 는 포인터 변수지만 100으로 초기화한 구조체 자체를 포인터로서 가리키고 있음
	a := &account{100} // *account

	withdrawFunc(a, 30)

	a.withdrawMethod(30) // 메소드를 입력하면 구조체의 필드가 아닌 구조체 타입에 포함된 메소드가 호출된다

	fmt.Printf("%d\n", a.balance)

	var b myInt

	// addFunc(b, 10)

	b = b.AddMethod(10)

	fmt.Println(b)
}
