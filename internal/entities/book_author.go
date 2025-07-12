package entities

import "gorm.io/gorm"

type BookAuthor struct {
	gorm.Model
	ID   int    `gorm:"primaryKey;autoIncrement"`
	Name string `gorm:"unique; not null"`
}

func NewBookAuthor(name string) *BookAuthor {
	return &BookAuthor{
		Name: name,
	}
}
