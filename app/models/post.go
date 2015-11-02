package model

type Post struct {
	ID    int
	Title string
	Body  string `sql:"type:text"`
}

type Posts []Post
