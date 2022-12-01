package models

import (
	"revel-app-demo/app"
	"time"

	"github.com/revel/revel"
)

type Book struct {
	ID          int       `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Author      string    `json:"author"`
	ReleaseDate time.Time `json:"release_date"`
	TotalPage   int       `json:"total_page"`
	Category    int       `json:"category"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

func (book *Book) Validate(v *revel.Validation) {
	v.Required(book.Title).Message("Title is required")
	v.Required(book.Author).Message("Author is required")
	v.Required(book.ReleaseDate).Message("Release date is required")
}

func (book *Book) GetList() []Book {
	books := []Book{}
	app.DB.Find(&books)

	return books
}

func (book *Book) Create() Book {
	b := Book{
		Title:       book.Title,
		Description: book.Description,
		Author:      book.Author,
		ReleaseDate: book.ReleaseDate,
		TotalPage:   book.TotalPage,
		Category:    book.Category,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}
	app.DB.Create(&book)

	return b
}

func (book *Book) GetByID(id int) Book {
	b := Book{}
	app.DB.Where("id = ?", id).First(&b)

	return b
}

func (book *Book) Update(id int, b Book) {
	b.ID = id
	b.UpdatedAt = time.Now()
	app.DB.Save(&b)
}

func (book *Book) DeleteByID(id int) {
	app.DB.Delete(&Book{}, id)
}
