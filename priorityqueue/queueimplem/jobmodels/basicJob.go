package queueimplem

import "time"

type BasicJob struct {
	Id         string    `json:"id"`
	UserId     string    `json:"userId"`
	PromptData string    `json:"promptData"`
	TimeIssued time.Time `json:"timeIssued"`
	Priority   int32     `json:"priority"`
	Index      int
}
