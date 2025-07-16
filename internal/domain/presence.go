package domain

import (
	"log/slog"
	"projeto-integrador-mdm/internal/db"
)

type Presence struct {
	NumberCard int64 `json:"number_card" validate:"required,gt=0"`
	MeetingID  int64 `json:"meeting_id"  validate:"required,gt=0"`
	IsPresence *bool `json:"is_presence" validate:"required"`
}

func (a Presence) ToCreateParams() db.CreatePresenceParams {
	slog.Debug("chamada de função ToCreateParams em Presence")
	return db.CreatePresenceParams{
		NumberCard: a.NumberCard,
		MeetingID:  a.MeetingID,
		IsPresence: *a.IsPresence,
	}
}

func (a Presence) ToUpdateParams() db.UpdatePresenceParams {
	slog.Debug("chamada de função ToUpdateParams em Presence")
	return db.UpdatePresenceParams{
		NumberCard: a.NumberCard,
		MeetingID:  a.MeetingID,
		IsPresence: *a.IsPresence,
	}
}
