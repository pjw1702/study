package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func uploadsHandler(w http.ResponseWriter, r *http.Request) {
	// 웹 브라우저로 부터 파일 업로드 받음
	uploadFile, header, err := r.FormFile("upload_file") // 전송된 파일은 리퀘스트에 실려서 오는데, 해당 파일을 읽어야 함: input Form으로 날라온 파일을 읽는 함수이다
	if err != nil {
		w.WriteHeader(http.StatusBadRequest) // 에러 내용 반환
		fmt.Fprint(w, err)
		return
	}
	defer uploadFile.Close()

	// 디렉토리 및 파일 핸들 생성(생성 경로 지정 포함): 저장할 파일 공간 생성
	dirname := "./uploads"
	os.MkdirAll(dirname, 0777)
	filepath := fmt.Sprintf("%s/%s", dirname, header.Filename) // 파일 이름은 헤더에 들어있다
	file, err := os.Create(filepath)
	defer file.Close() // 생성한 파일은 항상 닫아줘야 한다: 파일을  생성하면 파일 핸들을 사용하게 되는데, 이 파일 핸들은 OS의 자원이므로 닫아줘지 않으면 계속 자원을 사용하게 된다
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprint(w, err)
		return
	}

	// 저장할 파일 공간에 업로드 받은 파일 복사
	io.Copy(file, uploadFile)
	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, filepath)
}

func main() {
	http.HandleFunc("/uploads", uploadsHandler)
	http.Handle("/", http.FileServer(http.Dir("public")))

	http.ListenAndServe(":3000", nil)
}
