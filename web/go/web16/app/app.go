package app

import (
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/unrolled/render"
)

var rd *render.Render

type Todo struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	Completed bool      `json:"completed"`
	CreatedAt time.Time `json:"created_at"`
}

var todoMap map[int]*Todo // 맵이 데이터 영역에 할당되어 힙 메모리에 할당된 Todo 객체를 value로 갖고 있으므로 프로그램 내의 모든 스코프에서 todoMap 맵을 이용하여 Todo 객체에 접근 가능

func indexHandler(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "/todo.html", http.StatusTemporaryRedirect)
}

func getTodoListHandler(w http.ResponseWriter, r *http.Request) {
	list := []*Todo{}
	for _, v := range todoMap {
		list = append(list, v)
	}
	rd.JSON(w, http.StatusOK, list)
}

// 아이템 추가에 대한 테스트 코드
func addTestTodos() {
	todoMap[1] = &Todo{1, "Buy a milk", false, time.Now()}
	todoMap[2] = &Todo{2, "Excercise", true, time.Now()}
	todoMap[3] = &Todo{2, "Home work", false, time.Now()}
}

func MakeHandler() http.Handler {
	todoMap = make(map[int]*Todo)
	addTestTodos()

	rd = render.New()
	r := mux.NewRouter()

	r.HandleFunc("/todos", getTodoListHandler).Methods("GET")
	r.HandleFunc("/", indexHandler)

	return r
}
