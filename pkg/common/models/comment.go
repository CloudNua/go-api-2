package models

type Comment struct {
	// gorm.Model // adds ID, Title, Author & Slug
	// ID     string `json:"id" gorm:"primary_key"`
	Title  string `json:"title"`
	Author string `json:"author"`
	Slug   string `json:"slug"`
}

type AddCommentRequestBody struct {
	// gorm.Model
	// ID     string `json:"id"`
	Title  string `json:"title"`
	Author string `json:"author"`
	Slug   string `json:"slug"`
}

type UpdateCommentRequestBody struct {
	Title  string `json:"title"`
	Author string `json:"author"`
	Slug   string `json:"slug"`
}
