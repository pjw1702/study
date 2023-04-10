// Go언어는 파일 명에 _test를 붙여서 작성하면 테스트코드로 작동한다
package app_test

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"

	"encoding/json"
	"time"
)

// smartstreets/goconvey 패키지는 코드를 테스트할 때 굉장히 유용하다: 백그라운드로 실행되면서 파일이 갱신될 때마다 테스트를 진행한다
// go get github.com/smartystreets/goconvey
// go install github.com/smartystreets/goconvey
// goconvey 명령어 실행
// localhost:8080으로 접속하여 확인
// Go 패키지 설치 시, 실행 파일은 모두 GOPATH(윈도우 기준 C:\Users\User-name\go)\bin 아래에 설치되는데, goconvey와 같은 경우엔 goconvey.exe 파일로 설치된다

// assert 패키지는 첫 번째 인자로 지정한 http 상태 코드를 확인하여 자동으로 코드에 대한 에러처리를 한다
// go get github.com/stretchr/testify/assert

// 테스트 코드를 작성할 땐 함수 명이 Test로 시작하여야 하고 인자로 testing 패키지에 T 포인터를 인자로 받아야한다

type User struct {
	FirstName string    `json:"first_name"`
	LastName  string    `json:"last_name"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"created_at"`
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello World")
}

type fooHandler struct{}

func (f *fooHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	user := new(User)
	err := json.NewDecoder(r.Body).Decode(user)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, "Bad Request: ", err)
		return
	}
	user.CreatedAt = time.Now()

	data, _ := json.Marshal(user)

	w.Header().Add("content-type", "application/json")

	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, string(data))
}

type Handler interface {
	ServeHTTP(w http.ResponseWriter, r *http.Request)
}

func barHandler(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")
	if name == "" {
		name = "World"
	}
	fmt.Fprintf(w, "Hello %s!", name)
}

// mux 구현: Go에서의 인터페이스는 메소드의 집합이라면, mux는 http 핸들러의 집합이라고 할 수 있다
func NewHttpHandler() http.Handler {
	mux := http.NewServeMux()

	mux.HandleFunc("/", indexHandler)
	mux.HandleFunc("/bar", barHandler)
	mux.Handle("/foo", new(fooHandler))
	return mux
}

func TestIndexPathHandler(t *testing.T) {
	assert := assert.New(t)

	res := httptest.NewRecorder()
	req := httptest.NewRequest("GET", "/", nil)

	// indexHandler(res, req)
	mux := NewHttpHandler()
	mux.ServeHTTP(res, req)

	assert.Equal(http.StatusOK, res.Code)
	// response의 코드에서의 호출 결과에 대한 에러처리
	// 매번 검사에 대한 에러처리 코드를 수정하거나 구현하기엔 번거롭다
	// if res.Code != http.StatusBadRequest {
	// 	t.Fatal("Failed!! ", res.Code)
	// }

	// 자동으로 http 상태 체크를 하여 response에 대한 코드를 검사 후 에러처리를 한다
	assert.Equal(http.StatusOK, res.Code)
	// 요청에 대한 실제 결과 값(response에 대한 내용)은 response내의 Body에 내장되어 있다
	// 하지만 Body는 Buffer 클래스이므로 바로 가져올 수 없다: ioutil 패키지를 통해 가져와야 한다 (혹은 마샬링 등 데이터 직렬화 과정이 필요)
	data, _ := ioutil.ReadAll(res.Body)
	assert.Equal("Hello World", string(data)) // data는 res.Body의 내용을 가져온 byte 타입의 값이므로 string으로 형 변환을 하여야 한다
}

func TestBarPathHandler_WithoutName(t *testing.T) {
	assert := assert.New(t)

	res := httptest.NewRecorder()
	req := httptest.NewRequest("GET", "/bar", nil)

	// barHandler(res, req)	// 특정 핸들러를 직접 호출하는 방식은 mux를 바람직하게 사용하는 것이 아니다(정적 라우팅)

	mux := NewHttpHandler()
	mux.ServeHTTP(res, req) // req의 인자로서 넘겨주는 경로만 적절히 변경하면 유연하게 라우팅 가능하게 하는 것이 mux의 목적이다(동적 라우팅)

	assert.Equal(http.StatusOK, res.Code)
	data, _ := ioutil.ReadAll(res.Body)
	assert.Equal("Hello World!", string(data))
}

func TestBarPathHandler_WithName(t *testing.T) {
	assert := assert.New(t)

	res := httptest.NewRecorder()
	req := httptest.NewRequest("GET", "/bar?name=jwpark", nil)

	mux := NewHttpHandler()
	mux.ServeHTTP(res, req)

	assert.Equal(http.StatusOK, res.Code)
	data, _ := ioutil.ReadAll(res.Body)
	assert.Equal("Hello jwpark!", string(data))
}

func TestFooHandler_WithoutJson(t *testing.T) {
	assert := assert.New(t)

	// 받는 데이터를 코드 상에서 직접 입력하여 요청 및 http 상태 값 응답
	res := httptest.NewRecorder()
	req := httptest.NewRequest("POST", "/foo",
		strings.NewReader(`{"first_name":"jeong woo", "last_name":"park", "email":"pjw1702@huevertech.com"}`)) // JSON 포맷으로 입력(strings.NewReader()함수에 의해서 string 타입 데이터가 버퍼 형태인 ioreader로 변환된다)

	mux := NewHttpHandler()
	mux.ServeHTTP(res, req)

	//assert.Equal(http.StatusCreated, res.Code)
	assert.Equal(http.StatusOK, res.Code)

	// 요청 받은 데이터를 비교하여 user.FirstName과 user.LastName의 값이 각각 "jeong woo"와 "park"와 일치하는지 검사
	user := new(User)
	err := json.NewDecoder(res.Body).Decode(user)
	assert.Nil(err)
	assert.Equal("jeong woo", user.FirstName)
	assert.Equal("park", user.LastName)
}
