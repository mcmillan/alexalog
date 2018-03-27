package domain

type Intent struct {
	Name               string           `json:"name"`
	ConfirmationStatus string           `json:"confirmationStatus"`
	Slots              map[string]*Slot `json:"slots"`
}
