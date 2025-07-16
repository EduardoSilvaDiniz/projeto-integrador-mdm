package errs

import "errors"

var (
	ErrNotFound      = errors.New("registro não encontrado")
	ErrInvalidInput  = errors.New("dados inválidos")
	ErrAlreadyExists = errors.New("conflito de dados: o registro já está cadastrado")
	ErrUnauthorized  = errors.New("não autorizado")
	ErrInternalError = errors.New("erro interno")
)
