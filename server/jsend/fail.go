package jsend

import (
	"encoding/json"
	"net/http"
)

func Fail(w http.ResponseWriter, code int, data map[string]interface{}) {
	jdata, err := json.Marshal(data)
	if err != nil {
		jdata = []byte("null")
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write([]byte("{\"status\": \"fail\", \"data\": " + string(jdata) + "}"))
}
