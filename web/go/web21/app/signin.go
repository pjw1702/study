package app

import (
	"context"
	"crypto/rand"
	"encoding/base64"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"

	"fmt"
)

//	{
//	  "id": "114343132594774220545",
//	  "email": "qjw1702@gmail.com",
//	  "verified_email": true,
//	  "picture": "https://lh3.googleusercontent.com/a/default-user=s96-c"
//	}
type GoogleUserId struct {
	ID            string `json:"id"`
	Email         string `json:"email"`
	VerifiedEmail bool   `json:"verified_email"`
	Picture       string `json:"picture"`
}

const oauthGoogleUrlAPI = "https://www.googleapis.com/oauth2/v2/userinfo?access_token="

func googleLoginHandler(w http.ResponseWriter, r *http.Request) {

	var googleOauthConfig = oauth2.Config{
		RedirectURL:  "http://localhost:3000/auth/google/callback",
		ClientID:     os.Getenv("GOOGLE_CLIENT_ID"),
		ClientSecret: os.Getenv("GOOGLE_SECRET_KEY"),
		Scopes:       []string{"https://www.googleapis.com/auth/userinfo.email"},
		Endpoint:     google.Endpoint,
	}

	// Check environment values
	fmt.Println("env [GOOGLE_CLIENT_ID]:", googleOauthConfig.ClientID)
	fmt.Println("env [GOOGLE_SECRET_KEY]:", googleOauthConfig.ClientSecret)

	state := generateStateOauthCookie(w)                   // Generate temporary key for authentication to cookie
	url := googleOauthConfig.AuthCodeURL(state)            // Authenticate call back url through temporary key located cookie from client and return
	http.Redirect(w, r, url, http.StatusTemporaryRedirect) // Redirect call back url
}

func generateStateOauthCookie(w http.ResponseWriter) string {
	expiration := time.Now().Add(1 * 24 * time.Hour) // validate time for a day from current gime

	b := make([]byte, 16)
	rand.Read(b)                                  // Generate random numbers and store them in a byte slice
	state := base64.URLEncoding.EncodeToString(b) // Encode to base64(temporary key)
	cookie := &http.Cookie{
		Name:    "oauthstate",
		Value:   state,
		Expires: expiration,
	}
	http.SetCookie(w, cookie) // Getnerate temporary key for authentictaion and store cookie where client
	return state
}

func googleAuthCallback(w http.ResponseWriter, r *http.Request) {
	oauthstate, _ := r.Cookie("oauthstate")

	if r.FormValue("state") != oauthstate.Value { // Case of not match throuth compare temporary key where cookie with where server
		errMsg := fmt.Sprintf("Invalid google OAuth state cookie: %s state: %s\n", oauthstate.Value, r.FormValue("state"))
		log.Printf(errMsg)
		http.Error(w, errMsg, http.StatusInternalServerError)
		return
	}

	data, err := getGoogleUserInfo(r.FormValue("code"))
	if err != nil {
		log.Println(err.Error())
		http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
		return
	}

	// Store Id info into Sessoin Cookie
	var userInfo GoogleUserId
	err = json.Unmarshal(data, &userInfo)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	// Get a session. We're ignoring the error resulted from decoding an existing session: Get() always returns a session, even if empty
	session, err := store.Get(r, "session")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	// Set some session values.
	session.Values["id"] = userInfo.ID
	// Save it before we write to the response/return from the handler.
	err = session.Save(r, w)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
}

func getGoogleUserInfo(code string) ([]byte, error) {

	var googleOauthConfig = oauth2.Config{
		RedirectURL:  "http://localhost:3000/auth/google/callback",
		ClientID:     os.Getenv("GOOGLE_CLIENT_ID"),
		ClientSecret: os.Getenv("GOOGLE_SECRET_KEY"),
		Scopes:       []string{"https://www.googleapis.com/auth/userinfo.email"},
		Endpoint:     google.Endpoint,
	}

	token, err := googleOauthConfig.Exchange(context.Background(), code) // converts an authorization code into a token
	if err != nil {
		return nil, fmt.Errorf("Failed to Exchange %s\n", err.Error())
	}

	resp, err := http.Get(oauthGoogleUrlAPI + token.AccessToken) // Request informations of user through authenticate by access token
	if err != nil {
		return nil, fmt.Errorf("Failed to Get UserInfo %s\n", err.Error())
	}

	return ioutil.ReadAll(resp.Body)
}
