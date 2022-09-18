package main

import "fmt"

type ColorType int // type: 별칭(alias) 타입(실제 타입: int)
const (
	Red ColorType = iota // 100 + iota를 해도 상관 없다. 값의 저장이나 타입 자체 보다는 타입과 변수와의 상관 관계가 가지는 코드 상의 의미를 이해하는 것이 핵심 포인트이다
	Blue
	Green
	Yellow
)

func colorToString(color ColorType) string {
	switch color {
	case Red:
		return "Red"
	case Blue:
		return "Blue"
	case Green:
		return "Green"
	case Yellow:
		return "Yellow"
	default:
		return "Undefined"

	}
}

func getMyFavoriteColor() ColorType {
	return Red
}

func main() {
	// 함수 호출 인자에 또 다른 함수 호출의 return 값을 입력
	fmt.Println("My favorite color is", colorToString(getMyFavoriteColor()))

	// 다음과 같은 경우 Blue를 반환한다
	fmt.Println("My favorite color is", colorToString(1))
}
