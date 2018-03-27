package dialog

import "github.com/mcmillan/alexalog/domain"

type DialogConfirmIntent struct {
	Type          string        `json:"type"`
	UpdatedIntent domain.Intent `json:"updatedIntent"`
}
