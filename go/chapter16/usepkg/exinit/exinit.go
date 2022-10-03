package exinit

import (
	"fmt"
	// "go/chapter16/usepkg/custompkg"		// exinit과 custompkg가 서로 import 하려는 경우 import cycle이 발생하여 컴파일 에러가 발생한다
)

var (
	a = c + b // 4. a = 4 + 5 = 9
	b = f()   // 3. b = 5
	c = f()   // 2. a = c + b 초기화로 인해 c 연산이 먼저 진행된다 (c = 4)
	d = 3     // 1. d = 3
)

func init() { // 4. 변수가 모두 초기화되면 init()함수를 실행한다
	d++ // d = 6
	fmt.Println("exinit.init function", d)
}

func f() int {
	d++ // d = 4, d = 5
	fmt.Println("f() d:", d)
	return d
}

func PrintD() {
	fmt.Print("d:", d)
}
