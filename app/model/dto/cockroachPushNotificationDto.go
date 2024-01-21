package dto

type CockroachPushNotificationDto struct {
	Title        string `json:"title,omitempty"`
	Amount       uint32 `json:"amount,omitempty"`
	ReportedTime string `json:"reported_time,omitempty"`
}
