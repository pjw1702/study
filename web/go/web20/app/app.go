package app

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/unrolled/render"

	"strconv"

	"github.com/pjw1702/study/web/go/web20/model"
)

// todoMap의 맵 리스트를 model 패키지의 GetTodos() 함수를 통해 가져온다

var rd *render.Render = render.New()

// app 패키지 또한 DB 핸들러 사용에 대해, Close() 함수 호출을 통해 사용 완료 책임을 가져야 한다
type AppHandler struct {
	http.Handler
	db model.DBHandler
}

func (a *AppHandler) indexHandler(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "/todo.html", http.StatusTemporaryRedirect)
}

func (a *AppHandler) getTodoListHandler(w http.ResponseWriter, r *http.Request) {
	// list := []*model.Todo{}
	// for _, v := range todoMap {
	// 	list = append(list, v)
	// }
	list := a.db.GetTodos()
	rd.JSON(w, http.StatusOK, list)
}

// add 버튼 클릭 이벤트 처리
func (a *AppHandler) addTodoHandler(w http.ResponseWriter, r *http.Request) {
	name := r.FormValue("name")
	// id := len(todoMap) + 1
	// todo := &model.Todo{id, name, false, time.Now()}
	// todoMap[id] = todo
	todo := a.db.AddTodo(name)
	//rd.JSON(w, http.StatusOK, todo)
	rd.JSON(w, http.StatusCreated, todo)
}

type Success struct {
	Success bool `json:"success"`
}

func (a *AppHandler) removeTodoHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r) // delete 요청을 한 item의 정보를 맵으로 전달
	id, _ := strconv.Atoi(vars["id"])
	// if _, ok := todoMap[id]; ok {
	// 	delete(todoMap, id)                      // 해당 id에 대한 맵 삭제
	// 	rd.JSON(w, http.StatusOK, Success{true}) // 응답 시 성공 여부까지 같이 알림
	// } else {
	// 	rd.JSON(w, http.StatusOK, Success{false})
	// }
	ok := a.db.RemoveTodo(id)
	if ok {
		rd.JSON(w, http.StatusOK, Success{true})
	} else {
		rd.JSON(w, http.StatusOK, Success{false})
	}

}

func (a *AppHandler) completeTodoHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r) // delete 요청을 한 item의 정보를 맵으로 전달
	id, _ := strconv.Atoi(vars["id"])
	complete := r.FormValue("complete") == "true"
	// if todo, ok := todoMap[id]; ok {
	// 	todo.Completed = complete
	// 	rd.JSON(w, http.StatusOK, Success{true})
	// } else {
	// 	rd.JSON(w, http.StatusOK, Success{false})
	// }
	ok := a.db.CompleteTodo(id, complete)
	if ok {
		rd.JSON(w, http.StatusOK, Success{true})
	} else {
		rd.JSON(w, http.StatusOK, Success{false})
	}
}

func (a *AppHandler) Close() {
	a.db.Close()
}

func MakeHandler(filepath string) *AppHandler {
	// todoMap = make(map[int]*model.Todo)

	//rd = render.New()
	r := mux.NewRouter()
	a := &AppHandler{
		Handler: r,
		db:      model.NewDBHandler(filepath),
	}

	r.HandleFunc("/todos", a.getTodoListHandler).Methods("GET")
	r.HandleFunc("/todos", a.addTodoHandler).Methods("POST")
	r.HandleFunc("/todos/{id:[0-9]+}", a.removeTodoHandler).Methods("DELETE")
	r.HandleFunc("/complete-todo/{id:[0-9]+}", a.completeTodoHandler).Methods("GET")
	r.HandleFunc("/", a.indexHandler)

	return a
}
