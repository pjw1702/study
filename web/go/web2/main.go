package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

// 응답 시, {"FirstName":"","LastName":"","Email":"jwpark@huevertech.com","CreatedAt":"2023-04-10T10:42:56.2613474+09:00"}와 같이, 대문자 밑 언더 스코어("_") 없는 등의 차이가 보이는 결과 값을 응답한다: JSON의 문자 인식 및 표현과 go의 문자 인식 표현이 서로 다르기 때문이다
// 맞춰주는 작업이 필요하다
//
//	type User struct {
//		FirstName string
//		LastName  string
//		Email     string
//		CreatedAt time.Time
//	}
//
// 아래와 같이, `json:"first_name"`와 같은 작업을 어노테이션(annotation)이라고 한다
// 어노테이션을 명시하면, 마샬링 시 어노테이션의 내용에 맞게 변환하여 마샬링한다
type User struct {
	FirstName string    `json:"first_name"`
	LastName  string    `json:"last_name"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"created_at"`
}

type fooHandler struct{}

func (f *fooHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// request를 JSON 포맷으로 받는다
	// 리퀘스트 body에 데이터를 입력하는 방식이므로 URL?argument=value 형태로 웹 브라우저를 통해 요청하는 방식이 아니다
	user := new(User)                           // User 구조체 인스턴스 생성
	err := json.NewDecoder(r.Body).Decode(user) // JSON 포맷으로 파싱(Request 내의 Body에 JSON 포맷이 들어있다)하여 디코드(user의 데이터를 JSON포맷의 항목으로 한다)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, "Bad Request: ", err)
		return
	}
	user.CreatedAt = time.Now()

	data, _ := json.Marshal(user) // user의 데이터로 디코딩 하였던 JSON 포맷을 다시 JSON 포맷으로 인코딩(byte 타입의 array로 마샬링한다)

	// {"first_name":"tucker","last_name":"kim","email":"jwpark@huevertech.com","created_at":"2023-04-10T11:05:48.7380883+09:00"}과 같이 response 값을 JSON 타입이 아닌 TEXT 타입으로 응답한다(content-type: text/plain)
	// 헤더의 content-type을 바꿔준다
	w.Header().Add("content-type", "application/json")

	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, string(data)) // 마샬링한 data의 값은 byte이므로 string 타입으로 형 변환한다
}

type Handler interface {
	ServeHTTP(w http.ResponseWriter, r *http.Request)
}

func barHandler(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name") // 해당 URL로 요청받은 Request 값에서, name이라는 argument를 쿼리하여 그 argument의 값을 Get한다
	// localhost:3000/bar?name=bar
	if name == "" {
		name = "World"
	}
	fmt.Fprintf(w, "Hello %s!", name) // name 출력
}

func main() {

	// 정적 등록

	// HTTP 핸들러를 함수로 등록하는 경우: http.HandleFunc()
	// HTTP 핸들러를 인스턴스로 등록하는 경우: http.HandleFunc()

	// 인덱스 경로에 대한 HTTP 핸들러 등록
	// 첫 번째 페이지(인덱스 페이지)에 대한 리퀘스트에 대한 처리
	// http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
	// 	fmt.Fprint(w, "Hello World")
	// })
	// 첫 번째 인자("/")에 입력한 경로로 요청을 받는다
	// 두 번째 인자(func() 함수)에 대한 일을 처리한다
	// func() 함수에는 사용자가 요청(Request)한 정보를 가지고 있는 두 번째 인자와 요청에 대한 응답(Response)을 할 때 웹 페이지에 writing하기 위한 정보를 가지고 있는 두 번쩨 인자가 있다

	// bar 경로에 대한 HTTP 핸들러 등록
	// http.HandleFunc("/bar", func(w http.ResponseWriter, r *http.Request) {
	// 	fmt.Fprint(w, "Hello Bar!")
	// })
	// http.HandleFunc("/bar", barHandler)

	// HTTP Handler를 정의하는 또 다른 방식
	// 첫번째 파라미터로 URL (혹은 URL 패턴)을 받아들이고, 두번째 파라미터로 http.Handler 인터페이스를 갖는 객체를 받아들인다
	// http.Handle("/foo", new(fooHandler))	// new() 함수: 인스턴스 생성

	// 동적 등록

	// HTTP 핸들러를 경로에 따라 다르게 등록(각각의 경로로 분배)하는 것을 Mux라고 한다(라우터 라고도 함)
	// 동적 방식은 http에 정적으로 등록하는 것과 동일하게 적용하지만 라우터 인스턴스를 별도로 생성하여 그 인스턴스에 경로를 할당하는 방식으로 작용한다
	mux := http.NewServeMux() // 라우터 인스턴스 생성

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Hello World")
	})
	mux.HandleFunc("/bar", barHandler)
	mux.Handle("/foo", new(fooHandler))

	http.ListenAndServe(":3000", mux)
}
