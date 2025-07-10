package domain

import (
	"projeto-integrador-mdm/internal/database"
	"time"
)

type Presence struct {
	NumberCard int64     `json:"number_card"`
	MeetingID  int64     `json:"meeting_id"`
	Date       time.Time `json:"date"`
	Presence   bool      `json:"presence,omitempty"`
}

func (a Presence) ToCreateParams() database.CreatePresenceParams {
	return database.CreatePresenceParams{
		NumberCard: a.NumberCard,
		MeetingID:  a.MeetingID,
		Date:       a.Date,
		Present:    a.Presence,
	}
}
