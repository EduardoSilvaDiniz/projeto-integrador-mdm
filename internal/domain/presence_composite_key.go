package domain

import (
	"log/slog"
	"projeto-integrador-mdm/internal/db"
)

type PresenceByCompositeKey struct {
	NumberCard int64 `json:"number_card"`
	MeetingID  int64 `json:"meeting_id"`
}

func (a PresenceByCompositeKey) ToCreateParams() db.DeletePresenceByCompositeKeyParams {
	slog.Debug("chamada para função ToCreateParams em PresenceByCompositeKey")
	return db.DeletePresenceByCompositeKeyParams{
		NumberCard: a.NumberCard,
		MeetingID:  a.MeetingID,
	}
}
