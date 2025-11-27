package albums

import "gorm.io/gorm"

type Album struct {
	Title  string `json:"title" binding:"required"`
	Artist string `json:"artist" binding:"required"`
	gorm.Model
}
