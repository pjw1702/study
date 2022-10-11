// append() 동작 원리
/*
빈 공간이 충분한가?
YES: 빈 공간에 요소 추가
NO: 새로운 배열 할당 -> 복사 -> 요소 추가

남은 빈 공간 = cap - len

case1. 공간이 남을 경우
slice1이 가리키는 같은 배열을 가리키게 되고 {1, 2, 3, 4, 5}로 저장된다

case2. 공간이 남지 않을 경우
slice1의 배열을 복사하여 새로운 길이의 배열을 생성하여 slice1과 다른 곳을 가리키는 배열에 값을 저장한다

만약 slice1과 slice2가 같은 배열을 가리키는 경우, slice1[1] = 100을 수행한 경우 printf("%d", slice2[1])은 어떻게 될까?
	-> 같은 곳을 가리키므로 slice2[1]또한 100이 된다

슬라이스의 빈 공간 유무에 따라 append()를 통해 생성한 슬라이스의 요소 값이 변하므로 주의해야 한다
*/

package main

import "fmt"

func main() {
	var slice = []int{1, 2, 3} // 요소가 3개인 슬라이스

	// append(): 요소를 추가한 슬라이스 생성: 첫 번째 인자에 추가할 슬라이스, 두 번째 인자에 추가할 요소 값 (생성된 슬라이스 반환)
	// 첫 번째 인자로 들어갈 슬라이스는 바뀌지 않고 첫 번째 인자의 슬라이스에 두 번째 인자를 다음 요소 값으로서 추가
	slice2 := append(slice, 4) // 요소 추가 {1, 2, 3, 4}

	fmt.Println(slice)
	fmt.Println(slice2)
}
