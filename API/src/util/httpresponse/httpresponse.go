package httpresponse

import (
	"encoding/json"
	"log"
	"net/http"
)

// JSON return a HTTP response as json
func JSON(writer http.ResponseWriter, status int, model interface{}) {
	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(status)
	err := json.NewEncoder(writer).Encode(model)
	if err != nil {
		log.Fatal("Error writing response", err)
	}
}

// Error return an error as text
func Error(writer http.ResponseWriter, status int, message string) {
	writer.Header().Set("Content-Type", "text/plain")
	writer.WriteHeader(status)
	_, err := writer.Write([]byte(message))
	if err != nil {
		log.Fatal("Error writing response", err)
	}
}
