package model

import "time"

type Todo struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	Completed bool      `json:"completed"`
	CreatedAt time.Time `json:"created_at"`
}

// DB 핸들러에 대한 라이프사이클(DB 인스턴스에 대한 생명주기: 핸들러를 언제부터 언제까지 사용하는지에 대한 주기)에 대한 관리가 필요
// DB 핸들러에 대한 사용 완료를 Close() 함수를 호출하는 패키지로 책임을 전가해야 함
type DBHandler interface {
	GetTodos() []*Todo
	AddTodo(name string) *Todo
	RemoveTodo(id int) bool
	CompleteTodo(id int, complete bool) bool
	Close()
}

// Map initialize
func NewDBHandler(filepath string) DBHandler {
	//handler = newMemoryHandler()
	return newSqliteHandler(filepath)
}
