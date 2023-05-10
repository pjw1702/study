package main

import (
	"fmt"

	"github.com/pjw1702/study/web/go/web9/cipher"
	"github.com/pjw1702/study/web/go/web9/lzw"
)

type Component interface {
	Operator(string)
}

var sentData string
var recvData string

type SendComponent struct{}

func (self *SendComponent) Operator(data string) {
	// Send data
	sentData = data
}

type ZipComponent struct {
	com Component
}

func (self *ZipComponent) Operator(data string) {
	zipData, err := lzw.Write([]byte(data))
	if err != nil {
		// panic(err)
		fmt.Println("압축 실패!")
		return
	}
	self.com.Operator(string(zipData)) // 인터페이스를 통해, 본인이 가지고 있는 Operator 메서드를 압축한 데이터로 호출
}

type EncryptComponent struct {
	key string
	com Component
}

func (self *EncryptComponent) Operator(data string) {
	encryptData, err := cipher.Encrypt([]byte(data), self.key)
	if err != nil {
		// panic(err)
		fmt.Println("암호화 실패!")
		return
	}
	self.com.Operator(string(encryptData)) // 인터페이스를 통해, 본인이 가지고 있는 Operator 메서드를 암호화한 데이터로 호출
}

type DecryptComponent struct {
	key string
	com Component
}

func (self *DecryptComponent) Operator(data string) {
	decryptData, err := cipher.Decrypt([]byte(data), self.key)
	if err != nil {
		// panic(err)
		fmt.Println("복호화 실패!")
		return
	}
	self.com.Operator(string(decryptData))
}

type UnzipComponent struct {
	com Component
}

func (self *UnzipComponent) Operator(data string) {
	unzipData, err := lzw.Read([]byte(data))
	if err != nil {
		// panic(err)
		fmt.Println("압축해제 실패!")
		return
	}
	self.com.Operator(string(unzipData))
}

type ReadComponent struct{}

func (self *ReadComponent) Operator(data string) {
	recvData = data
}

func main() {
	// EncryptComponent가 ZipComponent를 가지며 ZipComponent가 SendComponent를 가진다
	// 암호화 -> 압축한 데이터를 SendComponent 객체를 통해 sentData에 저장한다
	// Decorator Pattern은 각 기능마다 포함되어 있거나 연결되어 있으므로, 부가적인 기능이다 제거할 기능은 해당 객체만 단계적으로 제거하거나 포함하여 추가하면 된다
	sender := &EncryptComponent{
		key: "abcde",
		com: &ZipComponent{
			com: &SendComponent{},
		},
	}

	sender.Operator("Hello World")

	fmt.Println(sentData)

	receiver := &UnzipComponent{
		com: &DecryptComponent{
			key: "abcde",
			com: &ReadComponent{},
		},
	}

	receiver.Operator(sentData)
	fmt.Println(recvData)
}
