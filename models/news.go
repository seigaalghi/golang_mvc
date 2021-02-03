package models

import "gorm.io/gorm"

// News adalah model atau schema untuk news
type News struct {
	gorm.Model
	Title    string `json:"title" gorm:"type:varchar(100)"`
	Content  string `json:"content"`
	AuthorID int    `json:"authorId"`
	Author   Author `json:"author" gorm:"foreignKey:AuthorID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}
