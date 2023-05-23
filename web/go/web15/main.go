package main

import (
	"context"
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gorilla/pat"
	"github.com/urfave/negroni"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"

	"crypto/rand"

	"github.com/joho/godotenv"
)

// var googleOauthConfig = oauth2.Config{
// 	RedirectURL:  "http://localhost:3000/auth/google/callback", // 구글에 등록할 리다이렉트 콜백 URL
// 	ClientID:     os.Getenv("GOOGLE_CLIENT_ID"),                // Getenv(): 클라이언트 ID를 OS의 환경 변수에서 가져옴
// 	ClientSecret: os.Getenv("GOOGLE_SECRET_KEY"),
// 	Scopes:       []string{"https://www.googleapis.com/auth/userinfo.email"}, // 요청하는 데이터의 스코프 지정 (해당 웹 사이트의 userinfo.email 정보로 권한을 갖고 접근하겠다는 의미)
// 	Endpoint:     google.Endpoint,                                            // 엔드포인트 지정: 구글 계정으로 로그인할 때 어느 엔드포인트로 로그인 하는지를 google 패키지에서 제공
// }

const oauthGoogleUrlAPI = "https://www.googleapis.com/oauth2/v2/userinfo?access_token=" // 유저 정보에 대해 request하는 경로

func googleLoginHandler(w http.ResponseWriter, r *http.Request) {

	var googleOauthConfig = oauth2.Config{
		RedirectURL:  "http://localhost:3000/auth/google/callback",
		ClientID:     os.Getenv("GOOGLE_CLIENT_ID"),
		ClientSecret: os.Getenv("GOOGLE_SECRET_KEY"),
		Scopes:       []string{"https://www.googleapis.com/auth/userinfo.email"},
		Endpoint:     google.Endpoint,
	}

	// 환경변수 출력 확인
	fmt.Println("env [GOOGLE_CLIENT_ID]:", googleOauthConfig.ClientID)
	fmt.Println("env [GOOGLE_SECRET_KEY]:", googleOauthConfig.ClientSecret)

	state := generateStateOauthCookie(w)                   // 브라우저의 쿠키에 임시 키를 발급하여 저장
	url := googleOauthConfig.AuthCodeURL(state)            // AuthCodURL(): URL 변조 공격 해킹을 방지하기 위해서, 리다이렉트 URL로 콜백 요청 시, 쿠키로 발급한 임시 key를 클라이언트로 부터 확인
	http.Redirect(w, r, url, http.StatusTemporaryRedirect) // 해당 웹 사이트로 로그인 시도 시, 구글 로그인 폼으로 리다이렉트
}

// 브라우저의 쿠키에 임시 키를 발급
func generateStateOauthCookie(w http.ResponseWriter) string {
	expiration := time.Now().Add(1 * 24 * time.Hour) // 현재 시각으로 부터 하룻 동안만 유효

	b := make([]byte, 16)
	rand.Read(b)                                  // 난수를 생성하여 랜덤한 값으로 바이트 슬라이스 b를 채움
	state := base64.URLEncoding.EncodeToString(b) // 랜덤 값을 base64로 인코딩
	cookie := &http.Cookie{
		Name:    "oauthstate",
		Value:   state,
		Expires: expiration,
	}
	http.SetCookie(w, cookie) // 임시 키 발급 및 쿠키에 저장
	return state
}

func googleAuthCallback(w http.ResponseWriter, r *http.Request) {
	oauthstate, _ := r.Cookie("oauthstate")

	if r.FormValue("state") != oauthstate.Value { // 해당 클라이언트의 쿠키의 값이 발급한 임시 키 값과 다른 경우
		log.Printf("Invalid google OAuth state cookie: %s state: %s\n", oauthstate.Value, r.FormValue("state"))
		http.Redirect(w, r, "/", http.StatusTemporaryRedirect) // 기본 경로로 리다이렉트
	}

	data, err := getGoogleUserInfo(r.FormValue("code")) // "code"를 통해 Google의 유저 정보를 Get
	if err != nil {
		log.Println(err.Error())
		http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
		return
	}

	fmt.Fprint(w, string(data))
}

func getGoogleUserInfo(code string) ([]byte, error) {

	var googleOauthConfig = oauth2.Config{
		RedirectURL:  "http://localhost:3000/auth/google/callback",
		ClientID:     os.Getenv("GOOGLE_CLIENT_ID"),
		ClientSecret: os.Getenv("GOOGLE_SECRET_KEY"),
		Scopes:       []string{"https://www.googleapis.com/auth/userinfo.email"},
		Endpoint:     google.Endpoint,
	}

	token, err := googleOauthConfig.Exchange(context.Background(), code) // 토큰을 발급 받음: 요청 받은 "code"와 토큰을 교환
	if err != nil {
		return nil, fmt.Errorf("Failed to Exchange %s\n", err.Error())
	}

	// 발급받은 토큰으로 유저의 정보를 request
	// AccessToken: 구글 API URL에 접근하기 위한 토큰으로써, 만료기간이 지나면 효력 소멸
	// RefreshToken: AccessToken의 효력을 지속적으로 갱신하기 위한 토큰
	resp, err := http.Get(oauthGoogleUrlAPI + token.AccessToken)
	if err != nil {
		return nil, fmt.Errorf("Failed to Get UserInfo %s\n", err.Error())
	}

	return ioutil.ReadAll(resp.Body)
}

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Println("Error loading .env file!: ", err.Error())
	}

	mux := pat.New()
	mux.HandleFunc("/auth/google/login", googleLoginHandler)
	mux.HandleFunc("/auth/google/callback", googleAuthCallback)
	n := negroni.Classic()
	n.UseHandler(mux)

	http.ListenAndServe(":3000", n)
}
