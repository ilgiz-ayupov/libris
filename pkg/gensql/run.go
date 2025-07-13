package gensql

import (
	"log/slog"

	"github.com/ilgiz-ayupov/libris/internal/entities"
)

func LoadData[T any](fn func() (T, error), log *slog.Logger, errNotFound error, msg string, attrs ...any) (T, error) {
	var zero T
	data, err := fn()
	switch err {
	case nil:
		return data, nil
	case entities.ErrNoData:
		log.With(attrs...).Warn(msg + errNotFound.Error())
		return zero, errNotFound
	default:
		attrs = append(attrs, slog.String("error", err.Error()))
		log.With(attrs...).Error(msg)
		return zero, entities.ErrInternalError
	}
}

func LoadCanNoData[T any](fn func() (T, error), log *slog.Logger, errNotFound error, msg string, attrs ...any) (T, error) {
	var zero T
	data, err := fn()
	switch err {
	case nil:
		return data, nil
	case entities.ErrNoData:
		log.With(attrs...).Warn(msg + errNotFound.Error())
		return zero, nil
	default:
		attrs = append(attrs, slog.String("error", err.Error()))
		log.With(attrs...).Error(msg)
		return zero, entities.ErrInternalError
	}
}

func LoadRequiredData[T any](fn func() (T, error), log *slog.Logger, errNotFound error, msg string, attrs ...any) (T, error) {
	var zero T
	data, err := fn()
	switch err {
	case nil:
		return data, nil
	case entities.ErrNoData:
		attrs = append(attrs, slog.String("error", errNotFound.Error()))
		log.With(attrs...).Error(msg)
		return zero, errNotFound
	default:
		attrs = append(attrs, slog.String("error", err.Error()))
		log.With(attrs...).Error(msg)
		return zero, entities.ErrInternalError
	}
}
