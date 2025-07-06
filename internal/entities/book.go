package entities

import (
	"database/sql"
	"time"
)

type Book struct {
	ID          int          `json:"id"`
	Title       string       `json:"title"`
	Description string       `json:"description"`
	Author      string       `json:"author"`
	Price       float64      `json:"price"`
	Year        int          `json:"year"`
	CreatedDate time.Time    `json:"created_date"`
	DeletedDate sql.NullTime `json:"deleted_date"`
}
