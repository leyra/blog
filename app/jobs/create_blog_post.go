package job

import (
	"leyra/app"
	"leyra/app/models"
)

type CreateBlogPost struct {
	Title  string
	Body   string
	UserID int
}

func (cbp CreateBlogPost) Handle() {
	// TODO: Validate these post fields here
	post := model.Post{
		UserID: cbp.UserID,
		Title:  cbp.Title,
		Body:   cbp.Body,
	}

	db := app.S.DB
	db.NewRecord(post)
	db.Create(&post)
}
