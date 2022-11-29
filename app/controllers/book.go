package controllers

import (
	"fmt"
	"revel-app-demo/app/models"
	"revel-app-demo/app/routes"

	"github.com/revel/revel"
)

type Books struct {
	ApiController
	// App
}

// type BookCreateReq struct {
// 	Title       string `json:"title"`
// 	Description string `json:"description"`
// }

func (c Books) New() revel.Result {
	return c.Render()
}

func (c Books) Create(book models.Book) revel.Result {
	fmt.Println(book)

	book.Validate(c.Validation)

	if c.Validation.HasErrors() {
		c.Flash.Error("Please fix errors below")
		c.Validation.Keep()
		c.FlashParams()
		return c.Redirect(routes.Books.New())
	}

	_ = book.Create()

	// return c.RenderJSON(&Response{OK, book})
	c.Flash.Success("Create successfully!")
	return c.Redirect(routes.Books.Index())
}

func (c Books) Index() revel.Result {
	// var books []models.Book
	b := models.Book{}
	books := b.GetList()

	return c.RenderJSON(&Response{OK, books})
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
