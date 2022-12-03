package models

import (
	"regexp"
	"revel-app-demo/app"
	"time"

	"github.com/revel/revel"
)

type User struct {
	ID        int       `json:"id"`
	Username  string    `json:"username"`
	Fullname  string    `json:"fullname"`
	Email     string    `json:"email"`
	Password  string    `json:"password"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

var usernameRegex = regexp.MustCompile(`^\w*$`)
var emailRegex = regexp.MustCompile(`([a-zA-Z0-9])+@gmail.com`)

func (user *User) Validate(v *revel.Validation) {
	v.Required(user.Username).Message("Username is required")
	v.Required(user.Fullname).Message("Fullname is required")
	v.Required(user.Email).Message("Email is required")
	v.Required(user.Password).Message("Password is required")

	v.Match(user.Email, emailRegex).Message("Email is invalid")
	v.Match(user.Username, usernameRegex).Message("Username is invalid")
}

func (user *User) Create() User {
	u := User{
		Username:  user.Username,
		Fullname:  user.Fullname,
		Email:     user.Email,
		Password:  user.Password,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	app.DB.Create(&user)

	return u
}

func (user *User) GetUserByName(username string) User {
	u := User{}

	app.DB.Where("username = ?", username).First(&u)

	return u
}
