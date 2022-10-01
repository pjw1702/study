/**
문자열 합산을 이용하여 문자열을 변경하는 경우, 문자열이 매우 길어지게 되면 글자 몇 개만 늘어나고 변경되는 경우에도 문자열 전체를 복사하는 과정이 이루어져 비효율적이다
builders 메소드는 내부적으로 미리 할당된 슬라이스를 가지고 있다
builders 메소드를 사용하면 공간의 재 할당할 필요없이 미리 할당된 공간에 값을 저장하므로 더욱 효율적이다
*/

package main

import (
	"fmt"
	"strings" // builders 메소드를 사용하기 위한 패키지
)

// 문자열 합산을 이용하여 소문자를 대문자로 바꾸는 함수
func ToUpper1(str string) string {
	var rst string
	for _, c := range str {
		// 소문자일 경우
		if c >= 'a' && c <= 'z' {
			rst += string('A' + (c - 'a')) // 해당 문자에 대응하는 대문자를 출력(모든 알파벳은 A 혹은 a에서 시작하므로 해당 문자에서 'a'를 빼면 A에서 해당 문자까지의 이동 길이를 계산할 수 있다)
		} else {
			rst += string(c)
		}
	}

	return rst
}

// strings 패키지의 builder 메소드를 이용하여 소문자를 대문자로 바꾸는 함수
func ToUpper2(str string) string {
	var builder strings.Builder
	for _, c := range str {
		// 소문자일 경우
		if c >= 'a' && c <= 'z' {
			builder.WriteRune('A' + (c - 'a')) // 해당 문자에 대응하는 대문자를 출력(모든 알파벳은 A 혹은 a에서 시작하므로 해당 문자에서 'a'를 빼면 A에서 해당 문자까지의 이동 길이를 계산할 수 있다)
		} else {
			builder.WriteRune(c)
		}
	}
	return builder.String()
}

// 결과 값은 같지만 내부적인 동작 구조는 완전히 다르다

func main() {
	var str string = "Hello World"

	fmt.Println(ToUpper1(str))
	fmt.Println(ToUpper2(str))
}
