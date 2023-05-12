package main

import (
	"log"
	"net/http"
	"time"

	"github.com/pjw1702/study/web/go/web10/decoHandler"
	"github.com/pjw1702/study/web/go/web10/myapp"
)

// 부가 기능
func logger(w http.ResponseWriter, r *http.Request, h http.Handler) {
	start := time.Now()
	log.Print("[LOGGER1] Started")
	h.ServeHTTP(w, r)                                                           // 핸들러 수행
	log.Println("[LOGGER1] Completed time: ", time.Since(start).Milliseconds()) // 특정 시각으로 부터 걸린 시간을 밀리 초로 계산
}

func NewHandler() http.Handler {
	mux := myapp.NewHandler()
	h := decoHandler.NewDecoHandler(mux, logger)
	return h
}

func main() {
	mux := NewHandler()

	http.ListenAndServe(":3000", mux)
}
