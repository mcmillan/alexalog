package dialog

import "github.com/mcmillan/alexalog/domain"

type DialogElicitSlot struct {
	Type          string        `json:"type"`
	SlotToElicit  string        `json:"slotToElicit"`
	UpdatedIntent domain.Intent `json:"updatedIntent"`
}
