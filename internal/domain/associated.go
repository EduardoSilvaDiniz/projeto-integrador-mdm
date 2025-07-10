package domain

import "projeto-integrador-mdm/internal/database"

type Associated struct {
	NumberCard int64  `json:"number_card"`
	Name       string `json:"name,omitempty"`
	GroupID    int64  `json:"group_id"`
}

func (a Associated) ToCreateParams() database.CreateAssociatedParams {
	return database.CreateAssociatedParams{
		NumberCard: a.NumberCard,
		Name:       a.Name,
		GroupID:    a.GroupID,
	}
}
