package main

import "fmt"

func main() {

	day := "thursday"

	switch day {
	// 두 개 이상의 값을 비교하여 하나라도 값이 일치(','와 '||'이 같은 기능을 한다)하면 해당 case의 문장이 실행된다
	case "monday", "tuesday":
		fmt.Println("월, 화요일은 수업 가는 날입니다.")
	case "wendsday", "tursday", "friday":
		fmt.Println("수, 목, 금요일은 실습 가는 날입니다/")
	}
}
