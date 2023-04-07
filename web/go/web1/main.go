package main

import (
	"fmt"
	"net/http"
)

type fooHandler struct{}

func (f *fooHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello Foo!")
}

type Handler interface {
	ServeHTTP(w http.ResponseWriter, r *http.Request)
}

func main() {
	// HTTP 핸들러를 함수로 등록하는 경우: http.HandleFunc()
	// HTTP 핸들러를 인스턴스로 등록하는 경우: http.HandleFunc()

	// 인덱스 경로에 대한 HTTP 핸들러 등록
	// 첫 번째 페이지(인덱스 페이지)에 대한 리퀘스트에 대한 처리
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Hello World")
	})
	// 첫 번째 인자("/")에 입력한 경로로 요청을 받는다
	// 두 번째 인자(func() 함수)에 대한 일을 처리한다
	// func() 함수에는 사용자가 요청(Request)한 정보를 가지고 있는 두 번째 인자와 요청에 대한 응답(Response)을 할 때 웹 페이지에 writing하기 위한 정보를 가지고 있는 두 번쩨 인자가 있다

	// bar 경로에 대한 HTTP 핸들러 등록
	http.HandleFunc("/bar", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Hello Bar!")
	})

	// HTTP Handler를 정의하는 또 다른 방식
	// 첫번째 파라미터로 URL (혹은 URL 패턴)을 받아들이고, 두번째 파라미터로 http.Handler 인터페이스를 갖는 객체를 받아들인다
	http.Handle("/foo", new(fooHandler))

	http.ListenAndServe(":3000", nil)
}
