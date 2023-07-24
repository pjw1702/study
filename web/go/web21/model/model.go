package model

import (
	"time"
)

type Todo struct {
	ID         int       `json:"id"`
	Name       string    `json:"name"`
	Compoleted bool      `json:"completed"`
	CreatedAt  time.Time `json:"CreatedAt"`
}

type DBHandler interface {
	GetTodos() []*Todo
	AddTodo(name string) *Todo
	RemoveTodo(id int) bool
	CompleteTodo(id int, complete bool) bool
	Close()
}

func NewDBHandler(filepath string) DBHandler {
	return newSqliteHandler(filepath)
}
