package main

import (
	"log"
	"net/http"
	"os"
	"encoding/json"
)

// HTTPResponse is the expected outgoing payload to a HTTP Trigger
// to function host.
type HTTPResponse struct {
	Outputs struct {
		Res struct {
			Body       string `json:"body"`
			StatusCode string `json:"statusCode"`
		} `json:"res"`
	}
	Logs        []string
	ReturnValue interface{}
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	log.Printf("Handler was called...\n")

	response := HTTPResponse{}
	response.Outputs.Res.StatusCode = "400"
	res, err := json.Marshal(response)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(res)
}

func main() {
	listenAddr := ":8080"
	if val, ok := os.LookupEnv("FUNCTIONS_CUSTOMHANDLER_PORT"); ok {
		listenAddr = ":" + val
	}
	http.HandleFunc("/TimerTrigger1", helloHandler)
	log.Printf("About to listen on %s. Go to https://127.0.0.1%s/", listenAddr, listenAddr)
	log.Fatal(http.ListenAndServe(listenAddr, nil))
}
