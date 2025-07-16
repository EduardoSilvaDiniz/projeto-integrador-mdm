package domain

import (
	"log/slog"
	"projeto-integrador-mdm/internal/db"
)

type PresenceByCompositeKey struct {
	NumberCard int64 `json:"number_card" validate:"required,gt=0"`
	MeetingID  int64 `json:"meeting_id"  validate:"required,gt=0"`
}

func (a PresenceByCompositeKey) ToCreateParams() db.DeletePresenceByCompositeKeyParams {
	slog.Debug("chamada de função ToCreateParams em PresenceByCompositeKey")
	return db.DeletePresenceByCompositeKeyParams{
		NumberCard: a.NumberCard,
		MeetingID:  a.MeetingID,
	}
}


func (a PresenceByCompositeKey) ToGetParams() db.GetPresenceByCompositeKeyParams {
	slog.Debug("chamada de função ToGetParams em PresenceByCompositeKey")
	return db.GetPresenceByCompositeKeyParams{
		NumberCard: a.NumberCard,
		MeetingID:  a.MeetingID,
	}
}

func (a PresenceByCompositeKey) ToDeleteParams() db.DeletePresenceByCompositeKeyParams{
	slog.Debug("chamada de função ToDeleteParams em PresenceByCompositeKey")
	return db.DeletePresenceByCompositeKeyParams{
		NumberCard: a.NumberCard,
		MeetingID:  a.MeetingID,
	}
}
