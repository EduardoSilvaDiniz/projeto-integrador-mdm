package domain

import "projeto-integrador-mdm/internal/database"

type PresenceCompositeKey struct {
	NumberCard int64 `json:"number_card"`
	MeetingID  int64 `json:"meeting_id"`
}

func (a PresenceCompositeKey) ToCreateParams() database.DeletePresenceByCompositeKeyParams {
	return database.DeletePresenceByCompositeKeyParams{
		NumberCard: a.NumberCard,
		MeetingID:  a.MeetingID,
	}
}
