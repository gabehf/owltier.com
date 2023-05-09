package jsend

import (
	"encoding/json"
	"net/http"
)

// A response denoting a successful request.
// More information about a success response can be found at:
// https://github.com/omniti-labs/jsend
func Success(w http.ResponseWriter, data map[string]interface{}) {
	jdata, err := json.Marshal(data)
	if err != nil {
		jdata = []byte("null")
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("{\"status\": \"success\", \"data\": " + string(jdata) + "}"))
}
