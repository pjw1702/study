package main

import "fmt"

/* 부자인 친구가 있는 경우 */
func HashRichFriend() bool {
	return true
}

/* 친구 수 */
func GetFriendCount() int {
	return 3
}

func main() {
	price := 35000 // 음식 값

	if price >= 50000 {
		if HashRichFriend() { // 5만원 이상을 가지고 있는 사람들 중에 부자가 있으면
			fmt.Println("앗 신발끈이 풀렸네")
		} else { // 부자가 없으면
			fmt.Println("나눠내자")
		}
	} else if price >= 30000 { // 5만원 미만이지만 3만원 이상일 때
		if GetFriendCount() > 3 { // 같이간 친구 수가 3명보다 많을 때
			fmt.Println("어이쿠 신발끈이...")
		} else {
			fmt.Println("사람도 얼마 없는데 나눠내자")
		}
	} else {
		fmt.Println("오늘은 내가 쏜다") // 음식값이 3만원 미만일 때
	}
}
