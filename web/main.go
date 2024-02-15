package main

import (
	"goWeb/myapp"
	"net/http"
)

func main() {
	http.ListenAndServe(":3000", myapp.NewHttpHandler())
}
