package shortcuts

import "github.com/mcmillan/alexalog/domain"

func SpeechResponse(text string) *domain.Response {
	return &domain.Response{
		OutputSpeech: &domain.ResponseOutputSpeech{
			Type: "PlainText",
			Text: text,
		},
	}
}
