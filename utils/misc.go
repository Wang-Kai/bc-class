package utils

import (
	"encoding/json"
	"net/http"
)

const (
	OK         = 200
	InterError = 500
)

type RespErrMsg struct {
	Code    uint   `json:"code"`
	Message string `json:"message"`
}

// RespMsg response body to http client
func RespMsg(w http.ResponseWriter, r *http.Request, msg interface{}) {
	switch m := msg.(type) {
	case error:
		respStruct(w, r, &RespErrMsg{Code: 500, Message: m.Error()})
	default:
		respStruct(w, r, m)
	}
}

func respStruct(w http.ResponseWriter, r *http.Request, obj interface{}) {
	bytes, err := json.Marshal(obj)
	if err != nil {
		respStruct(w, r, err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(bytes)
}
