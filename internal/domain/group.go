package domain

import (
	"log/slog"
	"projeto-integrador-mdm/internal/db"
	"time"
)

type Group struct {
	ID    int64     `json:"id"`
	Name  string    `json:"name"`
	Hours time.Time `json:"hours"`
}

func (a Group) ToCreateParams() db.CreateGroupParams {
	slog.Debug("chamada para função ToCreateParams em Group")
	return db.CreateGroupParams{
		Name:  a.Name,
		Hours: a.Hours,
	}
}

func (a Group) ToUpdateParams() db.UpdateGroupParams {
	slog.Debug("chamada para função ToUpdateParams em Group")
	return db.UpdateGroupParams{
		ID:    a.ID,
		Name:  a.Name,
		Hours: a.Hours,
	}
}
