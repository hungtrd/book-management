package controllers

import (
	"fmt"
	"revel-app-demo/app/models"
	"revel-app-demo/app/routes"

	"github.com/revel/revel"
	"golang.org/x/crypto/bcrypt"
)

type App struct {
	*revel.Controller
}

func (c App) Index() revel.Result {
	c.Flash.Success("Welcome, ahihi")
	return c.Render()
}

func (c App) Login() revel.Result {
	return c.Render()
}

func (c App) UserLogin() revel.Result {
	return c.Render()
}

func (c App) Signup() revel.Result {
	return c.Render()
}

func (c App) UserSignup(user models.User, cpassword string) revel.Result {
	fmt.Println("user: ", user)
	fmt.Println("cpassword: ", cpassword)

	c.Validation.Required(cpassword).
		Message("Confirm password is required")
	c.Validation.Required(cpassword == user.Password).
		Message("Password does not match")
	user.Validate(c.Validation)

	if c.Validation.HasErrors() {
		c.Validation.Keep()
		c.FlashParams()
		return c.Redirect(routes.App.Signup())
	}

	user.Password = hashPassword(user.Password)

	user.Create()

	return c.Redirect(routes.Books.Index())
}

func (c App) Logout() revel.Result {
	for k := range c.Session {
		delete(c.Session, k)
	}
	return c.Redirect(routes.App.Index())
}

func hashPassword(s string) string {
	hashed, _ := bcrypt.GenerateFromPassword(
		[]byte(s), bcrypt.DefaultCost)
	return string(hashed)
}

func checkPasswordHash(p, h string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(h), []byte(p))
	return err == nil
}
