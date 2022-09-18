/* 방에 전등을 켜고 끄는 상황의 예제 */
/* 비트 플래그는 비트에 따라 의미를 부여하기 위한 방법이다
해당 비트가 0이냐 1이냐에 따라 다르게 동작한다
변수의 많은 사용을 피함으로써 메모리 효율성과 가독성을 높일 수 있다*/
package main

import "fmt"

/* 각 방을 나타내는 비트 플래그 */
const (
	MasterRoom uint8 = 1 << iota
	LivingRoom
	BathRoom
	SmallRoom
)

/* 전등을 켜는 상황의 함수 */
func SetLight(rooms, room uint8) uint8 {
	return rooms | room
}

/* 불을 끄는 상황의 함수 */
func ResetLight(rooms, room uint8) uint8 {
	return rooms &^ room
}

/* 전등이 켜져있는 경우의 함수*/
func IsLightOn(rooms, room uint8) bool {
	return rooms&room == room
}

/* 전등을 켜는 함수 */
func TurnLights(rooms uint8) {
	if IsLightOn(rooms, MasterRoom) {
		fmt.Println("안방에 전등을 켠다")
	}
	if IsLightOn(rooms, LivingRoom) {
		fmt.Println("거실에 전등을 켠다")
	}
	if IsLightOn(rooms, BathRoom) {
		fmt.Println("화장실에 전등을 켠다")
	}
	if IsLightOn(rooms, SmallRoom) {
		fmt.Println("작은 방에 전등을 켠다")
	}
}

func main() {
	var rooms uint8 = 0                 // rooms: 0000 0000
	rooms = SetLight(rooms, MasterRoom) // rooms: 0000 0001
	rooms = SetLight(rooms, BathRoom)   // rooms: 0000 0101
	rooms = SetLight(rooms, SmallRoom)  // rooms: 0000 0111

	rooms = ResetLight(rooms, SmallRoom) // rooms: 0000 0101

	TurnLights(rooms) // 0000 0001에 해당하는 안방과 0000 0100에 하당하는 화장실에만 전등을 켠다
}
