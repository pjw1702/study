package main

import "fmt"

func addNum(slice []int) {
	slice = append(slice, 4)
}

func main() {
	slice := []int{1, 2, 3}
	addNum(slice)

	// addNum() 함수에 의한 append() 호출에도 불구하고 여전히 1, 2, 3을 호출한다 (값이 바뀌지 않는다)
	// main()의 silce를 인자로서 넘겨받을 때까지는 main()의 silce와 addNum()의 slice는 서로 같은 배열을 포인터로서 가리키게 된다
	// 하지만 append() 함수 호출로 인해 addNum()의 slice가 가리켜야 할 배열의 len과 cap의 길이가 늘어나므로 가리키던 배열을 복사하여 새로운 배열을 할당받아 새로운 공간의 배열을 가리키게 된다
	// 결과적으로 addNum()의 slice 요소 값만 변경되고 main()의 silce는 여전히 같은 배열을 가리키고 있으므로 변함이 없다
	fmt.Println(slice) // 출력 값: [1 2 3]
}
