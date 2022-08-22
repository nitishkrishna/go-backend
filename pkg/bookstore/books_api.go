package bookstore

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/nitish-krishna/go-backend/pkg/book"
	"gorm.io/gorm"
	"net/http"
)

type Bookstore struct {
	DB *gorm.DB
}

func (b *Bookstore) SetupRoutes(app *fiber.App) {
	api := app.Group("/api")
	api.Post("/books", b.CreateBook)
	api.Delete("/books/:id", b.DeleteBook)
	api.Get("/books/:id", b.GetBookByID)
	api.Get("/books", b.GetBooks)
}

func (b *Bookstore) MigrateBooks() error {
	err := b.DB.AutoMigrate(&book.Book{})
	return err
}

func (b *Bookstore) CreateBook(context *fiber.Ctx) error {
	book := book.Book{}

	err := context.BodyParser(&book)

	if err != nil {
		context.Status(http.StatusUnprocessableEntity).JSON(
			&fiber.Map{"message": "request failed"})
		return err
	}

	err = b.DB.Create(&book).Error
	if err != nil {
		context.Status(http.StatusBadRequest).JSON(
			&fiber.Map{"message": "could not create book"})
		return err
	}

	context.Status(http.StatusOK).JSON(&fiber.Map{
		"message": "book has been added"})
	return nil
}

func (b *Bookstore) DeleteBook(context *fiber.Ctx) error {
	bookModel := book.Book{}
	id := context.Params("id")
	if id == "" {
		context.Status(http.StatusInternalServerError).JSON(&fiber.Map{
			"message": "id cannot be empty",
		})
		return nil
	}

	err := b.DB.Delete(bookModel, id)

	if err.Error != nil {
		context.Status(http.StatusBadRequest).JSON(&fiber.Map{
			"message": "could not delete book",
		})
		return err.Error
	}
	context.Status(http.StatusOK).JSON(&fiber.Map{
		"message": "book delete successfully",
	})
	return nil
}

func (b *Bookstore) GetBooks(context *fiber.Ctx) error {
	bookModels := &[]book.Book{}

	err := b.DB.Find(bookModels).Error
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

func (b *Bookstore) GetBookByID(context *fiber.Ctx) error {

	id := context.Params("id")
	bookModel := &book.Book{}
	if id == "" {
		context.Status(http.StatusInternalServerError).JSON(&fiber.Map{
			"message": "id cannot be empty",
		})
		return nil
	}

	fmt.Println("the ID is", id)

	err := b.DB.Where("id = ?", id).First(bookModel).Error
	if err != nil {
		context.Status(http.StatusBadRequest).JSON(
			&fiber.Map{"message": "could not get the book"})
		return err
	}
	context.Status(http.StatusOK).JSON(&fiber.Map{
		"message": "book id fetched successfully",
		"data":    bookModel,
	})
	return nil
}
