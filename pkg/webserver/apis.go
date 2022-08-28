package webserver

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/nitish-krishna/go-backend/pkg/book"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"strconv"
)

func GetAllBooks(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(book.Books)
}

func GetBook(w http.ResponseWriter, r *http.Request) {
	// Read dynamic id parameter
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])

	// Iterate over all the mock books
	for _, bookObj := range book.Books {
		if bookObj.Id == id {
			// If ids are equal send bookObj as response
			w.Header().Add("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(bookObj)
			break
		}
	}
}

// TODO: Creates duplicates, fix

func AddBook(w http.ResponseWriter, r *http.Request) {
	// Read to request body
	defer r.Body.Close()
	body, err := ioutil.ReadAll(r.Body)

	if err != nil {
		log.Fatalln(err)
	}

	var bookObj book.Book
	json.Unmarshal(body, &bookObj)

	// Append to the Book mocks
	bookObj.Id = rand.Intn(100)
	book.Books = append(book.Books, bookObj)

	// Send a 201 created response
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode("Created")
}

func DeleteBook(w http.ResponseWriter, r *http.Request) {
	// Read the dynamic id parameter
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])

	// Iterate over all the mock Books
	for index, bookObj := range book.Books {
		if bookObj.Id == id {
			// Delete bookObj and send response if the bookObj Id matches dynamic Id
			book.Books = append(book.Books[:index], book.Books[index+1:]...)

			w.Header().Add("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode("Deleted")
			break
		}
	}
}

func UpdateBook(w http.ResponseWriter, r *http.Request) {
	// Read dynamic id parameter
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])

	// Read request body
	defer r.Body.Close()
	body, err := ioutil.ReadAll(r.Body)

	if err != nil {
		log.Fatalln(err)
	}

	var updatedBook book.Book
	json.Unmarshal(body, &updatedBook)

	// Iterate over all the mock Books
	for index, bookObj := range book.Books {
		if bookObj.Id == id {
			// Update and send response when bookObj Id matches dynamic Id
			bookObj.Title = updatedBook.Title
			bookObj.Author = updatedBook.Author
			bookObj.Desc = updatedBook.Desc
			bookObj.ISBN = updatedBook.ISBN

			book.Books[index] = bookObj

			w.Header().Add("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode("Updated")
			break
		}
	}
}
