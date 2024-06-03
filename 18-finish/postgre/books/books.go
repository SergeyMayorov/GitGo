package books

import "fmt"

type bookType struct {
	Title    string
	AuthorID int
	Year     int
	ISBN     string
}

func NewBook(title string, authorID int, year int, isbn string) *bookType {
	return &bookType{
		Title:    title,
		AuthorID: authorID,
		Year:     year,
		ISBN:     isbn,
	}
}

type BookCRUD interface {
	NewBook(title string, authorID int, year int, isbn string) *bookType
	Read(*bookType) string
	Update(b *bookType) bool
}

func Read(b *bookType) string {
	r := fmt.Sprintf("%s ", b.Title)
	return r
}
