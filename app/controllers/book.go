package controllers

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"revel-app-demo/app/models"

	"github.com/revel/revel"
)

type Books struct {
	ApiController
}

type BookCreateReq struct {
	Title       string `json:"title"`
	Description string `json:"description"`
}

func (c Books) Create() revel.Result {
	b := &models.Book{}
	c.Params.BindJSON(b)

	fmt.Println(b)

	b.Validate(c.Validation)
	if c.Validation.HasErrors() {
		return c.RenderJSON(&ErrorResponse{ERR_VALIDATE, "ErrorMessage(ERR_VALIDATE)"})
	}

	book := b.Create()

	return c.RenderJSON(&Response{OK, book})
}

func (c Books) Index() revel.Result {
	// var books []models.Book
	b := models.Book{}
	books := b.GetList()

	return c.RenderJSON(&Response{OK, books})
}

func BindJsonParams(i io.Reader, s interface{}) error {
	bytes, err := ioutil.ReadAll(i)
	if err != nil {
		return errors.New("can't read request data.")
	}

	if len(bytes) == 0 {
		return errors.New("data is nil")
	}

	return json.Unmarshal(bytes, s)
}
