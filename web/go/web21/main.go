package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/pjw1702/study/web/go/web21/app"
)

func main() {
	m := app.MakeHandler("./test.db")
	defer m.Close()
	// n := negroni.Classic()
	// n.UseHandler(m)

	log.Println("Started App")
	err := http.ListenAndServe(":3000", m)
	if err != nil {
		// panic(err)
		fmt.Println("Failed to start this app: %w", err)
		return
	}
}
