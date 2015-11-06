package controller

import (
	"bytes"
	"net/http"

	"gopkg.in/leyra/echo.v1"

	"leyra/app"
	"leyra/app/jobs"
	"leyra/app/models"
)

// Blog represents an instance of the Blog Controller. In this case, an empty
// buffer is created as the routes are being setup for funcs in this controller.
type Blog struct {
	Buffer *bytes.Buffer
}

// ListTemplate represents the data structure that will be passed through to the
// blog_list.html template.
type ListTemplate struct {
	Posts model.Posts
}

// ViewTemplate represents the data structure that will be passed through to the
// blog_view.html template.
type ViewTemplate struct {
	Post model.Post
}

// List will list the title of each blog post with a link to then go on to view
// the title and body of it using the view template.
func (b *Blog) List(c *echo.Context) error {
	posts := model.Posts{}
	app.S.DB.Find(&posts)

	app.S.View.ExecuteTemplate(b.Buffer, "blog_list.html", ListTemplate{
		Posts: posts,
	})

	return c.HTML(http.StatusOK, b.Buffer.String())
}

// View displays one blog post using the :id param passed through in the URL.
// This displays both the title and body for a given post.
func (b *Blog) View(c *echo.Context) error {
	post := model.Post{}
	app.S.DB.First(&post, c.Param("id"))

	app.S.View.ExecuteTemplate(b.Buffer, "blog_view.html", ViewTemplate{
		Post: post,
	})

	return c.HTML(http.StatusOK, b.Buffer.String())
}

// Create presents a form where the user can input the title and body of their
// new blog post.
func (b Blog) Create(c *echo.Context) error {
	if app.S.Get(c, "user") == nil {
		return c.Redirect(http.StatusMovedPermanently, "/")
	}

	app.S.View.ExecuteTemplate(b.Buffer, "blog_create.html", nil)

	return c.HTML(http.StatusOK, b.Buffer.String())
}

// Store saves the blog post.
func (b Blog) Store(c *echo.Context) error {
	create := job.CreateBlogPost{
		UserID: app.S.Get(c, "user").(int),
		Title:  c.Form("title"),
		Body:   c.Form("body"),
	}

	create.Handle()

	return c.Redirect(http.StatusMovedPermanently, "/")
}
