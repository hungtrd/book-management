package models

import (
	"errors"
	"regexp"

	"github.com/revel/revel"
)

type User struct {
	ID       int    `json:"id"`
	MailAddr string `json:"mailaddr"`
	Password string `json:"password"`
	Created  int64  `json:"-"`
	Updated  int64  `json:"-"`
}

func (u *User) Validate() error {
	var v revel.Validation

	v.Match(u.MailAddr, regexp.MustCompile(`([a-zA-Z0-9])+@gmail.com`))
	if v.HasErrors() {
		return errors.New("mail address is validate error")
	}

	v.Check(
		u.Password,
		revel.Required{},
		// revel.MinSize{4},
	)
	if v.HasErrors() {
		return errors.New("password is validate error")
	}

	return nil
}
