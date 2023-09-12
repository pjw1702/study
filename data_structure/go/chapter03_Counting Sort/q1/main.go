package main

import "fmt"

// Q1: 0 ~ 10까지의 수를 정렬하시오.
func main() {
	arr := []int{5, 1, 3, 2, 5, 2, 6, 8, 2, 0, 4, 5, 1, 6, 8, 2, 7, 9, 2, 1, 5, 6, 10}

	var count [11]int // 값의 범위가 10까지 이므로 인덱스는 0 ~ 10이다
	for i := 0; i < len(arr); i++ {
		count[arr[i]]++
	}

	fmt.Println("count: ", count)

	// case 1: 이중for문 사용
	// sorted := make([]int, 0, len(arr))
	// for i := 0; i < 11; i++ {
	// 	for j := 0; j < count[i]; j++ { // 값의 중복에 대한 정렬을 위함: ex) 5가 두 번 나오면 5 두 개를 나란히 정렬
	// 		sorted = append(sorted, i)
	// 	}
	// }

	// case 2: 단일 for문 사용
	// 중복되는 숫자가 많을 경우 용이
	// case 1보다 성능 면에선 더 효율적이지만 가독성이 떨어짐(이해하기 여려운 코드)
	for i := 1; i < 11; i++ {
		count[i] += count[i-1]
	}

	fmt.Println("count2: ", count)

	sorted := make([]int, len(arr))
	for i := 0; i < len(arr); i++ {
		sorted[count[arr[i]]-1] = arr[i] // 정렬된 배열의 해당 인덱스에 데이터 저장
		count[arr[i]]--                  // 해당 숫자에 대한 카운트 수는 한 번 삽입될 때마다 하나씩 줄임
	}

	fmt.Println(sorted)
}

// 실행 결과
// count:  [1 3 5 1 1 4 3 1 2 1 1]
// count2:  [1 4 9 10 11 15 18 19 21 22 23]	// case2는 카운팅 수가 데이터 갯수가 아닌 인덱스 번호를 기준으로 카운팅한다: ex) count[1]은 4로 출력이 되었는데, 이것은 sorted[4]부터 sorted[8]까지 모두 2로 채워져야 한다는 것을 의미한다
