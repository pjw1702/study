package singlelinkedlist

// any type
type Node[T any] struct {
	next  *Node[T]
	Value T
}

type LinkedList[T any] struct {
	root *Node[T]
	tail *Node[T]

	count int // 길이 구하는 두 번째 방법
}

// 뒷 방향으로 데이터 삽입(->)
func (list *LinkedList[T]) PushBack(value T) {
	node := &Node[T]{
		Value: value,
	}

	list.count++

	if list.root == nil {
		list.root = node
		list.tail = node
		return
	}

	// 순 방향이므로 기존의 tail이 가리키는 노드가 새로 추가할 노드를 가리키게 해야 함
	list.tail.next = node
	list.tail = node
}

// 앞 방향으로 데이터 삽입(<-)
func (list *LinkedList[T]) PushFront(value T) {
	node := &Node[T]{
		Value: value,
	}

	list.count++

	if list.root == nil {
		list.root = node
		list.tail = node
		return
	}

	// 역 방향이므로 새로 추가할 노드의 next가 기존의 root가 가리키는 노드를 가리키게 해야 함
	node.next = list.root
	list.root = node
}

// 첫 번째 데이터 노드 반환
func (list *LinkedList[T]) Front() *Node[T] {
	return list.root
}

// 맨 마지막 노드 반환
func (list *LinkedList[T]) Back() *Node[T] {
	return list.tail
}

// 길이(노드 갯수) 반환
func (list *LinkedList[T]) Count() int {
	node := list.root
	cnt := 0

	for ; node != nil; node = node.next {
		cnt++
	}
	return cnt
}

func (list *LinkedList[T]) Count2() int {
	return list.count
}

func (list *LinkedList[T]) GetAfter(index int) *Node[T] {
	if index >= list.Count2() {
		return nil
	}

	i := 0
	for node := list.root; node != nil; node = node.next {
		if i == index {
			return node
		}
		i++
	}
	return nil
}

func (list *LinkedList[T]) InsertAfter(node *Node[T], value T) {
	// 직전 노드에 리스트에 없는 노드를 입력하고 InsertAfter 실행 시, Count()메소드는 정상 실행되지만 Count2()메소드로는 fail이 발생한다
	// node와 newNode는 Node[T] 타입의 구조체로, LinkedList[T]타입의 구조체와 다른 구조체인데 list.count가 증가가되므로 발생하는 문제이다
	if !list.isIncluded(node) {
		return
	}
	newNode := &Node[T]{
		Value: value,
	}

	// Case 1
	// originNext := node.next		// 새 노드의 다음 노드를 백업
	// node.next = newNode
	// newNode.next = originNext

	// Case 2
	node.next, newNode.next = newNode, node.next
	list.count++
}

// 데이터 노드 추가시, 직전 노드가 리스트에 있는 노드인지 없는 노드인지를 판별
func (list *LinkedList[T]) isIncluded(node *Node[T]) bool {
	inner := list.root
	for ; inner != nil; inner = inner.next {
		if inner == node {
			return true
		}
	}
	return false
}

func (list *LinkedList[T]) InsertBefore(node *Node[T], value T) {
	if node == list.root {
		list.PushFront(value)
		return
	}
	prevNode := list.findPrevNode(node)
	if prevNode == nil {
		return
	}
	newNode := &Node[T]{
		Value: value,
	}
	prevNode.next, newNode.next = newNode, node
	list.count++
}

func (list *LinkedList[T]) findPrevNode(node *Node[T]) *Node[T] {
	inner := list.root
	for ; inner != nil; inner = inner.next {
		if inner.next == node {
			return inner
		}
	}
	return nil
}

// 맨 첫번째 노드부터 제거
// 1, 2, 3이 삽입되었을 시 PopFront()메소드를 한 번 호출하면 2, 3이 남는다
func (list *LinkedList[T]) PopFront() {
	if list.root == nil {
		return
	}
	// free()기능이 없으므로 next에 들어갈 주소 값을 nil로 입력
	// list.root = list.root.next
	// list.root.next = nil
	list.root.next, list.root = nil, list.root.next
	// 노드가 완전히 제거되었을 경우 root와 tail 모두 nil을 가리켜야 한다
	if list.root == nil {
		list.tail = nil
	}
	list.count--
}

func (list *LinkedList[T]) Remove(node *Node[T]) {
	if node == list.root {
		list.PopFront()
		return
	}

	// 노드가 2개 이상인데 현재 노드의 이전 노드가 없다는 것은 리스트에 노드가 존재하지 않는다는 것이다
	prev := list.findPrevNode(node)
	if prev == nil {
		return
	}
	// prev.next = node.next
	// node.next = nil
	prev.next, node.next = node.next, nil

	// 맨 마지막에 있는 노드를 삭제할 때, tail을 현재 노드의 직전 노드의 위치를 가리키도록 갱신해 주어야 한다
	if node == list.tail {
		list.tail = prev
	}
	list.count--
}
