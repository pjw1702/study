package main

import (
	"fmt"
)

func main() {
	var t [5]float64 = [5]float64{24.0, 25.9, 27.8, 26.9, 26.2}

	/** range는 여러 자료구조(slice, string, map 등)가 입력되었을 때 저장된 요소의 길이 만큼 순회하여 주는 기능이다*/
	for i, v := range t {
		fmt.Println(i, v)
	}

	/** '_'는 빈칸(blank) 지시자로, 반복문 실행에만 사용하고 변수를 사용하지 않을 경우 유용하다 */
	for _, v := range t {
		fmt.Println(v)
	}
}
