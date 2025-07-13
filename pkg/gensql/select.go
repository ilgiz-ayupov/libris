package gensql

import (
	"github.com/ilgiz-ayupov/libris/internal/entities"
	"gorm.io/gorm"
)

func Select[T any](query *gorm.DB, conds ...any) ([]T, error) {
	var data []T
	if err := query.Find(&data, conds...).Error; err != nil {
		return nil, err
	}

	if len(data) == 0 {
		return nil, entities.ErrNoData
	}

	return data, nil
}

func Get[T any](query *gorm.DB, conds ...any) (T, error) {
	var zero T

	var data T
	switch err := query.First(&data, conds...).Error; err {
	case nil:
		return data, nil
	case gorm.ErrRecordNotFound:
		return zero, entities.ErrNoData
	default:
		return zero, err
	}
}
