package main

import (
	"container/list"
	"fmt"
)

func main() {
	v := list.New()       // 리스트 초기화
	e4 := v.PushBack(4)   // 맨 뒤에 넣는 메서드 (***4), e4는 4를 가리킴
	e1 := v.PushFront(1)  // 멘 앞에 넣는 메서드 (1**4), e1는 1을 가리킴
	v.InsertBefore(3, e4) // 3이라는 값을 e4 전에 넣는 메서드 (1*3 4)
	v.InsertAfter(2, e1)  // 2라는 값을 e1 후에 넣는 메서드 (1 2 3 4)

	for e := v.Front(); e != nil; e = e.Next() { // e가 맨 처음엔 v.Front()값인 맨 앞의 1을 가리킴, e.Next()는 한 칸씩 다음 노드로 이동하여 가리키는 메서드
		fmt.Print(e.Value, "")
	}

	fmt.Println()
	for e := v.Back(); e != nil; e = e.Prev() { // e가 맨 처음엔 v.Back()값인 맨 뒤의 4을 가리킴, e.Prev()는 한 칸씩 이전 노드로 이동하여 가리키는 메서드
		fmt.Print(e.Value, "")
	}
}
