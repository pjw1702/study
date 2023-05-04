package myapp_test

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strconv"
	"strings"

	"testing"

	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"

	"encoding/json"
	"time"
)

// User struct
type User struct {
	ID        int       `json:"id"`
	Firstname string    `json:"first_name"`
	Lastname  string    `json:"last_name"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"created_at"`
}

var userMap map[int]*User
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

	w.Header().Add("Content-Type", "application/json") // 헤더 추가
	w.WriteHeader(http.StatusOK)
	data, _ := json.Marshal(user) // user를 JSON 포맷으로 마샬링
	fmt.Fprint(w, string(data))
}

func createUserHandler(w http.ResponseWriter, r *http.Request) {
	user := new(User)
	//user.CreatedAt = time.Now()
	err := json.NewDecoder(r.Body).Decode(user) // r.Body의 JSON 포맷을 user 구조체로 언마샬링
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
	userMap[user.ID] = user

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

func TestIndex(t *testing.T) {
	assert := assert.New(t)

	// 실제 서버가 아닌 테스트를 위한 목업 서버가 생성됨
	ts := httptest.NewServer(NewHandler())
	defer ts.Close()

	// Get 리퀘스트 생성
	resp, err := http.Get(ts.URL)
	assert.NoError(err)
	assert.Equal(http.StatusOK, resp.StatusCode)
	data, _ := ioutil.ReadAll(resp.Body)      // 데이터 포맷을 바이트 코드로 변환
	assert.Equal("Hello World", string(data)) // 바이트 코드를 String 타입으로 변환
}

func TestUsers(t *testing.T) {
	assert := assert.New(t)

	ts := httptest.NewServer(NewHandler())
	defer ts.Close()

	resp, err := http.Get(ts.URL + "/users") // "/users" 경로에 대한 헨들러를 등록하지 않아도 "/"(인덱스) 헨들러가 호출되어 테스트를 통과한다: "/" 하위 경로의 헨들러가 등록이 되어 있지 않으면 제일 상위인 "/" 헨들러가 호출된다
	assert.NoError(err)
	assert.Equal(http.StatusOK, resp.StatusCode)
	data, _ := ioutil.ReadAll(resp.Body)
	assert.Contains(string(data), "Get UserInfo") // Contains(): 특정 문자열이 포함되어 있어야 함을 명시하는 함수
	// assert.Contains(string(data), "Hello World")
}

func TestGetUserInfo(t *testing.T) {
	assert := assert.New(t)

	ts := httptest.NewServer(NewHandler())
	defer ts.Close()

	resp, err := http.Get(ts.URL + "/users/89") // "/users" 하위 경로의 헨들러가 정의되어 있지 않으면 등록되어 있는 상위 헨들러인 "/users" 헨들러가 호출됨
	assert.NoError(err)
	assert.Equal(http.StatusOK, resp.StatusCode)
	data, _ := ioutil.ReadAll(resp.Body)
	assert.Contains(string(data), "User Id:89")

	resp, err = http.Get(ts.URL + "/users/56")
	assert.NoError(err)
	assert.Equal(http.StatusOK, resp.StatusCode)
	data, _ = ioutil.ReadAll(resp.Body)
	assert.Contains(string(data), "User Id:56")
}

func TestCreateUser(t *testing.T) {
	assert := assert.New(t)

	ts := httptest.NewServer(NewHandler())
	defer ts.Close()

	// POST
	resp, err := http.Post(ts.URL+"/users", "application/json",
		strings.NewReader(`{"first_name":"jw", "last_name":"Park", "email":"qjw1702@gmail.com"}`)) // strings.NewReader 함수를 사용하면 문자열로 io.Reader 인터페이스를 따르는 인스턴스를 만들 수 있다(언마샬링을 하기 위함)
	assert.NoError(err)
	assert.Equal(http.StatusCreated, resp.StatusCode)
	// StatusCreated 후 user를 디코드하여 POST가 제대로 이루어졌는지 확인
	user := new(User)
	err = json.NewDecoder(resp.Body).Decode(user) // resp.Body에 있는 JSON 포맷을 디코더로 생성하여 user 객체를 디코드(user는 디코딩하기 위한 수단)
	assert.NoError(err)
	assert.NotEqual(0, user.ID)

	// POST한 결과를 GET
	id := user.ID
	resp, err = http.Get(ts.URL + "/users/" + strconv.Itoa(id)) // "/users"가 아닌 "/users/"가 들어갔는지 제대로 확인해야 한다 그렇지 않으면 unmarshalling 에러가 발생한다
	assert.NoError(err)
	assert.Equal(http.StatusOK, resp.StatusCode)

	// 또 다른 User를 사용한 인스턴스(user2)를 생성 했을 때 기존의 인스턴스(user)와의 정보가 일치하는지 확인
	user2 := new(User)
	err = json.NewDecoder(resp.Body).Decode(user2)
	assert.NoError(err)
	assert.Equal(user.ID, user2.ID)
	assert.Equal(user.Firstname, user2.Firstname) // user와 user2는 같은 info이다
}
