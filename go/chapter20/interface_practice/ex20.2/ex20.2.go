package main

import (
	"go/chatper20/interface_practice/fedex"
	korePost "go/chatper20/interface_practice/koreaPost"
)

// 인터페이스 없이 모듈과 메서드만 있는 경우, 아래와 같이 패키지를 바꾸면 소스코드 전체를 바꿔야 한다

// import "go/chatper20/interface_practice/fedex"
// func SendBook(name string, sender *fedex.FedexSender) {
// 	sender.Send(name)
// }

// func main() {
//	sender := &fedex.FedexSender{}
//	SendBook("어린 왕자", sender)
//	SendBook("그리스인 조르바", sender)
// }

// import "go/chatper20/interface_practice/korePost"
// func SendBook(name string, sender *korePost.PostSender) {
//	sender.Send(name)
// }

// func main() {
//	sender := &korePost.PostSender{}
//	SendBook("어린 왕자", sender)
//	SendBook("그리스인 조르바", sender)
// }

// main() 함수에서 메서드를 초기화하는 과정에서는 패키지 구분이 필요하지만, 그 외에의 함수나 메서드를 사용할 때에는 변경 작업이 필요하지 않다
// 덕 타이핑: Java에서의 implements 기능, 즉 Sender라는 인터페이스가 어떤 패키지의 타입과 연관되어 있는지 상관 없이, 그냥 Sender를 인터페이스로 선언하면 패키지 상관없이 모든 패키지에서의 주소 값을 전달받은 패키지의 타입과 연관짓겠다는 의미
type Sender interface {
	Send(percel string)
}

func SendBook(name string, sender Sender) {
	sender.Send(name)
}

func main() {
	sender := &korePost.PostSender{}
	SendBook("어린 왕자", sender)
	SendBook("그리스인 조르바", sender)
	sender2 := &fedex.FedexSender{}
	SendBook("어린 왕자", sender2)
	SendBook("그리스인 조르바", sender2)
}
