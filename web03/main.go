package main

import (
	"goWeb/web03/myapp"
	"net/http"
)

func main() {
	http.ListenAndServe(":8090", myapp.NewHandler())
}
