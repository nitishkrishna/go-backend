package catalog

import (
	"github.com/nitish-krishna/go-backend/pkg/book"
)

func (c *BookCatalog) MigrateIndexedBooks() error {
	err := c.DB.AutoMigrate(&book.IndexedGoodreadsBook{})
	return err
}
