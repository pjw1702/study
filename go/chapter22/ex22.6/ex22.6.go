package main

import "fmt"

type Product struct {
	Name  string
	Price int
}

func main() {
	m := make(map[int]Product)

	m[16] = Product{"볼펜", 500}
	m[46] = Product{"지우개", 200}
	m[78] = Product{"자", 1000}
	m[345] = Product{"샤프", 3000}
	m[897] = Product{"샤프심", 500}

	for k, v := range m {
		// k: 키, v: 값, m: 맵
		// 배열: 인덱스, 값
		fmt.Println(k, v)
	}
}

/*
출력 결과는 아래와 같다
78 {자 1000}
345 {샤프 3000}
897 {샤프심 500}
16 {볼펜 500}
46 {지우개 200}

맵은 순회 출력 시 순서 보장이 안된다

맵에는 크게 HashMap과 SortedMap으로 나뉘는데, HashMap은 순서가 무작위로 출력되고 SortedMap은 키 값의 순서대로 출력된다
Go에서는 HashMap을 채택하고 있으므로 무작위로 출력된다
*/
