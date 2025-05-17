package dto

import (
	"time"

	"gorm.io/gorm"
)

type Present struct {
	CreatedAt  time.Time      `json:"createdAt"`
	UpdatedAt  time.Time      `json:"updatedAt"`
	DeletedAt  gorm.DeletedAt `json:"deletedAt"`
	Meeting    Meeting        `json:"Meeting"`
	Associated Associated     `json:"Associated"`
	Date       string         `json:"date"`
	Present    bool           `json:"present"`
}
