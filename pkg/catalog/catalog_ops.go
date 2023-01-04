package catalog

import (
	"strings"

	"github.com/nitish-krishna/go-backend/pkg/book"
)

func (c *BookCatalog) MigrateBooks() error {
	err := c.DB.AutoMigrate(&book.GoodreadsBook{})
	if err != nil {
		return err
	}
	c.DB.Exec(TSVIndexQuery)
	return nil
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
	var books []*book.GoodreadsBookSearchResult
	// Gorm allows you to query only for the fields you need
	// https://gorm.io/docs/advanced_query.html
	err := c.DB.Model(&book.GoodreadsBook{}).Scopes(paginate(book.GoodreadsBook{}, pagination, c.DB)).Where("title_tsv @@ to_tsquery(?)", nlQuery).Select("Title", "Authors", "AverageRating", "NumPages").Find(&books).Error
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
