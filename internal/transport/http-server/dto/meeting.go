package dto

import (
	"time"

	"gorm.io/gorm"
)

type Meeting struct {
	CreatedAt time.Time      `json:"createdAt"`
	UpdatedAt time.Time      `json:"updatedAt"`
	DeletedAt gorm.DeletedAt `json:"deletedAt"`
	ProjectId int            `json:"projectId"`
	GroupId   int            `json:"groupId"`
	HubId     int            `json:"hubId"`
	Date      string         `json:"date"`
	Time      string         `json:"time"`
}
