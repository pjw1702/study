package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"time"
)

func main() {
	// lottery.exe, filename, count
	if len(os.Args) < 3 { // os.Args: 실행 인자
		fmt.Fprintln(os.Stderr, "Invalid arguments!\nlottery filename count") // fmt.Fprintln: 어느 곳에 write 할 것인지 지정 가능, os.Stderr: 표준 입출력
		return
	}

	/*
		Go는 에러를 처리할 때 보통 함수나 메서드의 마지막 반환 값으로 에러 상태를 반환한다
		에러가 발생하지 않으면 에러 상태로 nil을 반환하고, 에러가 발생하면 에러 상황에 맞는 error 값을 반환한다
		에러를 반환하는 함수를 호출할 때는 반드시 에러 값을 확인해서 에러 상황에 알맞은 적절한 처리를 해야 한다
		Go는 return 값을 동시에 여러번 반환할 수 있으므로 그 return값을 대입할 변수도 return값의 갯수와 순서에 맞게 지정해야 한다
	*/

	filename := os.Args[1]                 // 첫 번째 출력인자: filename
	count, err := strconv.Atoi(os.Args[2]) //strconv.Atoi(): os.Args[]는 string타입이므로 string -> int 형으로 임시적으로 형 변환을 가능하게 한다
	if err != nil {                        // 두 번째 출력인자인 count의 출력 값이 err 변수로 저장이 되었을 때, err 변수의 값이 nill이 아닌 경우
		fmt.Fprintln(os.Stderr, "cannot convert count to integer! count:", count)
		return
	}

	// filename의 파일 내용을 읽어서 후보자 목록을 생성하는 매서드
	candidates, err := readCandidates(filename)
	if err != nil {
		fmt.Fprintln(os.Stderr, "cannot read candidates file!", err)
		return
	}

	rand.Seed(time.Now().UnixNano()) //rand: 난수를 발생시키는 함수, Seed 매서드: 난수 발생을 매번 바꾸는 매서드, time.Now().UnixNano: 현재 시각으로 설정

	winners := make([]string, count) // 뽑힌 사람 저장: 슬라이스 string을 뽑힌 사람의 수(count)만큼 저장
	for i := 0; i < count; i++ {
		n := rand.Intn(len(candidates)) //random한 Integer 값을 뽑는 함수(0~N 사이의 값을 뽑음)
		winners[i] = candidates[n]
		candidates = append(candidates[:n], candidates[n+1:]...) // 0~(n-1)과 (n+1)~끝까지를 제외한 사람(n번째 사람), 즉 n번째에서 뽑힌 사람만 후보에서 제외
	}

	fmt.Println("Winners are !!!")
	for _, winner := range winners {
		fmt.Println(winner)
	}

}

func readCandidates(filename string) ([]string, error) {
	file, err := os.Open(filename) // 파일을 읽는 함수(파일 핸들과 에러 return)
	if err != nil {
		return nil, err
	}
	defer file.Close() // 파일을 닫는 함수(바로 실행되는 함수가 아니라 readCandidates함수가 종료되기 직전에 실행되어서 파일이 닫힌다)

	scanner := bufio.NewScanner(file) // 파일을 한 줄씩 읽기위함(파일 핸들러를 매개변수로 전달)
	var candidates []string
	for scanner.Scan() { // 한 줄씩 읽어오는 함수(못 읽어오는 경우, 즉 모두 읽은 경우 false를 반환하여 함수가 자동으로 종료 됨)
		candidates = append(candidates, scanner.Text()) // candidates변수에 scanner를 통해 읽어온 값(scanner.Text())을 추가
	}
	return candidates, nil
}
