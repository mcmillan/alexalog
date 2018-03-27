package domain

type SessionAttributes map[string]interface{}

type Session struct {
	New         bool              `json:"new"`
	SessionID   string            `json:"sessionId"`
	Attributes  SessionAttributes `json:"attributes"`
	Application Application       `json:"application"`
	User        User              `json:"user"`
}
