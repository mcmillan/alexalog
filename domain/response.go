package domain

type ResponseOutputSpeech struct {
	Type string `json:"type"`
	Text string `json:"text"`
	SSML string `json:"ssml"`
}

type ResponseCardImage struct {
	SmallImageURL string `json:"smallImageUrl"`
	LargeImageURL string `json:"largeImageUrl"`
}

type ResponseCard struct {
	Type    string             `json:"type"`
	Title   string             `json:"title"`
	Content string             `json:"content"`
	Text    string             `json:"text"`
	Image   *ResponseCardImage `json:"image"`
}

type ResponseReprompt struct {
	OutputSpeech *ResponseOutputSpeech `json:"outputSpeech"`
}

type Response struct {
	OutputSpeech     *ResponseOutputSpeech `json:"outputSpeech,omitempty"`
	Card             *ResponseCard         `json:"card,omitempty"`
	Reprompt         *ResponseReprompt     `json:"reprompt,omitempty"`
	ShouldEndSession bool                  `json:"shouldEndSession"`
	Directives       []interface{}         `json:"directives,omitempty"`
}
