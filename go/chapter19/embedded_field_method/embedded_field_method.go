// Go에서는 상속과 메소드 오버라이딩 기능이 없다
// embedded field는 편의상 자동화된 호출일 뿐, 상속과는 다르다
package main

import "fmt"

type User struct {
	Name string
	Age  int
}

func (u User) String() string {
	// formatter를 인자로 받아 string을 생성하여 반환하는 함수
	return fmt.Sprintf("%s, %d", u.Name, u.Age)
}

type VIPUser struct {
	User     // 내장된 필드 (embedded field)
	VIPLevel int
}

// 아래와 같이 선언된 경우, 1차적으로 해당하는 타입의 메소드를 찾지만 없으면 embedded field에 해당하는 타입의 메소드를 찾아서 호출한다
// func (v VIPUser) String() string {
// 	return fmt.Sprintf("%d", v.VIPLevel)
// }

func (v VIPUser) VipLevel() int {
	return v.VIPLevel
}

func main() {
	vip := VIPUser{User{"hwarang", 34}, 5}
	// String() 메소드는 User 타입의 메소드인데 VIPUser 타입의 변수인 vip에 의해 호출되고 있다
	fmt.Println(vip.String()) // 메소드가 속해있는 구조체에 embedded field가 존재하면, embedded field까지 자동으로 호출 가능하다

}
