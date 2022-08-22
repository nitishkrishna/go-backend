package book

import (
	"github.com/gorilla/mux"
	"net/http"
)

func SetupHandlers(router *mux.Router) {
	router.HandleFunc("/books", GetAllBooks).Methods(http.MethodGet)
	router.HandleFunc("/books/{id}", GetBook).Methods(http.MethodGet)
	router.HandleFunc("/books", AddBook).Methods(http.MethodPost)
	router.HandleFunc("/books/{id}", UpdateBook).Methods(http.MethodPut)
	router.HandleFunc("/books/{id}", DeleteBook).Methods(http.MethodDelete)
}
