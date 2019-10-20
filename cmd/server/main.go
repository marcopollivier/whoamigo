package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
	"time"
)

func main() {
	Init()
}

func Init() {
	http.HandleFunc("/", meHandler)
	log.Fatal(http.ListenAndServe(":8081", nil))
}

func now() string {
	var timeNow = time.Now()
	return timeNow.Format("2006-01-02T15:04:05.000Z")
}

func meHandler(responseWriter http.ResponseWriter, req *http.Request) {
	responseWriter.Header().Set("Content-Type", "application/json")

	var hostname, _ = os.Hostname()
	var payload = Me{hostname, now()}

	returnBody, _ := json.Marshal(payload)

	_, _ = responseWriter.Write(returnBody)
}

type Me struct {
	Me        string `json:"me"`
	Timestamp string `json:"timestamp"`
}
