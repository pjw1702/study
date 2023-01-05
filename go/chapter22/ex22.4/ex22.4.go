package main

import (
	"container/ring"
	"fmt"
)

func main() {
	r := ring.New(5)

	n := r.Len()

	for i := 0; i < n; i++ {
		r.Value = 'A' + i // r.Value: 가리키는 노드의 요소 값
		r = r.Next()
	}

	for j := 0; j < n; j++ {
		fmt.Printf("%c", r.Value)
		r = r.Next()
	}

	fmt.Println()
	for j := 0; j < n; j++ {
		fmt.Printf("%c", r.Value)
		r = r.Prev()
	}
}
