package controllers

import (
	"encoding/json"
	"errors"
	"io"
	"io/ioutil"
	"revel-app-demo/app/models"

	"github.com/revel/revel"
)

type ApiBook struct {
	ApiController
}

func (c *ApiBook) Create() revel.Result {
	book := &models.Book{}
	c.Params.BindJSON(book)

	book.Validate()
	if c.Validation.HasErrors() {
		return c.RenderJSON(&ErrorResponse{ERR_VALIDATE, ErrorMessage(ERR_VALIDATE)})
	}

	err := c.Txn.Insert(book)
	if err != nil {
		panic(err)
	}

	return c.RenderJSON(&Response{OK, book})
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
