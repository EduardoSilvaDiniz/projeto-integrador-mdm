package domain

import (
	"log/slog"
	"projeto-integrador-mdm/internal/db"
)

type Associated struct {
	GroupID    int64  `json:"group_id"`
	Name       string `json:"name,omitempty"`
	NumberCard int64  `json:"number_card"`
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
