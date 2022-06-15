package client_request

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func Hello() {
	resp, err := http.Get("http://localhost:8080/")
	if err != nil {
		fmt.Println(err)
	}
	body, readErr := ioutil.ReadAll(resp.Body)
	if readErr != nil {
		log.Fatal(readErr)
	}
	fmt.Println(string(body))
	defer resp.Body.Close()
}
