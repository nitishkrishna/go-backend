package docker_test

import (
	"github.com/gofiber/fiber/v2"
	"github.com/nitish-krishna/go-backend/pkg/book"
	"github.com/nitish-krishna/go-backend/pkg/bookstore"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Repository", func() {
	var testBookstore *bookstore.PostgresBookstore
	var testApp *fiber.App
	BeforeEach(func() {
		testBookstore = &bookstore.PostgresBookstore{DB: Db}
		err := testBookstore.MigrateBooks() // auto create tables
		Ω(err).To(Succeed())

		testApp = fiber.New()
		testBookstore.SetupRoutes(testApp)

		sampleData := &book.Book{
			Id:     1,
			Title:  "test",
			Author: "fake",
			Desc:   "book for testing",
			ISBN:   "abcd",
		}
		// create sample data
		err = Db.Create(sampleData).Error
		Ω(err).To(Succeed())
	})
	Context("Get", func() {
		It("Found Book", func() {
			book, err := testBookstore.GetBookByIdOp(1)

			Ω(err).To(Succeed())
			Ω(book.Title).To(Equal("test"))
			Ω(book.Author).To(Equal("fake"))
		})
		It("Not Found", func() {
			_, err := testBookstore.GetBookByIdOp(999)
			Ω(err).To(HaveOccurred())
		})
	})

	It("ListAll", func() {
		l, err := testBookstore.FindBooksOp()
		Ω(err).To(Succeed())
		Ω(*l).To(HaveLen(1))
	})

	Context("Save", func() {

		It("Create", func() {
			testBook := &book.Book{
				Id:     2,
				Title:  "newtest",
				Author: "newfake",
				Desc:   "new book for testing",
				ISBN:   "wxyz",
			}
			err := testBookstore.CreateBookOp(testBook)
			Ω(err).To(Succeed())
			Ω(testBook.Id).To(BeEquivalentTo(2))
		})
		It("Update", func() {
			foundBook, err := testBookstore.GetBookByIdOp(1)
			Ω(err).To(Succeed())

			foundBook.Title = "foo"
			err = testBookstore.SaveBookOp(foundBook)
			Ω(err).To(Succeed())

			updatedBook, err := testBookstore.GetBookByIdOp(1)
			Ω(err).To(Succeed())
			Ω(updatedBook.Title).To(BeEquivalentTo("foo"))
		})
	})
	It("Delete", func() {
		foundBook, err := testBookstore.GetBookByIdOp(1)
		Ω(err).To(Succeed())
		err = testBookstore.DeleteBookOp(*foundBook, 1)
		Ω(err).To(Succeed())
		_, err = testBookstore.GetBookByIdOp(1)
		Ω(err).To(HaveOccurred())
	})

})
