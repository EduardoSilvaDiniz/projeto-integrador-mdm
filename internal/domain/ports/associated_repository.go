package ports

import "chamada-pagamento-system/internal/transport/http-server/dto"

type AssociatedReposity interface {
	AssociatedInsert(assoc dto.Associated)
	AssociatedGet()
	AssociatedDelete(id int)
	AssociatedUpdate(id int)
}

type AssociatedReposityImpl struct {}



