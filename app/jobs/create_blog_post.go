package job

import (
	"leyra/app"
	"leyra/app/models"
)

type CreateBlogPost struct {
	Title string
	Body  string
}

func (cbp CreateBlogPost) Handle() {
	// TODO: Validate these post fields here
	post := model.Post{
		Title: cbp.Title,
		Body:  cbp.Body,
	}

	db := app.S.DB
	db.NewRecord(post)
	db.Create(&post)
}
