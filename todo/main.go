package main

import (
	"goWeb/todo/app"
	"log"
	"net/http"

	"github.com/joho/godotenv"
	"github.com/urfave/negroni"
)

func main() {
	m := app.MakeHandler("./test.db")
	defer m.Close()
	n := negroni.Classic()
	n.UseHandler(m)
	err := godotenv.Load()

	if err != nil {
		panic(err)
	}

	log.Println("Started App")
	err = http.ListenAndServe(":3000", n)
	if err != nil {
		panic(err)
	}
}
