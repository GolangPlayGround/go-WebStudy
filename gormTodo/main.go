package main

import (
	"goWeb/gormTodo/app"
	"log"
	"net/http"
)

func main() {
	m := app.MakeHandler("./test.db")
	defer m.Close()

	log.Println("Started App")
	err := http.ListenAndServe(":3000", m)
	if err != nil {
		panic(err)
	}
}
