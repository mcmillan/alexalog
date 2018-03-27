package shortcuts

import "github.com/mcmillan/alexalog/domain"

func SpeechResponse(text string, endSession bool) domain.ResponseBody {
	return domain.ResponseBody{
		Version: "1.0",
		Response: &domain.Response{
			OutputSpeech: &domain.ResponseOutputSpeech{
				Type: "PlainText",
				Text: text,
			},
			ShouldEndSession: endSession,
		},
	}
}
