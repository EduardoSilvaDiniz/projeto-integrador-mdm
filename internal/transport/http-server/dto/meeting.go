package dto

type Meeting struct {
	ProjectId int    `json:"project_id,omitempty"`
	GroupId   int    `json:"group_id,omitempty"`
	HubId     int    `json:"hub_id,omitempty"`
	Date      string `json:"date,omitempty"`
	Time      string `json:"time,omitempty"`
}
