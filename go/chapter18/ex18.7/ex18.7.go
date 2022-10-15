/* 슬라이싱
슬라이싱은 배열의 일부를 (포인터로서)가리키는(집어내는) 기능

배열을 슬라이싱의 결과는 슬라이스이다

array[startIdx:endindex] 형태로 사용할 수 있다

전체 배열 중 슬라이싱을 시작할 메모리 주소를 포인터로서 가리키는 슬라이스이다

// array[n:m]
data = slice.data[n] ~ slice.data[m-1]
len	=	m-n
cap	=	cap-n
*/

package main

import "fmt"

func main() {
	array := [5]int{1, 2, 3, 4, 5}

	slice := array[1:2]
	fmt.Println("array:", array)
	fmt.Println("slice:", slice, len(slice), cap(slice))

	// 슬라이싱은 새로운 배열을 복사하는 개념이 아닌 기본 배열의 일부를 가리키는 기능이다
	array[1] = 100
	fmt.Println("After change second element")
	fmt.Println("array:", array)
	fmt.Println("slice:", slice, len(slice), cap(slice)) // 100으로 변경

	// slice는 원래 array[1]을 가리키는 길이가 len이 1인 슬라이스 였지만, append() 함수로 인해 500이 추가가 되어 len이 2로 늘어난다
	slice = append(slice, 500)
	fmt.Println("After append 500")
	fmt.Println("array:", array) // 500 추가
	fmt.Println("slice:", slice, len(slice), cap(slice))

}
