package server

import (
	"encoding/json"
	"net/http"

	"github.com/mcmillan/alexalog/domain"
	log "github.com/sirupsen/logrus"
)

func handler(w http.ResponseWriter, r *http.Request) {
	log.Printf("%s %s", r.Method, r.RequestURI)

	if r.Method != "POST" {
		http.Error(w, "method not allowed", 405)
		return
	}

	var requestBody domain.RequestBody

	err := json.NewDecoder(r.Body).Decode(&requestBody)

	if err != nil {
		log.Error(err)
		http.Error(w, err.Error(), 500)
		return
	}

	log.Printf("requestType: %s", requestBody.Request.Type)

	responseBody := domain.ResponseBody{
		Version: "1.0",
		Response: &domain.Response{
			OutputSpeech: &domain.ResponseOutputSpeech{
				Type: "PlainText",
				Text: "my roflcopter goes soi soi soi soi soi",
			},
			ShouldEndSession: true,
		},
	}

	resp, err := json.Marshal(responseBody)

	if err != nil {
		log.Error(err)
		http.Error(w, err.Error(), 500)
		return
	}

	log.Print(string(resp))

	w.Write(resp)
}

func Start(listen string) {
	h := http.HandlerFunc(handler)

	log.Printf("Starting web server, listening on %s", listen)

	http.ListenAndServe(listen, h)
}
