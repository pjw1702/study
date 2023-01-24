package main

import (
	"bufio"
	"errors"
	"fmt"
	"strconv"
	"strings"
)

func MultipleFromString(str string) (int, error) {
	// bufio.NewScanner() = 스트림을 통해 어떠한 데이터를 한 단어씩(또는 한 줄씩 등, 특정 규칙으로) 가져올 수 있음
	// 인자로 io Reader를 받음 (strings.Newreader() = 특정 문자열을 읽어오는 Reader 객체를 만들어서 반환)
	scanner := bufio.NewScanner(strings.NewReader(str))
	scanner.Split(bufio.ScanWords)

	pos := 0
	a, n, err := readNextInt(scanner)
	if err != nil {
		// fmt.Errorf()가 'err'변수를 통해 얻은 에러를 감싸서 새로운 에러 객체를 생성
		return 0, fmt.Errorf("failed to readNext(), pos:%d err:%w", pos, err)
	}

	pos += n + 1
	b, _, err := readNextInt(scanner)
	if err != nil {
		// fmt.Errorf()가 'err'변수를 통해 얻은 에러를 감싸서 새로운 에러 객체를 생성
		return 0, fmt.Errorf("failed to readNext(), pos:%d err:%w", pos, err)
	}
	return a * b, nil

}

// 다음 단어를 읽어서 숫자로 반환
// 변환된 숫자, 읽은 글자 수, 에러를 반환
func readNextInt(scanner *bufio.Scanner) (int, int, error) {
	// 한 단어를 읽었는데 실패할 경우
	if !scanner.Scan() {
		return 0, 0, fmt.Errorf("failed to scan")
	}
	word := scanner.Text()
	// strconv.Atoi(): 문자열로 된 숫자를 정수형 숫자로 바꿈
	number, err := strconv.Atoi(word) // "24" -> 24로 바꿈, "abc" -> error 반환
	if err != nil {
		// fmt.Errorf()가 'err'변수를 통해 얻은 에러를 감싸서 새로운 에러 객체를 생성
		return 0, 0, fmt.Errorf("failed to convert word to int, word:%s, err:%w", word, err)
	}
	return number, len(word), nil
}

func readEq(eq string) {
	// readNextInt()와 MultipleFromString() 함수를 통해, 'err' 변수는 총 두 번 감싸진 에러를 반환받는다
	rst, err := MultipleFromString(eq)
	if err == nil {
		fmt.Println(rst)
	} else {
		fmt.Println(err)
		// 에러가 무엇인지 확인
		// 감싼 에러를 끄집어 내는 과정
		var numError *strconv.NumError
		// 두 번 감싸진 에러를 strconv.NumError 객체로 변환
		// strconv.NumError 객체로 변환이 가능하다는 소리는 convert 할 때 실패한 정보, 즉 strconv.Atoi(word)에서 에러가 발생한 경우 꺼내와서 그 객체를 출력
		if errors.As(err, &numError) {
			fmt.Println("NumberError", numError)
			// fmt.Println("NumberError", numError.Func)
		}
	}
}

func main() {
	readEq("123 3")
	readEq("123 abc")
}

/* 출력 결과
369
Failed to readNext(), pos:4 err:Failed to convert word to int, word:abc, err:strconv.Atoi: parsing "abc": invalid syntax
NumberError strconv.Atoi: parsing "abc": invalid syntax
*/

/* 호출 순서
* main -> ReadEq -> MultipleFromString -> readNextInt -> strconv.Atoi -> Error
	* 에러는 가장 중에 호출된 프로그램으로 부터 발생한다
	* 에러를 처리 못할 경우, 바로 상위의 호출 자에게 전달한다
	* 호출자에게 어느 부분에서 에러가 발생했는지를 명확하게 알려줄 필요가 있다
		* 어느 부분에서 에러가 발생했는지를 알리기 위해 각 호출자 마다의 정보를 래핑(wrapping)하여 전달한다
*/
