package domain

import (
	"log/slog"
	"projeto-integrador-mdm/internal/db"
)

type Presence struct {
	NumberCard int64 `json:"number_card"`
	MeetingID  int64 `json:"meeting_id"`
	IsPresence bool  `json:"is_presence,omitempty"`
}

func (a Presence) ToCreateParams() db.CreatePresenceParams {
	slog.Debug("chamada para função ToCreateParams em Presence")
	return db.CreatePresenceParams{
		NumberCard: a.NumberCard,
		MeetingID:  a.MeetingID,
		IsPresence: a.IsPresence,
	}
}

func (a Presence) ToUpdateParams() db.UpdatePresenceParams {
	slog.Debug("chamada para função ToUpdateParams em Presence")
	return db.UpdatePresenceParams{
		NumberCard: a.NumberCard,
		MeetingID:  a.MeetingID,
		IsPresence: a.IsPresence,
	}
}
