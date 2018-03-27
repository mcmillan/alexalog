package requests

import "github.com/mcmillan/alexalog/domain"

type IntentRequest struct {
	Type        string        `json:"type"`
	Timestamp   string        `json:"timestamp"`
	RequestID   string        `json:"requestId"`
	Locale      string        `json:"locale"`
	DialogState string        `json:"dialogState"`
	Intent      domain.Intent `json:"intent"`
}
