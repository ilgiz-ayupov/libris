package gensql

import (
	"github.com/ilgiz-ayupov/libris/internal/entities"
	"github.com/ilgiz-ayupov/libris/pkg/logger"
)

func LoadData[T any](fn func() (T, error), log logger.Logger, errNotFound error, msg string, logArgs ...any) (T, error) {
	var zero T

	data, err := fn()
	switch err {
	case nil:
		return data, nil
	case entities.ErrNoData:
		log.Warn(msg+errNotFound.Error(), logArgs...)
		return zero, errNotFound
	default:
		logArgs = append(logArgs, "error", err)
		log.Error(msg, logArgs...)
		return zero, entities.ErrInternalError
	}
}

func LoadCanNoData[T any](fn func() (T, error), log logger.Logger, errNotFound error, msg string, logArgs ...any) (T, error) {
	var zero T

	data, err := fn()
	switch err {
	case nil:
		return data, nil
	case entities.ErrNoData:
		log.Warn(msg+errNotFound.Error(), logArgs...)
		return zero, nil
	default:
		logArgs = append(logArgs, "error", err.Error)
		log.Error(msg, logArgs...)
		return zero, entities.ErrInternalError
	}
}

func LoadRequiredData[T any](fn func() (T, error), log logger.Logger, errNotFound error, msg string, logArgs ...any) (T, error) {
	var zero T

	data, err := fn()
	switch err {
	case nil:
		return data, nil
	case entities.ErrNoData:
		logArgs = append(logArgs, "error", errNotFound)
		log.Error(msg, logArgs...)
		return zero, errNotFound
	default:
		logArgs = append(logArgs, "error", err)
		log.Error(msg, logArgs...)
		return zero, entities.ErrInternalError
	}
}
