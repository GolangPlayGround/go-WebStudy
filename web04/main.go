package main

import (
	"goWeb/web04/decoHandler"
	"goWeb/web04/myapp"
	"log"
	"net/http"
	"time"
)

func logger(w http.ResponseWriter, r *http.Request, h http.Handler) {
	start := time.Now()
	log.Println("[LOGGER1] Started")
	h.ServeHTTP(w, r)
	log.Println("[LOGGER1] Completed time:", time.Since(start).Milliseconds())
}

func NewHandler() http.Handler {
	h := myapp.NewHandler()
	h = decoHandler.NewDecoHandler(h, logger)
	return h
}

func main() {
	mux := NewHandler()

	http.ListenAndServe(":3000", mux)
}
