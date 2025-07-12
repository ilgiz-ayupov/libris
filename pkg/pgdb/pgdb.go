package pgdb

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Connect(url string) (*gorm.DB, error) {
	return gorm.Open(postgres.Open(url), &gorm.Config{})
}
