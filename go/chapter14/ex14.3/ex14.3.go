package main

import "fmt"

type Data struct {
	value int      // 8바이트
	data  [200]int // 200*8 바이트
}

/*
// 다음과 같은 경우 출력 값은 0,0이 출력된다
func ChangeData(arg Data) {
	arg.value = 999
	arg.data[100] = 999
}

func main() {
	var data Data

	ChangeData(data)
	fmt.Printf("value = %d\n", data.value)
	fmt.Printf("data[100] = %d\n", data.data[100])

}
*/

// 다음과 같은 경우 출력 값은 999,999가 출력된다
func ChangeData(arg *Data) {

	arg.value = 999
	arg.data[100] = 999

	/** 다음과 같은 표현이 더 정확한 표현이다 */
	// (*arg).value = 999		// c/c++에서는 arg -> value = 999와 같이 표현한다
	// (*arg).data[100] = 999
}

func main() {
	var data Data

	ChangeData(&data)
	fmt.Printf("value = %d\n", data.value)
	fmt.Printf("data[100] = %d\n", data.data[100])

}
