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

func (c App) GetUser() revel.Result {
	if user := c.connected(); user != nil {
		c.ViewArgs["user"] = user
	}
	return nil
}

func (c App) Index() revel.Result {
	return c.Redirect(routes.Books.Index())
}

func (c App) Login() revel.Result {
	// check user login
	if user := c.connected(); user != nil {
		c.Flash.Error("You have logged in")
		return c.Redirect(routes.Books.Index())
	}

	return c.Render()
}

func (c App) UserLogin(username, password string) revel.Result {
	c.Validation.Required(username).
		Message("Username is required")
	c.Validation.Required(password).
		Message("Password is required")

	if c.Validation.HasErrors() {
		c.Flash.Error("Please fix errors below")
		c.Validation.Keep()
		c.FlashParams()
		return c.Redirect(routes.App.Login())
	}

	user := c.getUser(username)
	if user != nil {
		ok := checkPasswordHash(password, user.Password)
		if ok {
			c.Session["user"] = username
			c.Session.SetDefaultExpiration()
			c.Flash.Success("Welcome, " + username)
			return c.Redirect(routes.Books.Index())
		}
	}

	c.Flash.Out["username"] = username
	c.Flash.Error("Invalid username or password!")
	return c.Redirect(routes.App.Login())
}

func (c App) Signup() revel.Result {
	// check user login
	if user := c.connected(); user != nil {
		c.Flash.Error("You have logged in")
		return c.Redirect(routes.Books.Index())
	}

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
		c.Flash.Error("Please fix errors below")
		c.Validation.Keep()
		c.FlashParams()
		return c.Redirect(routes.App.Signup())
	}

	// Check username exist
	u := user.GetUserByName(user.Username)
	if u.ID != 0 {
		c.Flash.Error("Please fix errors below")
		c.Validation.Error("Username existed", user.Username).
			Message("Username existed").
			Key("user.Username")
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
	return c.Redirect(routes.Books.Index())
}

func (c App) connected() *models.User {
	if c.ViewArgs["user"] != nil {
		return c.ViewArgs["user"].(*models.User)
	}
	if username, ok := c.Session["user"]; ok {
		return c.getUser(username.(string))
	}

	return nil
}

func (c App) getUser(username string) (user *models.User) {
	user = &models.User{}
	_, err := c.Session.GetInto("fulluser", user, false)
	if err != nil {
		c.Log.Error("Session.GetInto failed: %w", err)
	}
	if user.Username == username {
		return user
	}

	u := user.GetUserByName(username)
	if u.ID == 0 {
		return nil
	}
	// c.Session["fulluser"] = u
	return &u
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
