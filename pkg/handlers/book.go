package handlers

import (
	"github.com/gorilla/mux"
	"github.com/nitish-krishna/go-backend/pkg/book"
	"net/http"
)

func SetupBookHandlers(router *mux.Router) {
	router.HandleFunc("/books", book.GetAllBooks).Methods(http.MethodGet)
	router.HandleFunc("/books/{id}", book.GetBook).Methods(http.MethodGet)
	router.HandleFunc("/books", book.AddBook).Methods(http.MethodPost)
	router.HandleFunc("/books/{id}", book.UpdateBook).Methods(http.MethodPut)
	router.HandleFunc("/books/{id}", book.DeleteBook).Methods(http.MethodDelete)
}
