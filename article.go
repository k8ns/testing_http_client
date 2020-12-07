package resttest

type Article struct {
	Id int `json:"id"`
	Title string `json:"title"validate:"required"`
	Author string `json:"author"validate:"required"`
	Body string `json:"body"validate:"required"`
}
