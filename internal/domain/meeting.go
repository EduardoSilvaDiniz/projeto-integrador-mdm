package domain

import (
	"log/slog"
	"projeto-integrador-mdm/internal/db"
	"time"
)

type Meeting struct {
	ID      int64     `json:"id"`
	GroupID int64     `json:"group_id"`
	Date    time.Time `json:"date"`
	Address string    `json:"address,omitempty"`
}

func (a Meeting) ToCreateParams() db.CreateMeetingParams {
	slog.Debug("chamada de função ToCreateParams em Meeting")
	return db.CreateMeetingParams{
		GroupID: a.ID,
		Address: a.Address,
		Date:    a.Date,
	}
}

func (a Meeting) ToUpdateParams() db.UpdateMeetingParams {
	slog.Debug("chamada de função ToUpdateParams em Associated")
	return db.UpdateMeetingParams{
		GroupID: a.ID,
		Address: a.Address,
		Date:    a.Date,
		ID:      a.ID,
	}
}
