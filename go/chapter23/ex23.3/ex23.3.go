package main

import (
	"fmt"
)

type PasswordError struct {
	Len        int
	RequireLen int
}

// error는 string 타입을 반환하는 Error()메소드를 가지는 인터페이스이다
// Error() 메소드에 string 타입을 반환: 에러처리 메소드로 사용될 수 있다(error 인터페이스 사용)
func (err PasswordError) Error() string {
	return "암호 길이가 짧습니다"
}

// 'PasswordError'구조체가 Error()메소드를 기본으로 가지고 있으므로(매개변수, return 없음) return 타입이 error 함수에서 반환 값으로 사용 가능하다
func RegisterAccount(name, password string) error {
	if len(password) < 8 {
		// 아래 세 가지 return은 모두 같은 값을 반환한다
		// return fmt.Errorf("...")
		// return errors.New("...")
		return PasswordError{len(password), 8}
	}
	// ...
	return nil
}

func main() {
	err := RegisterAccount("myID", "myPw")
	if err != nil {
		if errINfo, ok := err.(PasswordError); ok {
			fmt.Printf("%v Len:%d RequireLen:%d\n",
				errINfo, errINfo.Len, errINfo.RequireLen)
		} else {
			fmt.Println("회원 가입되었습니다.")
		}
	}
}
