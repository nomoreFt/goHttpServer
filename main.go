package main

import (
	"net/http"

	"github.com/nomoreFt/goHttpServer/myapp/myapp"
)

func main() {

	http.ListenAndServe(":3000", myapp.NewHttpHandler())
}
