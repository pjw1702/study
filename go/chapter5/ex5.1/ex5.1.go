package main

import "fmt" // go의 표준 입출력 함수를 제공하는 패키지

func main() {
	var a int = 10
	var b int = 20
	var f float64 = 32799438743.8297

	fmt.Print("a: ", a, "b: ", b)
	fmt.Println("a: ", a, "b: ", b, "f: ", f)
	fmt.Printf("a: %d, b: %d, f: %f\n", a, b, f) // %d, %f: 서식 문자(formatter)

}
