// 1. 0~99 사이 랜덤한 숫자 생성
// 2. 사용자 입력
// 3. 숫자 비교: 입력값이 크면 "큽니다" 작으면 "작습니다" 출력 (두 번 반복)
// 4. 입력값이 같으면 프로그램 종료

/** 랜덤
"math/rand" 패키지 사용
rand.Seed(seed int64) 함수로 랜덤 시드(Seed) 설정: 랜덤 생성 값을 입력 값(시드)에 따라 매번 다르게 출력하는 함수
rand.Intn(n int) int 함수로 0~(n-1) 사이의 값 생성(뽑음)
*/

// 컴퓨터는 랜덤과 그다지 친한 사이가 아니다: 컴퓨터는 계산이 탁월한 기계라 출력이 일정한 프로그램이 효율적인데 랜덤은 출력이 매번 바뀌므로 컴퓨터 입장에서 정확하지 않다
// 정확히는 실제 랜덤을 수행하는 것이 아닌 유사 랜덤을 수행하는 것이다

/** 랜덤 시드 선택
프로그램 실행 시마다 계속 변화하는 값
"time" 패키지는 시간과 관련된 기능 제공
time.Now() - Time 현재 시각 제공
func (t Time) UnixNano() int64 - Time 객체 t의 1970년 1월 1일 부터 현재까지 경과된 시간을 nano seconds 단위로 반환
*/

/** time 패키지
시각을 나타내는 Time 객체
시간을 나타내는 Duration 객체
타임존을 나타내는 Location
*/

/*
package main

import (
	"fmt"
	"math/rand"
	"time"
)
*/

// 랜덤 패키지 사용 예제
/*
func main() {
	t := time.Now()
	fmt.Println(t)
	rand.Seed(t.UnixNano())

	for i := 0; i < 10; i++ {
		fmt.Println(rand.Intn(100))
	}
}
*/

// 타임 패키지 사용 예제
/*
func main() {
	loc, _ := time.LoadLocation("Asia/Seoul")                               // 타임 존 생성
	const LongForm = "Jan 2, 2006 at 3:04pm"                                // 타임을 읽는 포맷 설정
	t1, _ := time.ParseInLocation(LongForm, "Jun 10, 2021 at 10:00pm", loc) // 특정 문자열을 통해서 타임을 읽어오는 함수: LongForm 포맷을 기준으로 시간 생성
	fmt.Println(t1, t1.Location(), t1.UTC())

	const shortForm = "2006-Jan-02"
	t2, _ := time.Parse(shortForm, "2021-Jun-14") // 타임 존 없이 시간을 읽어 옴: 그리니치 천문대의 시간을 기준으로 시간을 읽음
	fmt.Println(t2, t2.Location())

	t3, err := time.Parse("2021-06-01 15:20:21", "2021-06-14 20:04:05") // Go에서 설정된 기본 포맷과 맞지 않으므로 에러 발생
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(t3, t3.Location())

	d := t2.Sub(t1) // t2 - t1의 duration(시간)
	fmt.Println(d)
}
*/

package main

import "fmt"

func main() {
	fmt.Println("랜덤 패키지와 타임 패키지를 사용하는 예제")
}
