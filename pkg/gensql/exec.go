package gensql

import "github.com/jmoiron/sqlx"

func Exec(tx *sqlx.Tx, query string, args map[string]any) error {
	stmt, err := tx.PrepareNamed(query)
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(args)
	return err
}

func ExecReturnID[T any](tx *sqlx.Tx, query string, args map[string]any) (T, error) {
	var (
		zero T
		data T
	)

	stmt, err := tx.PrepareNamed(query)
	if err != nil {
		return zero, err
	}
	defer stmt.Close()

	if err := stmt.QueryRow(args).Scan(&data); err != nil {
		return zero, err
	}

	return data, nil
}
