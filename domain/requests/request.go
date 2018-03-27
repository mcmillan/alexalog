package requests

import (
	"errors"
	"fmt"
)

type genericRequest struct {
	Type string `json:"type"`
}

func Parse(req interface{}) (IntentRequest, LaunchRequest, SessionEndedRequest, error) {
	intentRequest := IntentRequest{}
	launchRequest := LaunchRequest{}
	sessionEndedRequest := SessionEndedRequest{}

	genericReq, ok := req.(genericRequest)

	if !ok {
		return intentRequest, launchRequest, sessionEndedRequest, errors.New("Unable to convert input to a request")
	}

	ok = false

	switch genericReq.Type {
	case "IntentRequest":
		intentRequest, ok = req.(IntentRequest)
	case "LaunchRequest":
		launchRequest, ok = req.(LaunchRequest)
	case "SessionEndedRequest":
		sessionEndedRequest, ok = req.(SessionEndedRequest)
	}

	if !ok {
		return intentRequest, launchRequest, sessionEndedRequest, fmt.Errorf("Unable to convert input to type %s", genericReq.Type)
	}

	return intentRequest, launchRequest, sessionEndedRequest, nil
}
