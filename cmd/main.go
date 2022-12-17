package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/nitish-krishna/go-backend/pkg/bookstore"
	"github.com/nitish-krishna/go-backend/pkg/catalog"
)

const DatasetFile = "test.csv"

func main() {

	/*
		router := mux.NewRouter()
		book.SetupHandlers(router)
		log.Println("API is running now!")
		_ = http.ListenAndServe(":4000", router)*/

	app := fiber.New()

	app.Use(cors.New(cors.Config{
		AllowOrigins: "http://localhost:5173",
		AllowHeaders: "Origin, Content-Type, Accept",
	}))

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

	var indexingErr error
	indexingFunc := func() error {
		indexingErr = c.IndexFullCatalog()
		return indexingErr
	}
	go indexingFunc()

	if indexingErr != nil {
		log.Fatal(err.Error())
	}

	_ = app.Listen(":4000")
}
