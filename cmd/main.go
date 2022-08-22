package main

import (
	"github.com/gorilla/mux"
	"github.com/nitish-krishna/go-backend/pkg/db"
	"github.com/nitish-krishna/go-backend/pkg/handlers"
	"log"
	"net/http"
)

func main() {

	dbConfig, err := db.ParsePostgresConfig(".env")
	if err != nil {
		log.Fatal("could not get db config")
	}

	_, err = db.NewPostgresConnection(dbConfig)
	if err != nil {
		log.Fatal("could not load the database")
	}
	router := mux.NewRouter()
	handlers.SetupBookHandlers(router)
	log.Println("API is running now!")
	_ = http.ListenAndServe(":4000", router)
}
