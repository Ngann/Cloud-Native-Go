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
	//Description `json:"description,omitempty"`
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

// added dictionary structure, the key is the isbn , the values are individual book, in memory data structure
var books = map[string]Book{
	"0345391802": {Title: "The Hitchhiker's Guide to the Galaxy", Author: "Douglas Adams", ISBN: "0345391802"},
	"0000000000": {Title: "Cloud Native Go", Author: "M.-Leander Reimer", ISBN: "0000000000"},
}

//BooksHandleFunc to be used as http.HandleFunc from Book API
func BooksHandleFunc(w http.ResponseWriter, r *http.Request){
	switch method := r.Method; method {
	case http.MethodGet:
		books := AllBooks()
		writeJSON(w,books)
	default:
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Unsupported request method"))
	}
}

//returns slice of Book in the system
func AllBooks() []Book {
	values := make([]Book, len(books))
	idx := 0
	for _, book := range books {
		values[idx] = book
		idx++
	}
	return values
}

//marshal the request and write the request
func writeJSON(w http.ResponseWriter, i interface{}) {
	b, err := json.Marshal(i)
	if err != nil {
		panic(err)
	}
	w.Header().Add("Content-Type", "application/json; charset=utf-8")
	w.Write(b)
}

// postman http://localhost:9090/api/books send a get request should return the Books Json

func BookHandleFunc(w http.ResponseWriter, r *http.Request){
	// take Books slice and marshall to json
	// we get back the byte array of success or err
	b, err :=json.Marshal(books)
	if err != nil {
		panic(err)
	}
	// Now we have Json data we tell the client
	w.Header().Add("Content-Type", "application/json; charset=utf-8")
	// then we write back the response structure
	w.Write(b)
}

// GetBook returns the book for a given ISBN
func GetBook(isbn string) (Book, bool) {
	book, found := books[isbn]
	return book, found
}

// CreateBook creates a new Book if it does not exist
func CreateBook(book Book) (string, bool) {
	_, exists := books[book.ISBN]
	if exists {
		return "", false
	}
	books[book.ISBN] = book
	return book.ISBN, true
}

// UpdateBook updates an existing book
func UpdateBook(isbn string, book Book) bool {
	_, exists := books[isbn]
	if exists {
		books[isbn] = book
	}
	return exists
}

// DeleteBook removes a book from the map by ISBN key
func DeleteBook(isbn string) {
	delete(books, isbn)
}
