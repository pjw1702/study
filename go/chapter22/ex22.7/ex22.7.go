package main

import "fmt"

func main() {
	m := make(map[int]int)
	m[1] = 0
	m[2] = 2
	m[3] = 3

	delete(m, 3)
	delete(m, 4)
	//fmt.Println(m[3])	// 0 출력
	//fmt.Println(m[1])	// 0 출력
	// m[1]은 원래부터 값이 0으로 초기화 했으므로 둘 다 0으로 출력했어도 값의 존재 여부는 출력 값 만으로는 확인할 수 없다

	v, ok := m[3]
	fmt.Println(v, ok)
	v, ok = m[1]
	fmt.Println(v, ok)
	/*
		출력 값은 아래와 같다
		0 false
		0 true
	*/
}
