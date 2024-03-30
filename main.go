package main

import (
	"log"
	"net/http"
)

func main() {
	fileServer := http.FileServer(http.Dir("./static"))

	http.Handle("/", fileServer)

	log.Println("starting server on addr", ":3000")
	http.ListenAndServe(":3000", nil)
}
