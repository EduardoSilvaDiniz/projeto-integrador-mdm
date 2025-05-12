package domain

type Meeting struct {
	Id        int    `json:"id"`
	ProjectId int    `json:"projectId"`
	GroupId   int    `json:"groupId"`
	HubId     int    `json:"hubId"`
	Date      string `json:"date"`
	Time      string `json:"time"`
}
