// Go is a statically type language
package main

import (

	"encoding/json"
	"net/http"
	"log"
	// "math/rand"
	// "strconv" // string convertor
	"github.com/gorilla/mux"
)

// Book struct (Model)

type Book struct {
	// JSON FIELD
	ID string `json:"id"`
	Isbn string `json:"isbn"`
	Title string `json:"title"`
	// Own struct
	Author *Author `json:"author"`
}

// Author struct

type Author struct {
	Firstname string `json:"firstname"`
	Lastname string `json:"lastname"`
}

// Init book var as a slice book struct

var books []Book

// Get books
// Any route handler function take response in request

func getBooks(w http.ResponseWriter, r *http.Request)  {
	// Set header value
	w.Header().Set("Cntent-Type", "application/json")
	json.NewEncoder(w).Encode(books)
}

// get single book
func getBook(w http.ResponseWriter, r *http.Request)  {
	
}

// create book
func createBook(w http.ResponseWriter, r *http.Request)  {
	
}

// Update Book
func updateBook(w http.ResponseWriter, r *http.Request)  {
	
}


// Delete Books
func deleteBook(w http.ResponseWriter, r *http.Request)  {
	
}

func main()  {
	// Init Router
	r := mux.NewRouter()
	
	// Mock Data 1

	books = append(books, Book{
		ID : "1",
		Isbn: "123456",
		Title: "Python Book",
		Author: &Author{Firstname: "John",Lastname: "Mark"}})
	
	// Mock Data 2
	books = append(books, Book{
		ID : "2",
		Isbn: "65442312",
		Title: "Java Book",
		Author: &Author{
			Firstname: "Mike",
			Lastname: "Martin"}})

	// router handlers / endpoints
	r.HandleFunc("/api/books", getBooks).Methods("GET")
	r.HandleFunc("/api/books/{id}", getBook).Methods("GET")
	r.HandleFunc("/api/books", createBook).Methods("POST")
	r.HandleFunc("/api/books/{id}", updateBook).Methods("PUT")
	r.HandleFunc("/api/books/{id}", deleteBook).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8000", r))

}
