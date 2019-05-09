package main

import (
	"fmt"
	"net/http"
	"os"
)

func main(){
	http.HandleFunc("/", index)
	http.ListenAndServe(port(), nil)

}

func index(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, "Hello Cloud Native Go")
}

//extract port parameter to make it configurable
// 8080 is still default

// In terminal $ export PORT=9090
//go run microservice.go
// go to postman http://localhost:9090/ , send a get request
// response should be " Hello Cloud Native"
func port() string {
	port := os.Getenv("PORT")
	if len(port) == 0 {
		port = "8080"
	}
	return ":" + port
}

