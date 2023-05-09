package jsend

import (
	"encoding/json"
	"net/http"
	"strconv"
)

func Error(w http.ResponseWriter, message string) {
	w.WriteHeader(500)
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte("{\"status\": \"error\", \"message\": " + string(message) + "}"))
}

func ErrorWithCode(w http.ResponseWriter, code int, message string) {
	w.WriteHeader(500)
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte("{\"status\": \"error\", \"message\": " + string(message) + ", \"code\":" + strconv.Itoa(code) + "}"))
}

func ErrorWithData(w http.ResponseWriter, message string, data map[string]interface{}) {
	jdata, err := json.Marshal(data)
	if err != nil {
		jdata = []byte("null")
	}
	w.WriteHeader(500)
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte("{\"status\": \"error\", \"message\": " + string(message) + ", \"data\":" + string(jdata) + "}"))
}
