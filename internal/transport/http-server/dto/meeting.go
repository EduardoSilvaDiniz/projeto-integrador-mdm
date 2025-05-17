package dto

type Meeting struct {
	ProjectId int    `json:"projectId"`
	GroupId   int    `json:"groupId"`
	HubId     int    `json:"hubId"`
	Date      string `json:"date"`
	Time      string `json:"time"`
}
