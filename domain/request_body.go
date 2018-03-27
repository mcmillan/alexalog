package domain

type RequestBody struct {
	Version string   `json:"version"`
	Session *Session `json:"session"`
	Context *Context `json:"context"`
	Request *Request `json:"request"`
}
