package entities

import "errors"

var (
	ErrInternalError = errors.New("внутренняя ошибка")
	ErrNoData        = errors.New("нет данных")
)
