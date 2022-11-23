package models

import (
	"errors"
	"regexp"

	"github.com/revel/revel"
)

type Book struct {
	ID       int    `json:"id"`
	MailAddr string `json:"mailaddr"`
	Password string `json:"password"`
	Created  int64  `json:"-"`
	Updated  int64  `json:"-"`
}

func (b *Book) Validate() error {
	var v revel.Validation

	v.Match(b.MailAddr, regexp.MustCompile(`([a-zA-Z0-9])+@gmail.com`))
	if v.HasErrors() {
		return errors.New("mail address is validate error")
	}

	v.Check(
		b.Password,
		revel.Required{},
		revel.MinSize{4},
	)
	if v.HasErrors() {
		return errors.New("password is validate error")
	}

	return nil
}
