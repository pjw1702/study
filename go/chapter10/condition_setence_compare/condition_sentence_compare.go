package main

import "fmt"

func main() {
	temp := 18

	switch true {
	// 값에 조건 연산자를 대입하여 조건문 처럼 사용할 수 있다. 하지만 이런 경우는 대부분 if문을 사용한다
	case temp < 10, temp > 30:
		fmt.Println("바깥 활동하기 좋은 날씨가 아닙니다.")
	case temp >= 10 && temp < 20:
		fmt.Println("약간 추울 수 있으니 가벼운 겉옷을 준비하세요.")
	case temp >= 15 && temp < 25:
		fmt.Println("야외 활동하기 좋은 날씨입니다.")
	default:
		fmt.Println("따뜻합니다.")
	}
}
