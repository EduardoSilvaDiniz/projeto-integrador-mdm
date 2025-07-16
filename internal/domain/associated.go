package domain

import (
	"log/slog"
	"projeto-integrador-mdm/internal/db"
)

type Associated struct {
	NumberCard int64  `json:"number_card"    validate:"required,gte=0,lte=99999999"`
	GroupID    int64  `json:"group_id"       validate:"required,gte=0,lte=99999999"`
	Name       string `json:"name,omitempty" validate:"required"`
}

func (a Associated) ToCreateParams() db.CreateAssociatedParams {
	slog.Debug("chamada de função ToCreateParams em Associated")
	return db.CreateAssociatedParams{
		GroupID:    a.GroupID,
		Name:       a.Name,
		NumberCard: a.NumberCard,
	}
}

func (a Associated) ToUpdateParams() db.UpdateAssociatedParams {
	slog.Debug("chamada de função ToUpdateParams em Associated")
	return db.UpdateAssociatedParams{
		GroupID:    a.GroupID,
		Name:       a.Name,
		NumberCard: a.NumberCard,
	}
}
