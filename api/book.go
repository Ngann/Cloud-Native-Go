package api

import (
	"encoding/json"
	"net/http"
	)

// Book type with Name, Author, ISBN
type Book struct {
	//define the book
}

//toJson to be used for marshalling of Book type
func (b Book) ToJSON() []byte {
	return nil
}

//FromJson to be used for unmarshalling of Book Type
func FromJson(data []byte) Book {

}

//BookHandleFunc to be used as http.HandleFunc fro Book API
func BooksHandleFunc(w http.ResponseWriter, r *http.Request){

}