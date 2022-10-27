// 인터페이스의 기본 값은 nil이다
// 다음과 같은 경우는, Attacker 인터페이스가 가리키고 있는 타입의 메서드가 정의되어 있지 않으므로 nil 값을 가진다
// 문법적으로는 이상이 없으므로 컴파일은 무리 없이 진행되지만, 실행 과정에서 런타임 에러가 발생한다
// panic: runtime error: invalid memory address or nil pointer dereference 와 같이 에러가 발생한다
package main

type Attacker interface {
	Attack()
}

func main() {
	var att Attacker
	att.Attack()
}
