package app_test

import (
	"bytes"
	"io"
	"mime/multipart"
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"

	"fmt"
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

func TestUploadTest(t *testing.T) {
	assert := assert.New(t) // assert 생성
	//path := `C:\images\Prometheus Architecture.PNG`
	path := "C:/images/Prometheus Architecture.PNG" // 역슬래쉬가 아닌 슬래쉬로 바꿔준다
	file, _ := os.Open(path)
	defer file.Close()

	// 웹으로 파일과 같은 데이터를 전송할 때 사용하는 포맷을 MIME 포맷이라고 한다(메일에서 파일 전송 시에 사용하는 포맷도 MIME이다)
	// MIME 포맷을 사용하기 위해 mulipart writer를 생성하여야 한다
	buf := &bytes.Buffer{} // mulipart writer를 사용하기 전에 버퍼를 우선 생성해 준 후 데이터를 버퍼에 싣는다
	writer := multipart.NewWriter(buf)
	multi, err := writer.CreateFormFile("upload_file", filepath.Base(path))
	assert.NoError(err)
	io.Copy(multi, file)
	writer.Close()

	res := httptest.NewRecorder()
	req := httptest.NewRequest("POST", "/uploads", buf)
	req.Header.Set("Content-Type", writer.FormDataContentType()) // 리퀘스트가 어떤 폼의 데이터인지를 검사해야한다 (컨텐츠 타입을 알려줘야 서버가 읽을 수 있다)

	uploadsHandler(res, req)
	assert.Equal(http.StatusOK, res.Code)

	// 업로드한 파일과 업로드 하기 위한 원본 파일이 같은지를 확인
	uploadFilePath := "./uploads/" + filepath.Base(path)
	_, err_info := os.Stat(uploadFilePath) // 해당 파일에 대한 info를 전달한다
	assert.NoError(err_info)

	uploadFile, _ := os.Open(uploadFilePath)
	originFile, _ := os.Open(path)
	defer uploadFile.Close()
	defer originFile.Close()

	// 이미지 포맷을 읽기 위해 바이트 타입의 슬라이스를 생성한다
	uploadData := []byte{}
	originData := []byte{}
	uploadFile.Read(uploadData)
	originFile.Read(originData)

	assert.Equal(originData, uploadData)
}
