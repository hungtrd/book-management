package models

import (
	"errors"
	"fmt"
	"revel-app-demo/app"
	"time"

	"github.com/revel/revel"
)

type Book struct {
	ID          int       `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

func (b *Book) Validate(v *revel.Validation) error {
	v.Check(
		b.Title,
		revel.Required{},
	)
	fmt.Println(v.Errors)

	if v.HasErrors() {
		return errors.New("title is validate error")
	}

	return nil
}

func (b *Book) GetList() []Book {
	books := []Book{}
	app.DB.Find(&books)

	return books
}

func (b *Book) Create() Book {
	book := Book{
		Title:       b.Title,
		Description: b.Description,
	}
	app.DB.Create(&book)

	return book
}

func (b *Book) GetByID(id int) Book {
	book := Book{}
	app.DB.Where("id = ?", id).First(&book)

	return book
}
