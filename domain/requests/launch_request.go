package requests

type LaunchRequest struct {
	Type      string `json:"type"`
	Timestamp string `json:"timestamp"`
	RequestID string `json:"requestId"`
	Locale    string `json:"locale"`
}
