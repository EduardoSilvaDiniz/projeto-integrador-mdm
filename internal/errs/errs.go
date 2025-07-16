package errs

import "errors"

var (
	ErrNotFound      = errors.New("registro não encontrado")
	ErrInvalidInput  = errors.New("dados inválidos")
	ErrConflict      = errors.New("conflito de dados")
	ErrUnauthorized  = errors.New("não autorizado")
	ErrInternalError = errors.New("erro interno")
)
