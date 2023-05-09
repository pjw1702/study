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

// PUT 요청 시, 특정 항목을 공백으로 요청을 하면 디폴트 값으로서의 공백이 아닌, 항목에 대한 빈 문자열으로 간주하여 해당 항목을 공백으로 갱신하지 않고 기존의 데이터를 그대로 유지시킨다
// 빈 문자열은 업데이트하지 않는다
// 실무에서는 업데이트용 객체를 따로 생성하여 업데이트할 항목만 플래그를 입력하여 업데이트한다
// type UpdateUser struct {
// 	ID        int       `json:"id"`
// 	Firstname string    `json:"first_name"`
// 	Lastname  string    `json:"last_name"`
// 	Email     string    `json:"email"`
// 	CreatedAt time.Time `json:"created_at"`

// 	UpdatedFirstName	bool	`json:"updated_first_name"`
// 	UpdatedLastName	bool	`json:"updated_last_name"`
// 	UpdatedEmail	bool	`json:"updated_email"`
// }

var userMap map[int]*User
var lastID int

func indexHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello World")
}

func usersHandler(w http.ResponseWriter, r *http.Request) {
	if len(userMap) == 0 {
		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w, "No Users")
		return
	}
	users := []*User{}
	for _, u := range userMap {
		users = append(users, u)
	}
	data, _ := json.Marshal(users)
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, string(data))
}

func getUserInfoHandler(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		w.WriteHeader(http.StatusBadGateway)
		fmt.Fprint(w, err)
		return
	}
	user, ok := userMap[id]
	if !ok {
		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w, "No User Id: ", id)
		return
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	data, _ := json.Marshal(user)
	fmt.Fprint(w, string(data))
}

func createUserHandler(w http.ResponseWriter, r *http.Request) {
	user := new(User)
	err := json.NewDecoder(r.Body).Decode(user)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, err)
		return
	}

	// Created User
	lastID++
	user.ID = lastID
	user.CreatedAt = time.Now()
	userMap[user.ID] = user

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	data, _ := json.Marshal(user)
	fmt.Fprint(w, string(data))
}

func deleteUserInfoHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		w.WriteHeader(http.StatusBadGateway)
		fmt.Fprint(w, err)
		return
	}
	_, ok := userMap[id] // 요청한 id가 맵에 없는 경우
	if !ok {
		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w, "No User Id: ", id)
		return
	}

	delete(userMap, id) // 맵에서 해당 id 유저를 삭제
	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, "Deleted User Id: ", id)
}

func updateUserHandler(w http.ResponseWriter, r *http.Request) {
	updateUser := new(User)
	err := json.NewDecoder(r.Body).Decode(updateUser) // json.NewDecoder(): JSON body에 대한 io reader
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, err)
		return
	}
	user, ok := userMap[updateUser.ID]
	if !ok {
		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w, "No User Id:", updateUser.ID)
		return
	}
	// 갱신 요청에 대한 항목 중, Firstname, Lastname, Email 항목이 비어있는 경우 공백으로 업데이트 하지 않고 기존 정보를 유지하기 위함
	if updateUser.Firstname != "" {
		user.Firstname = updateUser.Firstname
	}
	if updateUser.Lastname != "" {
		user.Lastname = updateUser.Lastname
	}
	if updateUser.Email != "" {
		user.Email = updateUser.Email
	}

	// Update
	//userMap[updateUser.ID] = user // userMap[]은 포인터 타입(*User)이므로 굳이 값을 대입시켜 줄 필요는 없다
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	//data, _ := json.Marshal(updateUser)
	data, _ := json.Marshal(user)
	fmt.Fprint(w, string(data))
}

// NewHandler make a new myapp handler
func NewHandler() http.Handler {
	userMap = make(map[int]*User)
	lastID = 0
	mux := mux.NewRouter()

	// 핸들러 등록
	mux.HandleFunc("/", indexHandler)
	mux.HandleFunc("/users", usersHandler).Methods("GET")
	mux.HandleFunc("/users", createUserHandler).Methods("POST")
	mux.HandleFunc("/users", updateUserHandler).Methods("PUT")
	mux.HandleFunc("/users/{id:[0-9]+}", getUserInfoHandler).Methods("GET")
	mux.HandleFunc("/users/{id:[0-9]+}", deleteUserInfoHandler).Methods("DELETE")

	return mux
}
