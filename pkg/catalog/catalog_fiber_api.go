package catalog

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
	"net/http"
	"strconv"
)

type BookCatalog struct {
	DB *gorm.DB
}

func (c *BookCatalog) SetupRoutes(app *fiber.App) {
	api := app.Group("/api")
	api.Get("/books/:id", c.GetBookByID)
	// TODO: Use for pagination
	// api.Get("/books", c.GetBooks)
}

func (c *BookCatalog) GetBookByID(context *fiber.Ctx) error {
	id := context.Params("id")
	if id == "" {
		context.Status(http.StatusInternalServerError).JSON(&fiber.Map{
			"message": "id cannot be empty",
		})
		return nil
	}

	fmt.Println("the ID is", id)

	bookId, err := strconv.Atoi(id)
	if err != nil {
		return err
	}

	bookModel, err := c.GetBookByIdOp(bookId)
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
