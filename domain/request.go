package domain

type Request struct {
	// All
	Type      string `json:"type"`
	Timestamp string `json:"timestamp"`
	RequestID string `json:"requestId"`
	Locale    string `json:"locale"`

	// IntentRequest
	DialogState string `json:"dialogState"`
	Intent      Intent `json:"intent"`

	// SessionEndedRequest
	Reason string `json:"reason"`
	Error  Error  `json:"error"`
}
