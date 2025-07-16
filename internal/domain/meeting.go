package domain

import (
	"log/slog"
	"projeto-integrador-mdm/internal/db"
	"time"
)

type Meeting struct {
	ID      int64     `json:"id"`
	GroupID int64     `json:"group_id"          validate:"required,gt=0"`
	Date    time.Time `json:"date"              validate:"required"`
	Address string    `json:"address,omitempty" validate:"required"`
}

func (a Meeting) ToCreateParams() db.CreateMeetingParams {
	slog.Debug("chamada de função ToCreateParams em Meeting")
	return db.CreateMeetingParams{
		GroupID: a.GroupID,
		Address: a.Address,
		Date:    a.Date,
	}
}

func (a Meeting) ToUpdateParams() db.UpdateMeetingParams {
	slog.Debug("chamada de função ToUpdateParams em Meeting")
	return db.UpdateMeetingParams{
		ID:      a.ID,
		GroupID: a.ID,
		Address: a.Address,
		Date:    a.Date,
	}
}
