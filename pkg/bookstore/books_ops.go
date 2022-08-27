package bookstore

import (
	"github.com/nitish-krishna/go-backend/pkg/book"
)

func (b *PostgresBookstore) MigrateBooks() error {
	err := b.DB.AutoMigrate(&book.Book{})
	return err
}

func (b *PostgresBookstore) CreateBookOp(bookObj *book.Book) error {
	return b.DB.Create(bookObj).Error
}

func (b *PostgresBookstore) SaveBookOp(bookObj *book.Book) error {
	return b.DB.Save(bookObj).Error
}

func (b *PostgresBookstore) DeleteBookOp(bookObj book.Book, id int) error {
	return b.DB.Delete(bookObj, id).Error
}

func (b *PostgresBookstore) FindBooksOp() (*[]book.Book, error) {
	bookModels := &[]book.Book{}

	err := b.DB.Find(bookModels).Error
	return bookModels, err
}

func (b *PostgresBookstore) GetBookByIdOp(id int) (*book.Book, error) {
	bookModels := &book.Book{}
	err := b.DB.Take(bookModels, id).Error
	return bookModels, err
}
