package domain

type ResponseBody struct {
	Version           string            `json:"version"`
	SessionAttributes SessionAttributes `json:"sessionAttributes"`
	Response          Response          `json:"response"`
}
