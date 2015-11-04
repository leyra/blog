package controller

import (
	"bytes"
	"net/http"

	"gopkg.in/leyra/echo.v1"

	"leyra/app"
	"leyra/app/models"
)

type Blog struct {
}

type ListTemplate struct {
	Posts model.Posts
}

type ViewTemplate struct {
	Post model.Post
}

// List will list the title of each blog post with a link to then go on to view
// the title and body of it using the view template.
func (b Blog) List(c *echo.Context) error {
	buff := new(bytes.Buffer)

	posts := model.Posts{}
	app.S.DB.Find(&posts)

	app.S.Template.ExecuteTemplate(buff, "list.html", ListTemplate{
		Posts: posts,
	})

	return c.HTML(http.StatusOK, buff.String())
}

// View displays one blog post using the :id param passed through in the URL.
// This displays both the title and body for a given post.
func (b Blog) View(c *echo.Context) error {
	post := model.Post{}
	app.S.DB.First(&post, c.Param("id"))

	buff := new(bytes.Buffer)
	app.S.Template.ExecuteTemplate(buff, "view.html", ViewTemplate{
		Post: post,
	})

	return c.HTML(http.StatusOK, buff.String())
}

// Create presents a form where the user can input the title and body of their
// new blog post.
func (b Blog) Create(c *echo.Context) error {
	buff := new(bytes.Buffer)
	app.S.Template.ExecuteTemplate(buff, "create.html", nil)

	return c.HTML(http.StatusOK, buff.String())
}

// Store saves the blog post.
func (b Blog) Store(c *echo.Context) error {
	post := model.Post{
		Title: c.Form("title"),
		Body:  c.Form("body"),
	}

	db := app.S.DB
	db.NewRecord(post)
	db.Create(&post)

	return c.Redirect(http.StatusMovedPermanently, "/")
}
