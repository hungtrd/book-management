package controllers

import (
	"fmt"
	"revel-app-demo/app/models"
	"revel-app-demo/app/routes"

	"github.com/revel/revel"
)

type Books struct {
	// ApiController
	App
}

func (c Books) New() revel.Result {
	return c.Render()
}

func (c Books) Create(book models.Book, cover []byte) revel.Result {
	fmt.Println(book)
	fmt.Println(string(cover))
	book.Cover = cover

	book.Validate(c.Validation)

	if c.Validation.HasErrors() {
		c.Flash.Error("Please fix errors below")
		c.Validation.Keep()
		c.FlashParams()
		return c.Redirect(routes.Books.New())
	}

	_ = book.Create()

	c.Flash.Success("Add book successfully!")
	return c.Redirect(routes.Books.Index())
}

func (c Books) Index() revel.Result {
	b := models.Book{}
	books := b.GetList()

	return c.Render(books)
}

func (c Books) Show(id int) revel.Result {
	// check user login
	if user := c.connected(); user == nil {
		c.Flash.Error("Please login first")
		return c.Redirect(routes.Books.Index())
	}

	// handle request
	b := models.Book{}

	book := b.GetByID(id)
	cover := book.CoverBs64()

	return c.Render(book, cover)
}

func (c Books) Update(id int, book models.Book) revel.Result {
	// check user login
	if user := c.connected(); user == nil {
		c.Flash.Error("Please login first")
		return c.Redirect(routes.Books.Index())
	}

	// handle request
	b := models.Book{}

	book.Validate(c.Validation)

	if c.Validation.HasErrors() {
		c.Flash.Error("Please fix errors below")
		c.Validation.Keep()
		c.FlashParams()
		return c.Redirect(routes.Books.Show(id))
	}

	b.Update(id, book)
	c.Flash.Success("Update book successfully!")
	return c.Redirect(routes.Books.Index())
}

func (c Books) Delete(id int) revel.Result {
	// check user login
	if user := c.connected(); user == nil {
		c.Flash.Error("Please login first")
		return c.Redirect(routes.Books.Index())
	}

	// handle request
	b := models.Book{}

	b.DeleteByID(id)

	c.Flash.Success("Delete book successfully!")
	return c.Redirect(routes.Books.Index())
}

// func BindJsonParams(i io.Reader, s interface{}) error {
// 	bytes, err := ioutil.ReadAll(i)
// 	if err != nil {
// 		return errors.New("can't read request data.")
// 	}

// 	if len(bytes) == 0 {
// 		return errors.New("data is nil")
// 	}

// 	return json.Unmarshal(bytes, s)
// }
