package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/pjw1702/study/web/go/web20/app"
	"github.com/urfave/negroni"
)

func main() {
	m := app.MakeHandler("./test.db")
	defer m.Close()
	n := negroni.Classic()
	n.UseHandler(m)

	log.Println("Started App")
	err := http.ListenAndServe(":3000", n)
	if err != nil {
		// panic(err)
		fmt.Errorf("Failed to start this app: %w", err)
		return
	}
}
