package http

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
	"time"
)

type HealthBody struct {
	Me        string `json:"me"`
	Timestamp string `json:"timestamp"`
}

func now() string {
	var timeNow = time.Now()
	return timeNow.Format("2006-01-02T15:04:05.000Z")
}

func meHandler(responseWriter http.ResponseWriter, req *http.Request) {
	responseWriter.Header().Set("Content-Type", "application/json")

	var hostname, _ = os.Hostname()
	profile := HealthBody{hostname, now()}

	returnBody, _ := json.Marshal(profile)

	_, _ = responseWriter.Write(returnBody)
}

func Init() {
	http.HandleFunc("/", meHandler)
	log.Fatal(http.ListenAndServe(":8081", nil))
}
