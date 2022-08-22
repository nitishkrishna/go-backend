package main

import (
	"github.com/gorilla/mux"
	"github.com/nitish-krishna/go-backend/pkg/book"
	"log"
	"net/http"
)

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/books", book.GetAllBooks).Methods(http.MethodGet)
	router.HandleFunc("/books/{id}", book.GetBook).Methods(http.MethodGet)
	router.HandleFunc("/books", book.AddBook).Methods(http.MethodPost)
	router.HandleFunc("/books/{id}", book.UpdateBook).Methods(http.MethodPut)
	router.HandleFunc("/books/{id}", book.DeleteBook).Methods(http.MethodDelete)
	log.Println("API is running now!")
	_ = http.ListenAndServe(":4000", router)
}
