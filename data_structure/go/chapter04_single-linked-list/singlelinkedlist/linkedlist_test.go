package singlelinkedlist

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPushBack(test *testing.T) {
	var list LinkedList[int]

	// 초기 상태이므로 nil인지 테스트
	assert.Nil(test, list.root)
	assert.Nil(test, list.tail)

	list.PushBack(1)

	// 뒷 방향으로 1을 삽입했으므로 NotNil인지 테스트
	assert.NotNil(test, list.root)

	// list.Front() 혹은 list.Back()에 1이 삽입되었는지 테스트
	assert.Equal(test, 1, list.Front().Value)
	assert.Equal(test, 1, list.Back().Value)

	list.PushBack(2)
	assert.NotNil(test, list.root)
	assert.Equal(test, 1, list.Front().Value)
	assert.Equal(test, 2, list.Back().Value)

	list.PushBack(3)
	assert.NotNil(test, list.root)
	assert.Equal(test, 1, list.Front().Value)
	assert.Equal(test, 3, list.Back().Value)

	// 데이터를 3개 삽입했을 때 길이가 3인지 테스트
	assert.Equal(test, 3, list.Count())
	assert.Equal(test, 3, list.Count2())

	// 각 노드마다의 데이터 테스트
	assert.Equal(test, 1, list.GetAfter(0).Value)
	assert.Equal(test, 2, list.GetAfter(1).Value)
	assert.Equal(test, 3, list.GetAfter(2).Value)
	assert.Nil(test, list.GetAfter(3))
}

func TestPushFront(test *testing.T) {
	var list LinkedList[int]

	assert.Nil(test, list.root)
	assert.Nil(test, list.tail)

	list.PushFront(1)
	assert.NotNil(test, list.root)
	assert.Equal(test, 1, list.Front().Value)
	assert.Equal(test, 1, list.Back().Value)

	list.PushFront(2)
	assert.NotNil(test, list.root)
	assert.Equal(test, 2, list.Front().Value)
	assert.Equal(test, 1, list.Back().Value)

	list.PushFront(3)
	assert.NotNil(test, list.root)
	assert.Equal(test, 3, list.Front().Value)
	assert.Equal(test, 1, list.Back().Value)

	assert.Equal(test, 3, list.Count())
	assert.Equal(test, 3, list.Count2())

	list.PushFront(4)
	assert.NotNil(test, list.root)
	assert.Equal(test, 4, list.Front().Value)
	assert.Equal(test, 1, list.Back().Value)

	assert.Equal(test, 4, list.Count())
	assert.Equal(test, 4, list.Count2())
}

func TestInsertAfter(test *testing.T) {
	var list LinkedList[int]

	list.PushBack(1)
	list.PushBack(2)
	list.PushBack(3)

	node := list.GetAfter(1)
	list.InsertAfter(node, 4)

	assert.Equal(test, 4, list.Count2())
	assert.Equal(test, 4, list.GetAfter(2).Value)
	assert.Equal(test, 3, list.Back().Value)
}

func TestInsertBefore(test *testing.T) {
	var list LinkedList[int]

	list.PushBack(1)
	list.PushBack(2)
	list.PushBack(3)

	node := list.GetAfter(1)
	list.InsertBefore(node, 4)

	assert.Equal(test, 4, list.Count2())
	assert.Equal(test, 4, list.GetAfter(1).Value)
	assert.Equal(test, 2, list.GetAfter(2).Value)
	assert.Equal(test, 3, list.Back().Value)

	// 리스트에 없는 노드 앞에 삽입 시도
	tempNode := &Node[int]{
		Value: 100,
	}
	list.InsertBefore(tempNode, 200)
	assert.Equal(test, 4, list.Count())
	assert.Equal(test, 4, list.Count2())
}

// 노드가 하나밖에 없는 경우에 현재 노드의 앞에 삽입하는 경우
func TestInsertBeforeRoot(test *testing.T) {
	var list LinkedList[int]

	list.PushBack(1)
	list.InsertBefore(list.GetAfter(0), 4)

	assert.Equal(test, 2, list.Count2())
	assert.Equal(test, 4, list.Front().Value)
	assert.Equal(test, 1, list.Back().Value)

}

func TestPopFront(test *testing.T) {
	var list LinkedList[int]

	list.PushBack(1)
	list.PushBack(2)
	list.PushBack(3)
	list.PopFront()

	assert.Equal(test, 2, list.Count())
	assert.Equal(test, 2, list.Count2())
	assert.Equal(test, 2, list.Front().Value)
	assert.Equal(test, 3, list.Back().Value)

	list.PopFront()

	assert.Equal(test, 1, list.Count())
	assert.Equal(test, 1, list.Count2())
	assert.Equal(test, 3, list.Front().Value)
	assert.Equal(test, 3, list.Back().Value)

	list.PopFront()

	assert.Equal(test, 0, list.Count())
	assert.Equal(test, 0, list.Count2())
	assert.Nil(test, list.Front())
	assert.Nil(test, list.Back())
}

func TestRemove(test *testing.T) {
	var list LinkedList[int]

	list.PushBack(1)
	list.PushBack(2)
	list.PushBack(3)

	list.Remove(list.GetAfter(1))
	assert.Equal(test, 2, list.Count())
	assert.Equal(test, 2, list.Count2())
	assert.Equal(test, 1, list.Front().Value)
	assert.Equal(test, 3, list.Back().Value)

	list.Remove(list.GetAfter(0))
	assert.Equal(test, 1, list.Count())
	assert.Equal(test, 1, list.Count2())
	assert.Equal(test, 3, list.Front().Value)
	assert.Equal(test, 3, list.Back().Value)

	// 없는 노드를 삭제하려 시도하므로 리스트의 갯수와 데이터에 변함이 없어야 한다
	list.Remove(&Node[int]{
		Value: 100,
	})
	assert.Equal(test, 1, list.Count())
	assert.Equal(test, 1, list.Count2())
	assert.Equal(test, 3, list.Front().Value)
	assert.Equal(test, 3, list.Back().Value)

	list.Remove(list.GetAfter(0))
	assert.Equal(test, 0, list.Count())
	assert.Equal(test, 0, list.Count2())
	assert.Nil(test, list.Front())
	assert.Nil(test, list.Back())
}
