package catalog

import "github.com/nitish-krishna/go-backend/pkg/book"

func (c *BookCatalog) GetBookByIdOp(id int) (*book.GoodreadsBook, error) {
	bookModels := &book.GoodreadsBook{}
	err := c.DB.Take(bookModels, id).Error
	return bookModels, err
}
