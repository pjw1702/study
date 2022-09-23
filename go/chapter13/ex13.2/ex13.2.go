package main

import "fmt"

type User struct {
	Name string
	ID   string
	Age  int
}

type VIPUser struct {
	UserInfo User
	VIPLevel int
	Price    int
}

func main() {
	user := User{"박정우", "pjw", 28}
	vip := VIPUser{
		User{"화랑", "hwarang", 48},
		3,
		250,
	}

	fmt.Printf("유저: %s, ID: %s, 나이: %d\n", user.Name, user.ID, user.Age)
	fmt.Printf("VTP 유저: %s, ID: %s, 나이: %d, VIP레벨: %d, VIP가격: %d만원\n ",
		vip.UserInfo.Name,
		vip.UserInfo.ID,
		vip.UserInfo.Age,
		vip.VIPLevel,
		vip.Price)
}
