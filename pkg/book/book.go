package book

type Book struct {
	Id     int    `gorm:"primary key;autoIncrement" json:"id"`
	Title  string `json:"title"`
	Author string `json:"author"`
	Desc   string `json:"desc"`
	ISBN   string `json:"isbn"`
}

var Books = []Book{
	{
		Id:     1,
		Title:  "Golang for beginners",
		Author: "Gopher",
		Desc:   "A beginner book for Golang",
		ISBN:   "abc",
	},
}
