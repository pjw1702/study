package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/antage/eventsource"
	"github.com/gorilla/pat"
	"github.com/urfave/negroni"

	"time"
)

type Message struct {
	Name string `json:"name"`
	Msg  string `json:"msg"`
}

var msgCh chan Message

func postMessageHandler(w http.ResponseWriter, r *http.Request) {
	msg := r.FormValue("msg")
	name := r.FormValue("name")
	// log.Println("postMessageHandler ", msg, name)
	sendMessage(name, msg)
}

func addUserHandler(w http.ResponseWriter, r *http.Request) {
	username := r.FormValue("name")
	// 등록된 유저의 이름을 해당 채팅에 연결된 다른 유저들에게 알림
	sendMessage("", fmt.Sprintf("add user: %s", username))
}

func leftUserHandler(w http.ResponseWriter, r *http.Request) {
	username := r.FormValue("username")
	// 삭제된 유저의 이름을 해당 채팅에 연결된 다른 유저들에게 알림
	sendMessage("", fmt.Sprintf("left user: %s", username))
}

func sendMessage(name, msg string) {
	// send message to every client
	msgCh <- Message{name, msg}
}

func processMsgCh(es eventsource.EventSource) {
	for msg := range msgCh {
		data, _ := json.Marshal(msg)
		es.SendEventMessage(string(data), "", strconv.Itoa(time.Now().Nanosecond()))
	}
}

func main() {

	// 이벤트 소스 소켓 오픈
	msgCh = make(chan Message)
	es := eventsource.New(nil, nil)
	defer es.Close()

	// 메세지 채널 실행
	go processMsgCh(es)

	mux := pat.New()
	mux.Post("/messages", postMessageHandler)
	mux.Handle("/stream", es) // 클라이언트가 '/stream' 경로로 요청시 이벤트 소스 소켓과 자동으로 연결 됨
	mux.Post("/users", addUserHandler)
	mux.Delete("/users", leftUserHandler)

	n := negroni.Classic()
	n.UseHandler(mux)

	http.ListenAndServe(":3000", n)
}
