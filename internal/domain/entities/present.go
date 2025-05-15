package domain

import "gorm.io/gorm"

type Present struct {
	gorm.Model
	Id         int        `json:"id"`
	Meeting    Meeting    `json:"Meeting"`
	Associated Associated `json:"Associated"`
	Date       string     `json:"date"`
	Present    bool       `json:"present"`
}
