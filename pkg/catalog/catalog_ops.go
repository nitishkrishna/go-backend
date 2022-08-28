package catalog

import "github.com/nitish-krishna/go-backend/pkg/book"

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
	err := c.DB.Take(bookModels, id).Error
	return bookModels, err
}
