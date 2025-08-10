package gensql

import (
	"database/sql"

	"github.com/ilgiz-ayupov/libris/internal/entities"
	"github.com/jmoiron/sqlx"
)

func Select[T any](tx *sqlx.Tx, query string, args map[string]any) ([]T, error) {
	stmt, err := tx.PrepareNamed(query)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	var data []T
	if err := stmt.Select(&data, args); err != nil {
		return nil, err
	}

	if len(data) == 0 {
		return nil, entities.ErrNoData
	}

	return data, nil
}

func Get[T any](tx *sqlx.Tx, query string, args map[string]any) (T, error) {
	var (
		zero T
		data T
	)

	stmt, err := tx.PrepareNamed(query)
	if err != nil {
		return zero, err
	}
	defer stmt.Close()

	switch err := stmt.Get(&data, args); err {
	case nil:
		return data, nil
	case sql.ErrNoRows:
		return zero, entities.ErrNoData
	default:
		return zero, err
	}
}

func SelectRebind[T any](tx *sqlx.Tx, query string, args map[string]any) ([]T, error) {
	query, params, err := sqlx.Named(query, args)
	if err != nil {
		return nil, err
	}

	query, params, err = sqlx.In(query, params...)
	if err != nil {
		return nil, err
	}

	query = tx.Rebind(query)

	var data []T
	err = tx.Select(&data, query, params...)
	if err != nil {
		return nil, err
	}

	if len(data) == 0 {
		return nil, entities.ErrNoData
	}

	return data, nil
}
