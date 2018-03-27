package dialog

import "github.com/mcmillan/alexalog/domain"

type DialogConfirmSlot struct {
	Type          string        `json:"type"`
	SlotToConfirm string        `json:"slotToConfirm"`
	UpdatedIntent domain.Intent `json:"updatedIntent"`
}
