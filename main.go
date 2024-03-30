package main

import (
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {
	r := chi.NewRouter()
	r.Use(middleware.Logger)

	fileServer := http.FileServer(http.Dir("./static"))

	r.Handle("/", fileServer)

	log.Println("starting server on addr", ":3000")
	http.ListenAndServe(":3000", r)
}
