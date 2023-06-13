package model

import "time"

type Todo struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	Completed bool      `json:"completed"`
	CreatedAt time.Time `json:"created_at"`
}

var todoMap map[int]*Todo // 맵이 데이터 영역에 할당되어 힙 메모리에 할당된 Todo 객체를 value로 갖고 있으므로 프로그램 내의 모든 스코프에서 todoMap 맵을 이용하여 Todo 객체에 접근 가능(메모리 맵)

// Map initialize
func init() {
	todoMap = make(map[int]*Todo)
}

func GetTodos() []*Todo {
	list := []*Todo{}
	for _, v := range todoMap {
		list = append(list, v)
	}
	return list
}

func AddTodo(name string) *Todo {
	id := len(todoMap) + 1
	todo := &Todo{id, name, false, time.Now()}
	todoMap[id] = todo

	return todo
}

func RemoveTodo(id int) bool {
	if _, ok := todoMap[id]; ok {
		delete(todoMap, id) // 해당 id에 대한 맵 삭제
		return true
	}

	return false
}

func CompleteTodo(id int, complete bool) bool {
	if todo, ok := todoMap[id]; ok {
		todo.Completed = complete
		return true
	}

	return false
}
