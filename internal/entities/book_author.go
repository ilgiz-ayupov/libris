package entities

import (
	"errors"

	"gorm.io/gorm"
)

var (
	ErrBookAuthorsNotFound = errors.New("авторы не найдены")
)

type BookAuthor struct {
	gorm.Model
	ID        int     `gorm:"primaryKey;autoIncrement"`
	FIO       string  `gorm:"unique; not null"`
	Biography string  `gorm:"not null"`
	Rating    float64 `gorm:"not null"`
}
