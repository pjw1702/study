package main

import "fmt"

func main() {
	a := 3 // int=int64
	var b float64 = 3.5

	var c int = int(b)  // c의 타입은 int 이지만 b의 타입은 float64이므로 대입할 수 없음, 실수형을 정수형으로 형 변환시 소숫점이 버림됨(3.5 -> 3)
	d := float64(a) * b // 타입이 다른 두 변수 a, b는 연산이 불가 함

	var e int64 = 7 // 타입 자체가 달라서(int, int64는 서로 다른 타입으로 인식 함) 같은 사이즈라도 대입이 안됨
	f := a * int(e)

	fmt.Println(a, b, c, d, e, f)
}
