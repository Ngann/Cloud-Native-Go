package main

import (
	"fmt"
	"net/http"
	"os"
)

func main(){
	http.HandleFunc("/", index)
	http.HandleFunc("api/echo", echo)
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


//implement echo function
// postman http://localhost:9090/api/echo?message=Cloud+Native+Go
// response "Hello Cloud Native Go"
// click header tab, notice the content type = text/plain
func echo(w http.ResponseWriter, r *http.Request){
	//extract the query parameter from the url and grab the first message
	message := r.URL.Query()["message"][0]

	w.Header().Add("Content-Type", "text/plain")
	fmt.Fprintf(w, message)
}