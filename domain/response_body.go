package domain

type ResponseBody struct {
	Version           string             `json:"version"`
	SessionAttributes *SessionAttributes `json:"sessionAttributes,omitempty"`
	Response          *Response          `json:"response"`
}
