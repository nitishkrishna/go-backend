package book

import (
	"gorm.io/gorm"
)

type GoodreadsBook struct {
	gorm.Model

	Id               uint    `gorm:"primarykey" json:"book_id" csv:"book_id"`
	Title            string  `json:"title" csv:"title"`
	Authors          string  `json:"authors" csv:"authors"`
	AverageRating    float32 `json:"average_rating" csv:"average_rating"`
	ISBN             string  `json:"isbn" csv:"isbn"`
	ISBN13           string  `json:"isbn13" csv:"isbn13"`
	LanguageCode     string  `json:"language_code" csv:"language_code"`
	NumPages         string  `json:"num_pages" csv:"num_pages"`
	RatingsCount     uint    `json:"ratings_count" csv:"ratings_count"`
	TextReviewsCount uint    `json:"text_reviews_count" csv:"text_reviews_count"`
	PublicationDate  string  `json:"publication_date" csv:"publication_date"`
	Publisher        string  `json:"publisher" csv:"publisher"`
}

type IndexedGoodreadsBook struct {
	GoodreadsBook
	TitleTSV  string `gorm:"->;type:tsvector GENERATED ALWAYS AS (to_tsvector(title)) STORED;default:(-)"`
	AuthorTSV string `gorm:"->;type:tsvector GENERATED ALWAYS AS (to_tsvector(authors)) STORED;default:(-)"`
}
