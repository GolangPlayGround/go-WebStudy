package main

import (
	"goWeb/todo/app"
	"log"
	"net/http"

	"github.com/urfave/negroni"
)

func main() {
	m := app.MakeHandler("./test.db")
	defer m.Close()
	n := negroni.Classic()
	n.UseHandler(m)

	log.Println("====Started App===")
	err := http.ListenAndServe(":8090", n)
	if err != nil {
		panic(err)
	}
}
