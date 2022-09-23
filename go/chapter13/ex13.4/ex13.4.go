package main

import "fmt"

type User struct {
	Name string
	ID   string
	Age  int
}

type VIPUser struct {
	User     // 내장된 필드
	VIPLevel int
	Price    int
	Name     string // 포함하는 구조체인 User의 Name 필드와 똑같은 이름을 가진 필드가 존재하는 경우
}

func main() {

	/*
		user := User{"박정우", "pjw", 28}
		vip := VIPUser{
			User{"화랑", "hwarang", 48},
			3,
			250,
		}

		fmt.Printf("유저: %s, ID: %s, 나이: %d\n", user.Name, user.ID, user.Age)

		// User타입의 필드 없이(생략) 바로 User 타입으로 접근 가능하다
		fmt.Printf("VTP 유저: %s, ID: %s, 나이: %d, VIP레벨: %d, VIP가격: %d만원\n ",
			vip.Name,
			vip.ID,
			vip.Age,
			vip.VIPLevel,
			vip.Price)

		// 유저: 박정우, ID: pjw, 나이: 28
		// VTP 유저: 화랑, ID: hwarang, 나이: 48, VIP레벨: 3, VIP가격: 250만원
	*/

	user := User{"박정우", "pjw", 28}
	vip := VIPUser{
		User{"화랑", "hwarang", 48},
		3,
		250,
		"아무개",
	}

	fmt.Printf("유저: %s, ID: %s, 나이: %d\n", user.Name, user.ID, user.Age)

	/** 포함되는 구조체인 User의 Name 필드와 똑같은 이름을 가진 필드가 존재하는 경우엔 포함하는 구조체의 필드의 우선순위가 더 높다 */
	fmt.Printf("VTP 유저: %s, ID: %s, 나이: %d, VIP레벨: %d, VIP가격: %d만원\n ",
		vip.Name, // 포함되는 구조체의 필드 값을 가져오고 싶을 경우엔 vip.User.Name와 같이 포함되는 구조체의 타입을 삽입하여 준다
		vip.ID,
		vip.Age,
		vip.VIPLevel,
		vip.Price)

}
