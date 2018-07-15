package utils

import (
	"encoding/json"
	"errors"
	"net/http"
)

// Request2Struct
func Request2Struct(w http.ResponseWriter, r *http.Request, req interface{}) {
	reqBody := r.Context().Value("reqBody")

	body, ok := reqBody.([]byte)
	if !ok {
		RespMsg(w, r, errors.New("Assert Error"))
		return
	}

	err := json.Unmarshal(body, req)
	if err != nil {
		RespMsg(w, r, err)
		return
	}
}
