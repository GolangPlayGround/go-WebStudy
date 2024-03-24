package main

import (
	"goWeb/todo/app"
	"log"
	"net/http"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}
	m := app.MakeHandler()
	defer m.Close()

	log.Println("Started App")
	err = http.ListenAndServe(":3000", m)
	if err != nil {
		panic(err)
	}
}
