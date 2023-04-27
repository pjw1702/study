package myapp

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func indexHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello World")
}

func usersHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Get UserInfo by /users/{id}")
}

func getUserInfoHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r) // id 파싱
	fmt.Fprint(w, "User Id:", vars["id"])
}

// NewHandler make a new myapp handler
func NewHandler() http.Handler {
	// mux := http.NewServeMux()
	mux := mux.NewRouter() // NewServeMux() 대신 Gorilla Mux 사용

	// 핸들러 등록
	mux.HandleFunc("/", indexHandler)
	mux.HandleFunc("/users", usersHandler)
	//mux.HandleFunc("/users/89", getUserInfo89Handler)	// id를 고정으로 지정하면 해당 id에 대한 헨들러를 일일이 모두 등록해야 한다: 요청하는 id에 대한 응답 처리를 동적으로 하기 위해 gorilla mux를 이용
	mux.HandleFunc("/users/{id:[0-9]+}", getUserInfoHandler) // github.com에서 gorilla mux 사용법 참고

	return mux
}
