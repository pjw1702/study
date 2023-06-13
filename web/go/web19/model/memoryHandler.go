package model

import "time"

// DB Memory 핸들링: 외부 공개 X
type memoryHandler struct {
	todoMap map[int]*Todo
}

func (m *memoryHandler) getTodos() []*Todo {
	list := []*Todo{}
	for _, v := range m.todoMap {
		list = append(list, v)
	}
	return list
}
func (m *memoryHandler) addTodo(name string) *Todo {
	id := len(m.todoMap) + 1
	todo := &Todo{id, name, false, time.Now()}
	m.todoMap[id] = todo

	return todo
}
func (m *memoryHandler) removeTodo(id int) bool {
	if _, ok := m.todoMap[id]; ok {
		delete(m.todoMap, id) // 해당 id에 대한 맵 삭제
		return true
	}

	return false
}
func (m *memoryHandler) completeTodo(id int, complete bool) bool {
	if todo, ok := m.todoMap[id]; ok {
		todo.Completed = complete
		return true
	}

	return false
}

func newMemoryHandler() dbHandler {
	m := memoryHandler{}
	m.todoMap = make(map[int]*Todo) // Map initialize
	return &m
}
