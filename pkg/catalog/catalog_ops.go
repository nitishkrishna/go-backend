package catalog

import (
	"strings"

	"github.com/nitish-krishna/go-backend/pkg/book"
)

func (c *BookCatalog) MigrateBooks() error {
	err := c.DB.AutoMigrate(&book.GoodreadsBook{})
	return err
}

func (c *BookCatalog) BulkInsertBooks(books *[]book.GoodreadsBook) error {
	err := c.DB.Create(books).Error
	return err
}

func (c *BookCatalog) GetBookByIdOp(id int) (*book.GoodreadsBook, error) {
	bookModels := &book.GoodreadsBook{}
	err := c.DB.Where(&book.GoodreadsBook{Id: uint(id)}).First(&bookModels).Error
	return bookModels, err
}

func (c *BookCatalog) SearchBookByNameOp(query string, pagination Pagination) (*Pagination, error) {
	nlQuery := strings.Join(strings.Split(query, " "), "|")
	var books []*book.IndexedGoodreadsBook
	err := c.DB.Scopes(paginate(books, pagination, c.DB)).Where("title_tsv @@ to_tsquery(?)", nlQuery).Find(&books).Error
	// Need to bulk translate these to fetch from GoodreadsBook table instead of indexed table
	pagination.Rows = books
	return &pagination, err
}

func (c *BookCatalog) GetBooksOp(pagination Pagination) (*Pagination, error) {
	var books []*book.GoodreadsBook
	err := c.DB.Scopes(paginate(books, pagination, c.DB)).Find(&books).Error
	pagination.Rows = books
	return &pagination, err
}

func (c *BookCatalog) GetBooksTotalOp() int64 {
	var books []*book.GoodreadsBook
	var count int64
	c.DB.Model(&books).Count(&count)
	return count
}
