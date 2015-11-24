package controller

import (
	"bytes"
	"net/http"

	"gopkg.in/leyra/echo.v1"

	"leyra/app"
	"leyra/app/jobs"
)

// Auth represents an instance of the Auth Controller. In this case, an empty
// buffer is created as the routes are being setup for funcs in this controller.
type Auth struct {
	Controller
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
	authenticate := job.AuthenticateUser{
		c.Form("email"),
		c.Form("password"),
	}

	userID := authenticate.Handle().ID

	if userID == 0 {
		return c.Redirect(http.StatusMovedPermanently, "/login")
	}

	app.S.Set(c, "user", userID)

	return c.Redirect(http.StatusMovedPermanently, "/")
}

// RegisterForm presents the user with a form containing the relevant fields to
// register a new user to the site.
func (a Auth) RegisterForm(c *echo.Context) error {
	if IsAuthenticated(c) {
		return c.Redirect(http.StatusMovedPermanently, "/")
	}

	app.S.View.ExecuteTemplate(a.Buffer, "auth_register.html", nil)

	return c.HTML(http.StatusOK, a.Buffer.String())
}

// Register collects the form data from the registration form and attempts to
// create a user using these credentials.
func (a Auth) Register(c *echo.Context) error {
	create := job.CreateUser{
		FirstName: c.Form("first_name"),
		LastName:  c.Form("last_name"),
		Email:     c.Form("email"),
		Password:  c.Form("password"),
	}

	user := create.Handle()

	app.S.Set(c, "user", user.ID)

	return c.Redirect(http.StatusMovedPermanently, "/")
}

// Logout clears the user's session "user" data.
func (a Auth) Logout(c *echo.Context) error {
	app.S.Set(c, "user", nil)

	return c.Redirect(http.StatusMovedPermanently, "/")
}
