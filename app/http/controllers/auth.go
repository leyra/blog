package controller

import (
	"bytes"
	"net/http"

	"gopkg.in/leyra/echo.v1"

	"leyra/app"
)

// Auth represents an instance of the Auth Controller. In this case, an empty
// buffer is created as the routes are being setup for funcs in this controller.
type Auth struct {
	Buffer *bytes.Buffer
}

// LoginForm presents the user with a form containing email and password fields
// for it then to be posted to controller.Auth.Login(c).
func (a *Auth) LoginForm(c *echo.Context) error {
	app.S.View.ExecuteTemplate(a.Buffer, "auth_login.html", nil)

	return c.HTML(http.StatusOK, a.Buffer.String())
}

// Login will check to see if the credentials are correct then proceed to log
// the user in if possible.
func (a Auth) Login(c *echo.Context) error {

	return c.HTML(http.StatusOK, "Logged in")
}

// RegisterForm presents the user with a form containing the relevant fields to
// register a new user to the site.
func (a Auth) RegisterForm(c *echo.Context) error {
	return c.HTML(http.StatusOK, "Register")
}
