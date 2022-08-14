package models

import "gorm.io/gorm"

type Comment struct {
	gorm.Model        // adds ID, CreatedAt, UpdatedAt & DeletedAt
	ID         string `json:"id"`
	Title      string `json:"title"`
	Author     string `json:"author"`
	Slug       string `json:"slug"`
}
