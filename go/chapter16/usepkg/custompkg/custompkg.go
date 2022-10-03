package custompkg

import (
	"fmt"
	"go/chapter16/usepkg/exinit" // main에서도 import를 하고 custompkg에서도 import하여 총 두번 import 하지만 초기화는 최초 한 번만 진행하여 출력값은 직접 호출한 함수 PrintD를 제외하곤 한 번만 출력된다
)

// 함수 뿐만 아니라 구조체, 변수 등도 모두 첫 알파벳이 대문자로 시작하여야 외부로 공개할 수 있고, 소문자이면 공개하지 않는다
type Student struct {
	Name  string
	Age   int
	score int // 외부로 공개하지 않으므로 import 시킨 프로그램 내에서 사용할 수 없다
}

var Ratio int
var ttt int // 외부로 공개하지 않으므로 import 시킨 프로그램 내에서 사용할 수 없다

const PI = 3.14
const pI2 = 3.1415 // 외부로 공개하지 않으므로 import 시킨 프로그램 내에서 사용할 수 없다

func PrintCustom() {
	fmt.Println("This is custom package!")
	exinit.PrintD()
}

// 다음과 같은 함수가 존재하는 경우, 함수가 외부로 공개되지 않아서 main 프로그램에서 다음 함수를 사용하고 빌드 시 컴파일 에러가 발생한다
// 함수의 첫 알파벳이 대문자이면 외부로 공개하고, 소문자이면 공개하지 않는다
// 알파벳의 대 소문자를 구분할 수 없으므로 함수명을 한글로 지정할 수 없다
/**
func printcustom2() {
	fmt.Println("This is custom package222!")
}
*/
