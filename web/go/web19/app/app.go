package app

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/unrolled/render"

	"strconv"

	"github.com/pjw1702/study/web/go/web19/model"
)

// todoMap의 맵 리스트를 model 패키지의 GetTodos() 함수를 통해 가져온다

var rd *render.Render

func indexHandler(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "/todo.html", http.StatusTemporaryRedirect)
}

func getTodoListHandler(w http.ResponseWriter, r *http.Request) {
	// list := []*model.Todo{}
	// for _, v := range todoMap {
	// 	list = append(list, v)
	// }
	list := model.GetTodos()
	rd.JSON(w, http.StatusOK, list)
}

// add 버튼 클릭 이벤트 처리
func addTodoHandler(w http.ResponseWriter, r *http.Request) {
	name := r.FormValue("name")
	// id := len(todoMap) + 1
	// todo := &model.Todo{id, name, false, time.Now()}
	// todoMap[id] = todo
	todo := model.AddTodo(name)
	//rd.JSON(w, http.StatusOK, todo)
	rd.JSON(w, http.StatusCreated, todo)
}

type Success struct {
	Success bool `json:"success"`
}

func removeTodoHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r) // delete 요청을 한 item의 정보를 맵으로 전달
	id, _ := strconv.Atoi(vars["id"])
	// if _, ok := todoMap[id]; ok {
	// 	delete(todoMap, id)                      // 해당 id에 대한 맵 삭제
	// 	rd.JSON(w, http.StatusOK, Success{true}) // 응답 시 성공 여부까지 같이 알림
	// } else {
	// 	rd.JSON(w, http.StatusOK, Success{false})
	// }
	ok := model.RemoveTodo(id)
	if ok {
		rd.JSON(w, http.StatusOK, Success{true})
	} else {
		rd.JSON(w, http.StatusOK, Success{false})
	}

}

func completeTodoHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r) // delete 요청을 한 item의 정보를 맵으로 전달
	id, _ := strconv.Atoi(vars["id"])
	complete := r.FormValue("complete") == "true"
	// if todo, ok := todoMap[id]; ok {
	// 	todo.Completed = complete
	// 	rd.JSON(w, http.StatusOK, Success{true})
	// } else {
	// 	rd.JSON(w, http.StatusOK, Success{false})
	// }
	ok := model.CompleteTodo(id, complete)
	if ok {
		rd.JSON(w, http.StatusOK, Success{true})
	} else {
		rd.JSON(w, http.StatusOK, Success{false})
	}
}

func MakeHandler() http.Handler {
	// todoMap = make(map[int]*model.Todo)

	rd = render.New()
	r := mux.NewRouter()

	r.HandleFunc("/todos", getTodoListHandler).Methods("GET")
	r.HandleFunc("/todos", addTodoHandler).Methods("POST")
	r.HandleFunc("/todos/{id:[0-9]+}", removeTodoHandler).Methods("DELETE")
	r.HandleFunc("/complete-todo/{id:[0-9]+}", completeTodoHandler).Methods("GET")
	r.HandleFunc("/", indexHandler)

	return r
}
