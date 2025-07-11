package domain

import "projeto-integrador-mdm/internal/db"

type Associated struct {
	GroupID    int64  `json:"group_id"`
	Name       string `json:"name,omitempty"`
	NumberCard int64  `json:"number_card"`
}

func (a Associated) ToCreateParams() db.CreateAssociatedParams {
	return db.CreateAssociatedParams{
		GroupID:    a.GroupID,
		Name:       a.Name,
		NumberCard: a.NumberCard,
	}
}

func (a Associated) ToUpdateParams() db.UpdateAssociatedParams {
	return db.UpdateAssociatedParams{
		GroupID:    a.GroupID,
		Name:       a.Name,
		NumberCard: a.NumberCard,
	}
}

func (a Associated) ToStruct() db.Associated{
	return db.Associated{
		GroupID:    a.GroupID,
		Name:       a.Name,
		NumberCard: a.NumberCard,
	}
}
