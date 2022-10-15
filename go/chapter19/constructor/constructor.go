package main

import (
	"fmt"
	"time"
)

func main() {
	// 생성자를 통해 클래스나 메소드를 생성하도록 규정한 문법은 없다
	// 어떠한 방법으로도 클래스나 메소드를 생성할 수 있다
	var t1 = time.Timer{}
	var t2 = time.NewTimer(time.Second)
	fmt.Println(t1)
	fmt.Println(t2)
}
