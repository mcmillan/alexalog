package server

import (
	"encoding/json"
	"net/http"

	"github.com/mcmillan/alexalog/domain"
	"github.com/mcmillan/alexalog/shortcuts"
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

	responseBody := shortcuts.SpeechResponse("hi there!", true)

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
