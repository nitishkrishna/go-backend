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
	err := c.DB.Where(&book.GoodreadsBook{BookId: id}).First(&bookModels).Error
	return bookModels, err
}

func (c *BookCatalog) GetBooksOp(pagination Pagination) (*Pagination, error) {
	var books []*book.GoodreadsBook
	err := c.DB.Scopes(paginate(books, pagination, c.DB)).Find(&books).Error
	pagination.Rows = books
	return &pagination, err
}
