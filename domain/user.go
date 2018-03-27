package domain

type User struct {
	UserID      string      `json:"userId"`
	AccessToken string      `json:"accessToken`
	Permissions interface{} `json:"permissions"`
}
