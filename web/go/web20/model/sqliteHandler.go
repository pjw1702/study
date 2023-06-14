package model

import (
	"database/sql"
	"time"

	_ "github.com/mattn/go-sqlite3"
)

type sqliteHandler struct {
	db *sql.DB
}

func (s *sqliteHandler) GetTodos() []*Todo {
	todos := []*Todo{}
	rows, err := s.db.Query("SELECT id, name, completed, createdAt FROM todos")
	if err != nil {
		panic(err)
	}
	defer rows.Close()
	for rows.Next() {
		var todo Todo
		rows.Scan(&todo.ID, &todo.Name, &todo.Completed, &todo.CreatedAt) // 읽어온 DB의 데이터를 객체에 저장
		todos = append(todos, &todo)
	}
	return todos
}

func (s *sqliteHandler) AddTodo(name string) *Todo {
	stmt, err := s.db.Prepare("INSERT INTO todos (name, completed, createdAt) VALUES (?, ?, datetime('now'))") // s.db.Prepare(): 다음 쿼리를 실행하기 위한 prepared statement를 생성하기 위한 메서드
	if err != nil {
		panic(err)
	}
	rst, err := stmt.Exec(name, false) // prepared statement에 전달 받은 쿼리 문 실행
	if err != nil {
		panic(err)
	}
	id, _ := rst.LastInsertId() // LastInsertId()메서드: "auto increment"로 설정한 컬럼의 값을 반환, 데이터를 추가한 후의 정보: 쿼리 문 실행 시, name, completed, createdAT 레코드 값은 알 수 있지만 auto increment로 설정한 id 값은 알지 못한다

	// 쿼리 실행 후에 추가된 데이터를 객체에 저장
	var todo Todo
	todo.ID = int(id)
	todo.Name = name
	todo.Completed = false
	todo.CreatedAt = time.Now()

	return &todo
}

func (s *sqliteHandler) RemoveTodo(id int) bool {
	stmt, err := s.db.Prepare("DELETE FROM todos WHERE id=?")
	if err != nil {
		panic(err)
	}
	rst, err := stmt.Exec(id)
	if err != nil {
		panic(err)
	}
	cnt, _ := rst.RowsAffected()
	return cnt > 0
}

// 기존 레코드에 completed 컬럼 값만 변경
func (s *sqliteHandler) CompleteTodo(id int, complete bool) bool {
	stmt, err := s.db.Prepare("UPDATE todos SET completed=? WHERE id=?") // id가 일치 하는 레코드만 completed 컬럼 값 변경
	if err != nil {
		panic(err)
	}
	rst, err := stmt.Exec(complete, id)
	if err != nil {
		panic(err)
	}
	cnt, _ := rst.RowsAffected() // 실행한 쿼리에 대해 변경된 컬럼 개수를 알려 줌: 업데이트가 정상적으로 이루어졌는지 여부를 확인할 수 있음
	return cnt > 0
}

func (s *sqliteHandler) Close() {
	s.db.Close()
}

func newSqliteHandler(filepath string) DBHandler {
	database, err := sql.Open("sqlite3", filepath)
	if err != nil {
		panic(err)
	}
	statement, _ := database.Prepare( // SQL Query
		`CREATE TABLE IF NOT EXISTS todos (
			id        INTEGER  PRIMARY KEY AUTOINCREMENT,
			name      TEXT,
			completed BOOLEAN,
			createdAt DATETIME
		)`)
	statement.Exec() // SQL 실행
	return &sqliteHandler{db: database}
}
