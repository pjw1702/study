package main

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/gorilla/pat"
	"github.com/unrolled/render"
	"github.com/urfave/negroni"
)

var rd *render.Render

type User struct {
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"created_at"`
}

func getUserInfoHandler(w http.ResponseWriter, r *http.Request) {
	user := User{Name: "jwpark", Email: "pjw1702@gmail.com"}

	rd.JSON(w, http.StatusOK, user)
	// w.Header().Add("Content-type", "application/json")
	// w.WriteHeader(http.StatusOK)
	// data, _ := json.Marshal(user)
	// fmt.Fprint(w, string(data))
}

func addUserHandler(w http.ResponseWriter, r *http.Request) {
	user := new(User)
	err := json.NewDecoder(r.Body).Decode(user)
	if err != nil {
		// w.WriteHeader(http.StatusBadRequest)
		// fmt.Fprint(w, err)
		rd.Text(w, http.StatusBadRequest, err.Error()) // err.Error() 메서드를 호출하면 err이 받은 에러 메시지를 출력한다
		return
	}
	user.CreatedAt = time.Now()
	rd.JSON(w, http.StatusOK, user)
	// w.Header().Add("Content-type", "application/json")
	// w.WriteHeader(http.StatusOK)
	// data, _ := json.Marshal(user)
	// fmt.Fprint(w, string(data))
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	// tmpl, err := template.New("Hello").ParseFiles("templates/hello.tmpl")
	// if err != nil {
	// 	// w.WriteHeader(http.StatusInternalServerError)
	// 	// fmt.Fprint(w, err)
	// 	rd.Text(w, http.StatusBadRequest, err.Error())
	// 	return
	// }
	// tmpl.ExecuteTemplate(w, "hello.tmpl", "JWPark")
	// rd.HTML(w, http.StatusOK, "body", "JWPark") // render 패키지는 템플릿을 등록할 때, 확장자를 제거하고 등록한다
	user := User{Name: "jwpark", Email: "pjw1702@gmail.com"}
	rd.HTML(w, http.StatusOK, "body", user)
}

func main() {
	// rd = render.New() // unrolled/render 패키지: JSON, XML, Binary나 httl template 등의 파일들을 response하는 코드를 구현할 때, 복잡한 코드 구현 없이 더욱 편리하게 동작할 수 있도록 제공
	// render 패키지는 기본적으로 .tmpl 파일만 읽도록 정해져있으므로 .tmpl 외의 확장자 파일도 읽도록 설정을 해야한다
	rd = render.New(render.Options{
		Directory:  "template",
		Extensions: []string{".html", ".tmpl"},
		Layout:     "hello",
	})
	mux := pat.New()

	// mux.HandleFunc("/users", getUserInfoHandler).Methods("GET")
	// mux.HandleFunc("/users", addUserHandler).Methods("PUT")""

	// gorilla/pat 사용
	mux.Get("/users", getUserInfoHandler)
	mux.Post("/users", addUserHandler)
	mux.Get("/hello", helloHandler)

	// mux.Handle("/", http.FileServer(http.Dir("public")))
	// negroni 패키지를 이용하여 http 미들웨어를 실행할 때 구현하는 부가적인 기능을 쉽게 제공받을 수 있다
	// http.FileServer()함수를 호출하여 public 디렉토리를 인자로 전달할 필요 없이, public이라는 디렉토리내에 파일이 있으면 해당 파일을 인식하여 index 페이지로 출력하도록 설계되어 있다
	// 웹 서버 실행 시, http 요청에 대한 로그 출력 기능또한 제공한다
	n := negroni.Classic()
	n.UseHandler(mux)

	// http.ListenAndServe(":3000", mux)
	http.ListenAndServe(":3000", n)
}
