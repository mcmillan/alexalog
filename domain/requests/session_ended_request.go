package requests

import "github.com/mcmillan/alexalog/domain"

type SessionEndedRequest struct {
	Type      string       `json:"type"`
	Timestamp string       `json:"timestamp"`
	RequestID string       `json:"requestId"`
	Locale    string       `json:"locale"`
	Reason    string       `json:"reason"`
	Error     domain.Error `json:"error"`
}
