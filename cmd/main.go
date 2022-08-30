package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/nitish-krishna/go-backend/pkg/bookstore"
	"github.com/nitish-krishna/go-backend/pkg/catalog"
	"log"
)

const DatasetFile = "test.csv"

func main() {

	/*
		router := mux.NewRouter()
		book.SetupHandlers(router)
		log.Println("API is running now!")
		_ = http.ListenAndServe(":4000", router)*/

	app := fiber.New()

	b, err := bookstore.InitializeBookstore()
	if err != nil {
		log.Fatal(err.Error())
	}
	b.SetupRoutes(app)

	c, err := catalog.InitializeCatalog()
	if err != nil {
		log.Fatal(err.Error())
	}
	c.SetupRoutes(app)

	_ = app.Listen(":4000")
}
