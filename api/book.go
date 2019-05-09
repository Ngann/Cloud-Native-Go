package api

import (
	"encoding/json"
	_"encoding/json"
	"net/http"
	)

// Book type with Name, Author, ISBN
type Book struct {
	Title string `json:"title"`
	Author string `json:"author"`
	ISBN string `json:"isbn"`
}

// ToJSON to be used for marshalling of Book type
func (b Book) ToJSON() []byte {
	ToJSON, err := json.Marshal(b)
	if err != nil {
		panic(err)
	}
	return ToJSON
}

// FromJSON to be used for unmarshalling of Book type
func FromJSON(data []byte) Book {
	book := Book{}
	err := json.Unmarshal(data, &book)
	if err != nil {
		panic(err)
	}
	return book
}

var Books = []Book{
	 {Title: "The Hitchhiker's Guide to the Galaxy", Author: "Douglas Adams", ISBN: "0345391802"},
	 {Title: "Cloud Native Go", Author: "M.-Leander Reimer", ISBN: "0000000000"},
}

//BookHandleFunc to be used as http.HandleFunc fro Book API
func BooksHandleFunc(w http.ResponseWriter, r *http.Request){
	// take Books slice and marshall to json
	// we get back the byte array of success or err
	b, err :=json.Marshal(Books)
	if err != nil {
		panic(err)
	}
	// Now we have Json data we tell the client
	w.Header().Add("Content-Type", "application/json; charset=utf-8")
	// then we write back the response structure
	w.Write(b)
}

// postman http://localhost:9090/api/books send a get request should return the Books Json