package domain

type Device struct {
	DeviceID            string   `json:"deviceId"`
	SupportedInterfaces []string `json:"supportedInterfaces"`
}
