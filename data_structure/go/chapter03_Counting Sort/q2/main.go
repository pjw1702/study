package main

import "fmt"

// Q2: 알파벳 소문자로 이루어진 문자열 중 가장 많이 나오는 알파벳 캐릭터를 출력하시오
// a - z: 26문자
// 만약 문자열이 스트림으로 입력되는 경우에는 어떻게 할 것인가?: 실시간으로 입력이 되므로 Counting Sort로 해결할 수 없다
// 스트림인 경우엔 힙을 사용: 최대 최솟값에 대한 해결을 할 때 유용
func main() {
	str := "aksfhdklljsahflksahfdlskahfklsajflhakfdhjasflfakj"

	var count [26]int
	for i := 0; i < len(str); i++ {
		count[str[i]-'a']++
	}

	maxCount := 0
	var maxCh byte
	for i := 0; i < 26; i++ {
		if count[i] > maxCount { // 알파벳에 대해 카운팅 된 숫자 중 가장 큰 값 판별
			maxCount = count[i]
			maxCh = byte('a' + i)
		}
	}

	fmt.Printf("Max character: %c, count: %d", maxCh, maxCount)
}
