package model

type Post struct {
	Title string
	Body  string `sql:"type:text"`
}
