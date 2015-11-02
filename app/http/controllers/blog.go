package controller

import (
	"bytes"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"

	"gopkg.in/leyra/echo.v1"

	"leyra/app"
	"leyra/app/models"
)

type Blog struct {
}

// List will list the title of each blog post with a link to then go on to view
// the title and body of it using the view template.
func (b Blog) List(c *echo.Context) error {
	db := app.S.DB

	posts := model.Posts{}
	db.Find(&posts)

	html, err := ioutil.ReadFile("./app/views/blog/list.html")

	if err != nil {
		log.Fatal(err)
	}

	t, err := template.New("list").Parse(string(html))
	buff := new(bytes.Buffer)

	data := struct {
		Posts model.Posts
	}{
		Posts: posts,
	}

	err = t.Execute(buff, data)

	if err != nil {
		log.Fatal(err)
	}

	// Need to get all blog posts and list here
	return c.HTML(http.StatusOK, buff.String())
}

// View displays one blog post using the :id param passed through in the URL.
// This displays both the title and body for a given post.
func (b Blog) View(c *echo.Context) error {
	id := c.Param("id")

	db := app.S.DB
	post := model.Post{}

	db.First(&post, id)

	html, err := ioutil.ReadFile("./app/views/blog/view.html")

	if err != nil {
		log.Fatal(err)
	}

	t, err := template.New("view").Parse(string(html))

	buff := new(bytes.Buffer)

	data := struct {
		Post model.Post
	}{
		Post: post,
	}

	err = t.Execute(buff, data)

	if err != nil {
		log.Fatal(err)
	}

	return c.HTML(http.StatusOK, buff.String())
}

// Create presents a form where the user can input the title and body of their
// new blog post.
func (b Blog) Create(c *echo.Context) error {
	template, err := ioutil.ReadFile("./app/views/blog/create.html")
	if err != nil {
		panic(err)
	}

	return c.HTML(http.StatusOK, string(template))
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
