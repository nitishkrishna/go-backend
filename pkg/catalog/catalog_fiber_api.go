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

const catalogEnvFilePath = ".catalogenv"

type BookCatalog struct {
	DB *gorm.DB
}

func InitializeCatalog() (*BookCatalog, error) {
	dbConfig, err := db.ParsePostgresConfig(catalogEnvFilePath)
	if err != nil {
		return nil, fmt.Errorf("could not get db config: %w", err)
	}

	dbObj, err := db.NewPostgresConnection(dbConfig)
	if err != nil {
		return nil, fmt.Errorf("could not load the database: %w", err)
	}

	c := BookCatalog{DB: dbObj}
	err = c.MigrateBooks()
	if err != nil {
		return nil, fmt.Errorf("could not migrate db: %w", err)
	}

	_ = dataset.ReadFromCSV("./pkg/dataset/books.csv")
	// Save all the records at once in the database
	/*err = c.BulkInsertBooks(books)
	if err != nil {
		return nil, fmt.Errorf("could not bulk insert into db: %w", err)
	}*/

	return &c, nil
}

func (c *BookCatalog) SetupRoutes(app *fiber.App) {
	api := app.Group("/catalog")
	api.Get("/books/:id", c.GetBookByID)
	api.Get("/books", c.GetBooks)
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

func (c *BookCatalog) GetBooks(context *fiber.Ctx) error {
	page, _ := strconv.Atoi(context.Params("page"))
	if page <= 0 {
		page = 1
	}
	limit, _ := strconv.Atoi(context.Params("limit"))
	if limit <= 0 || limit > 1000 {
		limit = 10
	}
	sort := context.Params("sort")
	var sortWithDirection string
	if sort != "" {
		direction := context.Params("sortDesc")
		if direction != "" {
			if direction == "true" {
				sortWithDirection = sort + " desc"
			} else if direction == "false" {
				sortWithDirection = sort + " asc"
			}
		}
	}

	pagination := newPagination(limit, page, sortWithDirection)
	paginatedResult, err := c.GetBooksOp(*pagination)

	if err != nil {
		jsonErr := context.Status(http.StatusBadRequest).JSON(
			&fiber.Map{"message": "could not get the books"})
		if jsonErr != nil {
			return jsonErr
		}
		return err
	}
	context.Status(http.StatusOK).JSON(&fiber.Map{
		"message": "books fetched successfully",
		"data":    paginatedResult.Rows,
	})
	if err != nil {
		return err
	}
	return nil
}
