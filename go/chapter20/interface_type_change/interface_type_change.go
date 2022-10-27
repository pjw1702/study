package main

import "fmt"

type Attacker interface {
	Attack()
}

type Monster struct {
	Lv int
}

func (m *Monster) Attack() {
	fmt.Println("Monster Attack")
}

func DoAttack(att Attacker) {
	if att != nil {
		att.Attack()

		var monster *Monster
		monster = att.(*Monster) // 인터페이스의 타입 변환
		fmt.Println(monster.Lv)
	}
}

func main() {
	DoAttack(&Monster{20})
}

// 다음과 같은 경우 런타임 에러가 발생한다

// type Attacker interface {
//	Attack()
// }

// type Monster struct {
//	Lv int
// }

// type Player struct {
//	Lv int
// }

// func (m *Player) Attack() {
//	fmt.Println("Player Attack")
// }

//func (m *Monster) Attack() {
//	fmt.Println("Monster Attack")
// }

// func DoAttack(att Attacker) {
//	if att != nil {
//		att.Attack()

//		var monster *Monster
//		monster = att.(*Monster)		// 인터페이스의 타입 변환을 하지 못한다: 매개 변수의 인자로 Player 타입을 포인터로 받았기 때문이다
//		fmt.Println(monster.Lv)
//	}
// }

// func main() {
//	DoAttack(&Player{20})
// }
