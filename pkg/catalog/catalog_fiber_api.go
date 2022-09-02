package catalog

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/nitish-krishna/go-backend/pkg/dataset"
	"github.com/nitish-krishna/go-backend/pkg/db"
	"gorm.io/gorm"
	"net/http"
	"strconv"
)

const catalogEnvFilePath = ".env"

type BookCatalog struct {
	DB *gorm.DB
}

func InitializeCatalog() (*BookCatalog, error) {
	dbConfig, err := db.ParsePostgresConfig(catalogEnvFilePath)
	if err != nil {
		return nil, fmt.Errorf("could not get db config: %w", err)
	}

	dbConfig.DBName = "catalog"

	dbObj, err := db.NewPostgresConnection(dbConfig)
	if err != nil {
		return nil, fmt.Errorf("could not load the database: %w", err)
	}

	c := BookCatalog{DB: dbObj}
	err = c.MigrateBooks()
	if err != nil {
		return nil, fmt.Errorf("could not migrate db: %w", err)
	}

	books := dataset.ReadFromCSV("./pkg/dataset/books.csv")
	// Save all the records at once in the database
	c.DB = c.DB.Session(&gorm.Session{CreateBatchSize: 1000})
	err = c.BulkInsertBooks(books)
	if err != nil {
		return nil, fmt.Errorf("could not bulk insert into db: %w", err)
	}

	return &c, nil
}

func (c *BookCatalog) SetupRoutes(app *fiber.App) {
	api := app.Group("/catalog")
	api.Get("/books/:id", c.GetBookByID)
	// TODO: Use for pagination
	// api.Get("/books", c.GetBooks)
}

func (c *BookCatalog) GetBookByID(context *fiber.Ctx) error {
	bookID := context.Params("id")
	if bookID == "" {
		context.Status(http.StatusInternalServerError).JSON(&fiber.Map{
			"message": "book id cannot be empty",
		})
		return nil
	}

	fmt.Println("the Book ID is", bookID)

	bookIdInt, err := strconv.Atoi(bookID)
	if err != nil {
		return err
	}

	bookModel, err := c.GetBookByIdOp(bookIdInt)
	if err != nil {
		jsonErr := context.Status(http.StatusBadRequest).JSON(
			&fiber.Map{"message": "could not get the book"})
		if jsonErr != nil {
			return jsonErr
		}
		return err
	}
	err = context.Status(http.StatusOK).JSON(&fiber.Map{
		"message": "book id fetched successfully",
		"data":    bookModel,
	})
	if err != nil {
		return err
	}
	return nil
}
