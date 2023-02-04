// 일반 함수는 상태를 가질 수 없지만 함수 리터럴은 내부 상태를 가질 수 있다
// 캡쳐는 값 복사가 아닌 레퍼런스 복사
// 즉, 포인터가 복사된다고 볼 수 있다
// 함수 리터럴의 매개 변수는 Low-Level의 포인터 변수로서 존재하고 외부 변수의 주소 값을 복사하여 수행한다
package main

import "fmt"

func main() {
	i := 0

	f := func() {
		// 지역 변수나 매개 변수로 i를 선언하지 않아도 i를 사용할 수 있고 컴파일 에러도 발생하지 않는다
		// 외부의 변수를 캡쳐하여 내부에서 사용 가능하다 (내부 상태를 가질 수 있다)
		i += 10
	}

	i++

	f()

	fmt.Println(i)
}