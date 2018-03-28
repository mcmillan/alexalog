package server

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/mcmillan/alexalog/config"
	"github.com/mcmillan/alexalog/domain"
	"github.com/mcmillan/alexalog/domain/directives"
	"github.com/mcmillan/alexalog/shortcuts"
	log "github.com/sirupsen/logrus"
)

var globalConfig *config.Config

func handleError(w http.ResponseWriter, err error) {
	log.Error(err)

	response := domain.ResponseBody{
		Version:  "1.0",
		Response: shortcuts.SpeechResponse(fmt.Sprintf("Something went wrong: %s", err.Error())),
	}

	err = json.NewEncoder(w).Encode(response)

	if err != nil {
		log.Error(err)
		http.Error(w, "internal server error", 500)
	}
}

func handler(w http.ResponseWriter, r *http.Request) {
	log.Printf("%s %s", r.Method, r.RequestURI)

	if r.Method != "POST" {
		http.Error(w, "method not allowed", 405)
		return
	}

	var requestBody domain.RequestBody

	err := json.NewDecoder(r.Body).Decode(&requestBody)

	if err != nil {
		return
	}

	if requestBody.Request.Type != "IntentRequest" {
		http.Error(w, "not intentrequest", 422)
		return
	}

	var intentConfig *config.Intent

	for _, intent := range globalConfig.Intents {
		if intent.Name != requestBody.Request.Intent.Name {
			continue
		}

		intentConfig = intent
	}

	if intentConfig == nil {
		err = fmt.Errorf("cannot find intent %s", requestBody.Request.Intent.Name)
		handleError(w, err)
		return
	}

	responseBody := domain.ResponseBody{
		Version:  "1.0",
		Response: &domain.Response{},
	}

	if requestBody.Request.DialogState == "STARTED" {
		err := intentConfig.OnStart(requestBody.Request.Intent)
		if err != nil {
			handleError(w, err)
			return
		}
		responseBody.Response.Directives = []interface{}{
			directives.DialogDelegate{
				Type:          "Dialog.Delegate",
				UpdatedIntent: requestBody.Request.Intent,
			},
		}
	} else if requestBody.Request.DialogState != "COMPLETED" {
		responseBody.Response.Directives = []interface{}{
			directives.DialogDelegate{Type: "Dialog.Delegate"},
		}
	} else {
		response, err := intentConfig.OnComplete(requestBody.Request.Intent)
		if err != nil {
			handleError(w, err)
			return
		}
		responseBody.Response = response
	}

	resp, err := json.Marshal(responseBody)

	if err != nil {
		handleError(w, err)
		return
	}

	w.Write(resp)
}

func Start(listen string, cfg *config.Config) {
	globalConfig = cfg

	h := http.HandlerFunc(handler)

	log.Printf("Starting web server, listening on %s", listen)

	http.ListenAndServe(listen, h)
}
