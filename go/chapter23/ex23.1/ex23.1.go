package main

import (
	"bufio"
	"fmt"
	"os"
)

func ReadFile(filename string) (string, error) {
	// 매개변수 'filname'을 통해 입력 받은 파일을 os.Open() 함수를 통해 연다
	file, err := os.Open(filename)
	if err != nil {
		return "", err
	}

	// 파일은 한 번 열리면 파일 핸들을 닫아줘야 한다
	defer file.Close()
	// Reader 인스턴스를 통해 파일의 데이터를 가져올 수 있다
	// 인자로 인터페이스를 받는다
	rd := bufio.NewReader(file)
	// string 타입의 데이터를 가져온다
	line, _ := rd.ReadString('\n') // 반환 값으로 string 타입의 파일 데이터와 error(io.EOF)를 반환하는데, 크게 중요하지 않으므로 _ 처리하였다
	return line, nil
}

func WriteFile(filename string, line string) error {
	// filename에 해당하는 파일 생성
	// 에러: 디스크의 용량이 없거나 파일을 생성할 권한이 없는 경우 등에 발생한다
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	// 표준 출력의 스트림을 화면이 아닌 다른 곳으로 변경할 수 있다
	// 매개 변수로 파일이 들어가면 파일으로 출력된다
	_, err = fmt.Fprintln(file, line)
	return err
}

const filename string = "data.txt"

func main() {
	// 파일을 읽음
	line, err := ReadFile(filename)
	// 파일을 읽는데 실패할 경우 파일을 새로 생성
	if err != nil {
		err = WriteFile(filename, "This is WriteFile")
		if err != nil {
			fmt.Println("파일 생성에 실패했습니다.", err)
			return // 프로그램 종료
		}
	}
	// 파일을 처음 읽었을 때 에러가 발생하지 않으면 파일을 다시 읽기
	// 파일을 새로 생성했을 경우 생성한 파일의 데이터를 확인할 수 있다
	line, err = ReadFile(filename)
	if err != nil {
		fmt.Println("파일 읽기에 실패했습니다", err)
		return
	}

	fmt.Println("파일 내용:", line)

}
