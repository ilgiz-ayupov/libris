package entities

import (
	"errors"
)

var (
	ErrBookAuthorsNotFound = errors.New("авторы не найдены")
)

type BookAuthor struct {
	ID        int     `db:"author_id"`
	FIO       string  `db:"fio"`
	Biography string  `db:"biography"`
	Rating    float64 `db:"rating"`
}
