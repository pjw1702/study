package main

import "fmt"

// Type Generic
type Node[T any] struct {
	next *Node[T] // 다음 노드를 가리키는 포인터
	prev *Node[T] // 이전 노드를 가리키는 포인터
	val  T
}

func main() {
	root := &Node[int]{nil, nil, 10}
	// root.next.prev = root
	root.next = &Node[int]{nil, root, 20}
	root.next.next = &Node[int]{nil, root.next, 30}

	tail := root.next.next

	// n.next(다음 노드)가 nil일 때까지(데이터가 존재하는 모든 노드) 반복
	for n := root; n != nil; n = n.next {
		fmt.Printf("node val: %d\n", n.val)
	}

	fmt.Println("")

	// n.prev(이전 노드)가 nil일 때까지(데이터가 존재하는 모든 노드를 반대로) 반복
	for n := tail; n != nil; n = n.prev {
		fmt.Printf("node val: %d\n", n.val)
	}
}

/* 출력 값
node val: 10
node val: 20
node val: 30

node val: 30
node val: 20
node val: 10
*/
