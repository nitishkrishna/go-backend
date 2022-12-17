package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"

	"github.com/nitish-krishna/go-backend/pkg/bookstore"
	"github.com/nitish-krishna/go-backend/pkg/catalog"
)

const DatasetFile = "test.csv"

func main() {

	app := fiber.New()

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)

	serverShutdown := make(chan struct{})

	go func() {
		<-c
		fmt.Println("Gracefully shutting down...")
		_ = app.Shutdown()
		serverShutdown <- struct{}{}
	}()

	app.Use(cors.New(cors.Config{
		AllowOrigins: "http://localhost:5173",
		AllowHeaders: "Origin, Content-Type, Accept",
	}))

	bookstore, err := bookstore.InitializeBookstore()
	if err != nil {
		log.Fatal(err.Error())
	}
	bookstore.SetupRoutes(app)

	catalog, err := catalog.InitializeCatalog()
	if err != nil {
		log.Fatal(err.Error())
	}
	catalog.SetupRoutes(app)

	if err := app.Listen(":4000"); err != nil {
		log.Panic(err)
	}

	<-serverShutdown
}
