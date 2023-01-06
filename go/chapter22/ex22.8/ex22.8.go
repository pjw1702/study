package main

import "fmt"

const M = 10

// 해쉬의 원리 기능을 하는 함수
func hash(d int) int {
	return d % M
}

func main() {
	// 배열 선언
	m := [M]string{}

	// 배열 초기화
	m[hash(23)] = "송하나"
	m[hash(259)] = "백두산"

	// 나누는 값이 상수 M이므로 입력과 출력이 같고 Access 시간도 일정하다
	fmt.Printf("%d = %s\n", 23, m[hash(23)])
	fmt.Printf("%d = %s\n", 259, m[hash(259)])

	/*
		출력 결과는 아래와 같다
		23 = 송하나
		259 = 백두산
	*/
}
