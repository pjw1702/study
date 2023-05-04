package myapp

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"strconv"

	"github.com/gorilla/mux"
)

// User struct
type User struct {
	ID        int       `json:"id"`
	Firstname string    `json:"first_name"`
	Lastname  string    `json:"last_name"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"created_at"`
}

var userMap map[int]*User // 본 프로그램에서는 DB를 사용하지 않으므로 DB 대신 해시 테이블(맵)에 저장
var lastID int

func indexHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello World")
}

func usersHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Get UserInfo by /users/{id}")
}

func getUserInfoHandler(w http.ResponseWriter, r *http.Request) {
	// vars := mux.Vars(r) // id 파싱
	// fmt.Fprint(w, "User Id:", vars["id"])	// POST를 하는 경우 json 포맷인 '{'가 먼저 와야하는데 'U'가 먼저 오므로 포맷 불일치 에러 발생
	vars := mux.Vars(r) // Mux.Vars는 String 타입 맵을 반환
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		w.WriteHeader(http.StatusBadGateway)
		fmt.Fprint(w, err)
		return
	}
	user, ok := userMap[id]
	if !ok {
		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w, "No User Id:", id)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Add("Content-Type", "application/json") // 헤더 추가
	data, _ := json.Marshal(user)                      // user를 JSON 포맷으로 마샬링
	fmt.Fprint(w, string(data))
}

func createUserHandler(w http.ResponseWriter, r *http.Request) {
	user := new(User)
	err := json.NewDecoder(r.Body).Decode(user) // user 파싱
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, err)
		return
	}

	// Created User
	//user.ID = 2	// ID를 수동으로 지정하는 방식이 아닌 매번 1 씩 증가하는, 동적으로 지정하는 방식으로 교체
	lastID++
	user.ID = lastID
	user.CreatedAt = time.Now()
	userMap[user.ID] = user // 해시 테이블에 POST 요청으로 받은 정보를 User 객체를 통해 저장

	w.Header().Add("Content-Type", "application/json") // 헤더 추가
	w.WriteHeader(http.StatusCreated)
	data, _ := json.Marshal(user)
	fmt.Fprint(w, string(data))
}

// NewHandler make a new myapp handler
func NewHandler() http.Handler {
	userMap = make(map[int]*User)
	lastID = 0
	// mux := http.NewServeMux()
	mux := mux.NewRouter()

	// 핸들러 등록
	mux.HandleFunc("/", indexHandler)
	//mux.HandleFunc("/users", usersHandler)
	mux.HandleFunc("/users", usersHandler).Methods("GET")
	mux.HandleFunc("/users", createUserHandler).Methods("POST")
	mux.HandleFunc("/users/{id:[0-9]+}", getUserInfoHandler)

	return mux
}
