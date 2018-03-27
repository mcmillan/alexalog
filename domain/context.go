package domain

type SystemContext struct {
	ApiAccessToken string       `json:"apiAccessToken"`
	ApiEndpoint    string       `json:"apiEndpoint"`
	Application    *Application `json:"application"`
	Device         *Device      `json:"device"`
	User           *User        `json:"user"`
}

type Context struct {
	System *SystemContext `json:"System"`
}
