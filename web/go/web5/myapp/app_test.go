package myapp_test

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
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
