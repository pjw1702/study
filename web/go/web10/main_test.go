package main

import (
	"bufio"
	"bytes"
	"io/ioutil"
	"log"
	"testing"

	"github.com/stretchr/testify/assert"

	"net/http"
	"net/http/httptest"
)

func TestIndexPage(t *testing.T) {
	assert := assert.New(t)

	ts := httptest.NewServer(NewHandler())
	defer ts.Close()

	resp, err := http.Get(ts.URL)
	assert.NoError(err)
	assert.Equal(http.StatusOK, resp.StatusCode)
	data, _ := ioutil.ReadAll(resp.Body)
	assert.Equal("Hello World", string(data))
}

func TestDecoHandler(t *testing.T) {
	assert := assert.New(t)

	ts := httptest.NewServer(NewHandler())
	defer ts.Close()

	// 메인 폴더에 로그가 찍히는 지 확인
	buf := &bytes.Buffer{} // 버퍼 생성
	log.SetOutput(buf)     // Standard 로그를 저장하기 위한 버퍼 생성

	resp, err := http.Get(ts.URL)
	assert.NoError(err)
	assert.Equal(http.StatusOK, resp.StatusCode)

	// Get 요청 시 찍히는 로그가 버퍼에 잘 저장되었는지 확인
	r := bufio.NewReader(buf)
	line, _, err := r.ReadLine() // 버퍼에 저장되어 있는 데이터를 한 줄씩 읽음
	assert.NoError(err)
	assert.Contains(string(line), "[LOGGER1] Started")

	line, _, err = r.ReadLine()
	assert.NoError(err)
	assert.Contains(string(line), "[LOGGER1] Completed")
}
