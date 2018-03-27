package directives

import "github.com/mcmillan/alexalog/domain"

type DialogDelegate struct {
	Type          string        `json:"type"`
	UpdatedIntent domain.Intent `json:"updatedIntent`
}
