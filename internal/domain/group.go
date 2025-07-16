package domain

import (
	"log/slog"
	"projeto-integrador-mdm/internal/db"
	"time"
)

type Group struct {
	ID    int64     `json:"id"    validate:"required,gt=0"`
	Name  string    `json:"name"  validate:"required"`
	Hours time.Time `json:"hours" validate:"required"`
}

func (a Group) ToCreateParams() db.CreateGroupParams {
	slog.Debug("chamada de função ToCreateParams em Group")
	return db.CreateGroupParams{
		Name:  a.Name,
		Hours: a.Hours,
	}
}

func (a Group) ToUpdateParams() db.UpdateGroupParams {
	slog.Debug("chamada de função ToUpdateParams em Group")
	return db.UpdateGroupParams{
		ID:    a.ID,
		Name:  a.Name,
		Hours: a.Hours,
	}
}
