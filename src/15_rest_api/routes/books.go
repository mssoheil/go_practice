package routes

import (
	"encoding/json"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type Book struct {
	ID     string  `json:"id"`
	Isbn   string  `json:"isbn"`
	Title  string  `json:"title"`
	Author *Author `json:"author"`
}

type Author struct {
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
}

var books []Book

func setContentType(response http.ResponseWriter) {
	response.Header().Set("Content-Type", "application/json")
}

func getBooks(response http.ResponseWriter, request *http.Request) {
	setContentType(response)

	json.NewEncoder(response).Encode(books)
}

func findBook(id string) *Book {
	for _, item := range books {
		if item.ID == id {
			return (&item)
		}
	}

	return nil

}

func getBook(response http.ResponseWriter, request *http.Request) {
	setContentType(response)

	params := mux.Vars(request)

	book := findBook(params["id"])

	if book != nil {
		json.NewEncoder(response).Encode(book)

	} else {
		response.WriteHeader(http.StatusNotFound)
		json.NewEncoder(response).Encode("Item not found")
	}

}

func createId() string {
	return strconv.Itoa(rand.Intn(10000000))
}

func createSingleBook(request *http.Request) Book {
	var book Book
	json.NewDecoder(request.Body).Decode(&book)

	book.ID = createId()
	books = append(books, book)

	return book
}

func createBook(response http.ResponseWriter, request *http.Request) {
	setContentType(response)

	book := createSingleBook(request)

	response.WriteHeader(http.StatusCreated)
	json.NewEncoder(response).Encode(book)

}

func updateBook(response http.ResponseWriter, request *http.Request) {
	setContentType(response)

	params := mux.Vars(request)

	foundBook := findBook(params["id"])

	if foundBook == nil {
		response.WriteHeader(http.StatusNotFound)
		json.NewEncoder(response).Encode("Item not found")
	}

	for index, item := range books {
		if item.ID == params["id"] {
			books = append(books[:index], books[index+1:]...)
			book := createSingleBook(request)
			json.NewEncoder(response).Encode(book)
			return
		}
		json.NewEncoder(response).Encode(books)
	}

}

func deleteBook(response http.ResponseWriter, request *http.Request) {
	setContentType(response)

	params := mux.Vars(request)

	foundBook := findBook(params["id"])

	if foundBook == nil {
		response.WriteHeader(http.StatusNotFound)
		json.NewEncoder(response).Encode("Item not found")
	}

	for index, item := range books {
		if item.ID == params["id"] {
			books = append(books[:index], books[index+1:]...)
			break
		}
		json.NewEncoder(response).Encode(books)
	}

}

func BookRoute(router *mux.Router) {
	books = append(books, Book{ID: "1", Isbn: "45377", Title: "Book One", Author: &Author{FirstName: "John", LastName: "Doe"}})
	books = append(books, Book{ID: "2", Isbn: "54377", Title: "Book Two", Author: &Author{FirstName: "Jane", LastName: "Smith"}})

	router.HandleFunc("/api/books", getBooks).Methods("GET")
	router.HandleFunc("/api/books/{id}", getBook).Methods("GET")
	router.HandleFunc("/api/books", createBook).Methods("POST")
	router.HandleFunc("/api/books/{id}", updateBook).Methods("PUT")
	router.HandleFunc("/api/books/{id}", deleteBook).Methods("DELETE")

}
