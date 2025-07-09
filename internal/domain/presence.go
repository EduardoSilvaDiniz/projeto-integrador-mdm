package domain

type Presence struct {
	ID         int64  `json:"id"`
	IDMeeting  int64  `json:"meeting"`
	NumberCard int64  `json:"associated"`
	Date       string `json:"date,omitempty"`
	Present    bool   `json:"present,omitempty"`
}
