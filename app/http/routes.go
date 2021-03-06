package http

import (
	"bytes"

	"gopkg.in/leyra/echo.v1"

	"leyra/app/http/controllers"
)

// Route currently creates a new instance of echo and attaches routes to
// patterns that can be defined in this file. I'm still not so sure about how
// all of this should work - for now it's fine though.
func Route() *echo.Echo {
	e := echo.New()

	// Blog Controller
	e.Get("/", routeBlogList)
	e.Get("/new", routeBlogCreate)
	e.Post("/new", routeBlogStore)
	e.Get("/post/:id", routeBlogPost)

	// Auth Controller
	e.Get("/login", routeLoginForm)
	e.Post("/login", routeLogin)

	e.Get("/register", routeRegisterForm)
	e.Post("/register", routeRegister)

	e.Get("/logout", routeLogout)

	return e
}

// Inject an empty instance of bytes.Buffer into this controller as it's used
// multiple times and makes the code a little bit cleaner.
func blogController() *controller.Blog {
	c := new(controller.Blog)
	c.Buffer = new(bytes.Buffer)

	return c
}

func routeBlogList(c *echo.Context) error {
	return blogController().List(c)
}

func routeBlogCreate(c *echo.Context) error {
	return blogController().Create(c)
}

func routeBlogStore(c *echo.Context) error {
	return blogController().Store(c)
}

func routeBlogPost(c *echo.Context) error {
	return blogController().View(c)
}

// Inject an empty instance of bytes.Buffer into this controller as it's used
// multiple times and makes the code a little bit cleaner.
func authController() *controller.Auth {
	c := new(controller.Auth)
	c.Buffer = new(bytes.Buffer)

	return c
}

func routeLoginForm(c *echo.Context) error {
	return authController().LoginForm(c)
}

func routeLogin(c *echo.Context) error {
	return authController().Login(c)
}

func routeRegisterForm(c *echo.Context) error {
	return authController().RegisterForm(c)
}

func routeRegister(c *echo.Context) error {
	return authController().Register(c)
}

func routeLogout(c *echo.Context) error {
	return authController().Logout(c)
}
