package main

import (
	"api_server/handler"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter().StrictSlash(true)
	r.HandleFunc("/", handler.Hello)
	log.Fatal(http.ListenAndServe(":8080", r))
}
