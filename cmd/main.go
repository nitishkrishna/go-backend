package main

import (
	"github.com/gofiber/fiber/v2"

	"github.com/nitish-krishna/go-backend/pkg/bookstore"
	"github.com/nitish-krishna/go-backend/pkg/db"
	"log"
)

const DatasetFile = "test.csv"

func main() {

	dbConfig, err := db.ParsePostgresConfig(".env")
	if err != nil {
		log.Fatal("could not get db config")
	}

	dbObj, err := db.NewPostgresConnection(dbConfig)
	if err != nil {
		log.Fatal("could not load the database")
	}

	/*
		router := mux.NewRouter()
		book.SetupHandlers(router)
		log.Println("API is running now!")
		_ = http.ListenAndServe(":4000", router)*/

	app := fiber.New()
	b := bookstore.PostgresBookstore{DB: dbObj}
	err = b.MigrateBooks()
	if err != nil {
		log.Fatal("could not migrate db")
	}
	b.SetupRoutes(app)
	_ = app.Listen(":4000")
}
