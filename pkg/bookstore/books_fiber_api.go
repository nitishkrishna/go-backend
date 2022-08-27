package bookstore

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/nitish-krishna/go-backend/pkg/book"
	"gorm.io/gorm"
	"net/http"
	"strconv"
)

type PostgresBookstore struct {
	DB *gorm.DB
}

func (b *PostgresBookstore) SetupRoutes(app *fiber.App) {
	api := app.Group("/api")
	api.Post("/books", b.CreateBook)
	api.Delete("/books/:id", b.DeleteBook)
	api.Get("/books/:id", b.GetBookByID)
	api.Get("/books", b.GetBooks)
}

func (b *PostgresBookstore) CreateBook(context *fiber.Ctx) error {
	bookObj := book.Book{}

	err := context.BodyParser(&bookObj)

	if err != nil {
		jsonErr := context.Status(http.StatusUnprocessableEntity).JSON(
			&fiber.Map{"message": "request failed"})
		if jsonErr != nil {
			return jsonErr
		}
		return err
	}

	err = b.CreateBookOp(&bookObj)
	if err != nil {
		context.Status(http.StatusBadRequest).JSON(
			&fiber.Map{"message": "could not create bookObj"})
		return err
	}

	context.Status(http.StatusOK).JSON(&fiber.Map{
		"message": "bookObj has been added"})
	return nil
}

func (b *PostgresBookstore) DeleteBook(context *fiber.Ctx) error {
	bookModel := book.Book{}
	id := context.Params("id")
	if id == "" {
		context.Status(http.StatusInternalServerError).JSON(&fiber.Map{
			"message": "id cannot be empty",
		})
		return nil
	}

	bookId, err := strconv.Atoi(id)
	if err != nil {
		return err
	}

	err = b.DeleteBookOp(bookModel, bookId)

	if err != nil {
		context.Status(http.StatusBadRequest).JSON(&fiber.Map{
			"message": "could not delete book",
		})
		return err
	}
	context.Status(http.StatusOK).JSON(&fiber.Map{
		"message": "book delete successfully",
	})
	return nil
}

func (b *PostgresBookstore) GetBooks(context *fiber.Ctx) error {
	bookModels, err := b.FindBooksOp()
	if err != nil {
		context.Status(http.StatusBadRequest).JSON(
			&fiber.Map{"message": "could not get books"})
		return err
	}

	context.Status(http.StatusOK).JSON(&fiber.Map{
		"message": "books fetched successfully",
		"data":    bookModels,
	})
	return nil
}

func (b *PostgresBookstore) GetBookByID(context *fiber.Ctx) error {
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

	bookModel, err := b.GetBookByIdOp(bookId)
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
