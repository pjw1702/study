package main

import (
	"fmt"
	"go/chapter16/usepkg/custompkg" // 초기화한 모듈의 경로와 패키지 명을 맞춰줘야 한다
	"go/chapter16/usepkg/exinit"    // 어떤 패키지를 import 하면 그 패키지는 초기화를 진행한다. 초기화를 진행하면 변수들이 초기화 되는데, 만약 init()함수가 존재 한다면 변수 초기화 완료된 다음 실행한다

	"github.com/guptarohit/asciigraph"
	"github.com/tuckersGo/musthaveGo/ch16/expkg"
)

func main() {
	custompkg.PrintCustom()

	s := custompkg.Student{}
	s.Name = "박정우"
	s.Age = 28
	fmt.Println(s.Name, s.Age)
	fmt.Println(custompkg.PI)
	custompkg.Ratio = 10
	fmt.Println(custompkg.Ratio)
	expkg.PrintSample()

	data := []float64{3, 4, 5, 6, 9, 7, 5, 8, 5, 10, 2, 7, 2, 5, 6}

	graph := asciigraph.Plot(data)
	fmt.Println(graph)

	//exinit
	custompkg.PrintCustom()
	exinit.PrintD()
}
