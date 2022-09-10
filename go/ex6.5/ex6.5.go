/*실수 오차*/
package main

import "fmt"

func main() {
	var a float64 = 0.1
	var b float64 = 0.2
	var c float64 = 0.3

	fmt.Printf("%f + %f == %f: %v\n", a, b, c, a+b == c)
	fmt.Println(a + b) // 컴퓨터의 트랜지스터가 실수 표현식에서 소수부를 정확하게 표현을 하지 못함 그래서 0.3에 최대한 가까운 숫자를 표현 함
}
