package infra

import "chamada-pagamento-system/internal/domain"

func StartDb(size int) map[int]domain.Associated {
	db := make(map[int]domain.Associated, size)
	return db
}
