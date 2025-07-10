package domain

import (
	"projeto-integrador-mdm/internal/database"
	"time"
)

type Meeting struct {
	ID      int64     `json:"id"`
	Date    time.Time `json:"date"`
	Address string    `json:"address,omitempty"`
}

func (a Meeting) ToCreateParams() database.CreateMeetingParams {
	return database.CreateMeetingParams{
		GroupID: a.ID,
		Date:    a.Date,
		Address: a.Address,
	}
}
