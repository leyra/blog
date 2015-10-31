package controller

import (
	"io/ioutil"
	"net/http"

	"gopkg.in/leyra/echo.v1"

	"leyra/app"
	"leyra/app/models"
)

type Blog struct {
}

func (b Blog) List(c *echo.Context) error {
	// Need to get all blog posts and list here
	return c.HTML(http.StatusOK, "List")
}

func (b Blog) Create(c *echo.Context) error {
	template, err := ioutil.ReadFile("./app/views/blog/create.html")
	if err != nil {
		panic(err)
	}

	return c.HTML(http.StatusOK, string(template))
}

func (b Blog) Store(c *echo.Context) error {
	post := model.Post{
		Title: c.Form("title"),
		Body:  c.Form("body"),
	}

	db := app.S.DB
	db.NewRecord(post)
	db.Create(&post)

	return c.HTML(http.StatusOK, "Store")
}
