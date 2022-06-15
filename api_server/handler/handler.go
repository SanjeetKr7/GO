package handler

import (
	"fmt"
	"net/http"
)

func Hello(w http.ResponseWriter, r *http.Request) {
	// fmt.Fprintf(w, "Welcome home!")
	fmt.Println("Hello World!")
	w.Write([]byte("Hello World"))
}
