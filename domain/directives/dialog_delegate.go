package directives

import "github.com/mcmillan/alexalog/domain"

type DialogDelegate struct {
	Type          string         `json:"type,omitempty"`
	UpdatedIntent *domain.Intent `json:"updatedIntent,omitempty"`
}
