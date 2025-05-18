package dto

import (
	"time"

	"gorm.io/gorm"
)

type Present struct {
	CreatedAt  time.Time      `json:"created_at"`
	UpdatedAt  time.Time      `json:"updated_at"`
	DeletedAt  gorm.DeletedAt `json:"deleted_at"`
	Meeting    Meeting        `json:"meeting,omitempty"`
	Associated Associated     `json:"associated,omitempty"`
	Date       string         `json:"date,omitempty"`
	Present    bool           `json:"present,omitempty"`
}
