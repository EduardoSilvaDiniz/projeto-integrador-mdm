package dto

import (
	"time"

	"gorm.io/gorm"
)

type Meeting struct {
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at"`
	ProjectId int            `json:"project_id,omitempty"`
	GroupId   int            `json:"group_id,omitempty"`
	HubId     int            `json:"hub_id,omitempty"`
	Date      string         `json:"date,omitempty"`
	Time      string         `json:"time,omitempty"`
}
