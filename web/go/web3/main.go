package main

import (
	"net/http"

	app "github.com/pjw1702/study/web/go/web3/myapp"
)

func main() {

	http.ListenAndServe(":3000", app.NewHttpHandler())
}
