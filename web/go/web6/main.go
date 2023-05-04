package main

import (
	"net/http"

	"github.com/pjw1702/study/web/go/web6/myapp"
)

func main() {
	http.ListenAndServe(":3000", myapp.NewHandler())
}
