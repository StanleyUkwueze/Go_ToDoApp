package models

import "gorm.io/gorm"

type TaskModel struct {
	gorm.Model
	Title       string
	IsCompleted bool
	UserId      uint
}
