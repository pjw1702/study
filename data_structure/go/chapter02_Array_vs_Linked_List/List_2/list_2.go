package main

import "fmt"

type Node[T any] struct {
	next *Node[T]
	val  T
}

func Append[T any](node *Node[T], next *Node[T]) *Node[T] {
	node.next = next
	return next
}

func main() {
	// 노드 초기화
	root := &Node[int]{nil, 10}
	node := root

	// 삽입
	for i := 0; i < 9; i++ {
		node = Append(node, &Node[int]{nil, (i + 2) * 10})
	}

	for n := root; n != nil; n = n.next {
		fmt.Println("val: ", n.val)
	}

	fmt.Println()
	fmt.Println()

	// 데이터 20 뒤에 25 삽입
	// next 포인트만 살짝 바꿔주면 된다
	//node = root.next          // node.val = 20
	//originalNext := node.next // node.val = 30
	//node = Append(node, &Node[int]{nil, 25})
	//node.next = originalNext

	//for n := root; n != nil; n = n.next {
	//	fmt.Println("val: ", n.val)
	//}

	// 일곱번 째 데이터 검색
	// array는 a[6]과 같이 바로 요소를 찾을 수 있지만 리스트는 연결된 노드 수 대로 따라서 찾아야 한다
	node = root
	for i := 0; i < 6; i++ {
		node = node.next
	}
	fmt.Println("7th val: ", node.val)

}

/* 삽입 출력 결과
val:  10
val:  20
val:  30
val:  40
val:  50
val:  60
val:  70
val:  80
val:  90
val:  100


val:  10
val:  20
val:  25
val:  30
val:  40
val:  50
val:  60
val:  70
val:  80
val:  90
val:  100
*/

/* 검색 출력 결과
val:  10
val:  20
val:  30
val:  40
val:  50
val:  60
val:  70
val:  80
val:  90
val:  100


7th val:  70
*/
