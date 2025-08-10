package entities

import (
	"errors"
)

var (
	ErrBookPublisherNotFound = errors.New("редактор не найден")
)

type BookPublisher struct {
	ID          int    `db:"publisher_id"`
	Name        string `db:"name"`
	Description string `db:"description"`
}
