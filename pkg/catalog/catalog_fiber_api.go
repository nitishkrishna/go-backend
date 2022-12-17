package catalog

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/pkg/errors"
	"gorm.io/gorm"

	"github.com/nitish-krishna/go-backend/pkg/book"
	"github.com/nitish-krishna/go-backend/pkg/dataset"
	"github.com/nitish-krishna/go-backend/pkg/db"
)

const catalogEnvFilePath = ".env"
const TSVIndexQuery = `CREATE INDEX IF NOT EXISTS tsv_title_idx ON catalog USING GIN (titletsv)`

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
		return nil, fmt.Errorf("could not migrate db for books: %w", err)
	}

	var count int64
	c.DB.Model(&book.GoodreadsBook{}).Count(&count)

	if count == 0 {
		books := dataset.ReadFromCSV("./pkg/dataset/books.csv")
		// Save all the records at once in the database
		c.DB = c.DB.Session(&gorm.Session{CreateBatchSize: 1000})
		err = c.BulkInsertBooks(books)
		if err != nil {
			return nil, fmt.Errorf("could not bulk insert into db: %w", err)
		}
	}

	return &c, nil
}

func (c *BookCatalog) IndexFullCatalog() error {
	err := c.MigrateIndexedBooks()
	if err != nil {
		return errors.Wrapf(err, "could not migrate db for indexed books")
	}

	var count, indexedCount int64
	c.DB.Model(&book.GoodreadsBook{}).Count(&count)
	c.DB.Model(&book.IndexedGoodreadsBook{}).Count(&indexedCount)
	if count != indexedCount {
		c.DB.Migrator().DropTable(&book.IndexedGoodreadsBook{})
		err := c.MigrateIndexedBooks()
		c.DB.Exec(TSVIndexQuery)
		if err != nil {
			return errors.Wrapf(err, "could not migrate db for indexed books")
		}
		var bookRecords []*book.GoodreadsBook
		err = c.DB.Find(&bookRecords).Error
		if err != nil {
			return errors.Wrapf(err, "could not read all books from DB")
		}

		for _, bookRecord := range bookRecords {
			indexedBookRecord := book.IndexedGoodreadsBook{
				GoodreadsBook: *bookRecord,
			}
			res := c.DB.Create(&indexedBookRecord)
			if res.Error != nil {
				return errors.Wrapf(res.Error, "could not index record for book: %v", bookRecord.Title)
			}
			fmt.Println("Still indexing...")
		}
	}

	fmt.Println("Finished indexing books")
	testIndexedBook := &book.IndexedGoodreadsBook{}
	err = c.DB.Last(&testIndexedBook).Error
	if err != nil {
		return errors.Wrapf(err, "could not find last indexed book")
	}

	lexemes := testIndexedBook.TitleTSV
	fmt.Printf("lexemes of last book: %s", lexemes)

	return nil
}

func (c *BookCatalog) SetupRoutes(app *fiber.App) {
	api := app.Group("/catalog")
	api.Get("/books/:id", c.GetBookByID)
	api.Get("/books", c.GetBooks)
	api.Get("/total", c.GetBookTotal)
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
	page, _ := strconv.Atoi(context.Query("page"))
	if page <= 0 {
		page = 1
	}
	limit, _ := strconv.Atoi(context.Query("limit"))
	if limit <= 0 || limit > 1000 {
		limit = 10
	}
	sort := context.Query("sort")
	var sortWithDirection string
	if sort != "" {
		direction := context.Query("sortDesc")
		if direction != "" {
			if direction == "true" {
				sortWithDirection = sort + " desc"
			} else if direction == "false" {
				sortWithDirection = sort + " asc"
			}
		}
	}

	pagination := newPagination(limit, page, sortWithDirection)

	fmt.Printf("Getting books, limit %d, page %d, sort %s\n", limit, page, sortWithDirection)

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

func (c *BookCatalog) GetBookTotal(context *fiber.Ctx) error {

	count := c.GetBooksTotalOp()

	fmt.Println("the Book count total is: ", count)

	err := context.Status(http.StatusOK).JSON(&fiber.Map{
		"message": "book count fetched successfully",
		"data":    count,
	})
	if err != nil {
		return err
	}
	return nil
}
