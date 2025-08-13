package helper

import (
	"encoding/json"
	"net/http"
)

func WithJSON(resW http.ResponseWriter, code int, payload interface{}){
resW.Header().Set("Content-Type", "application/json")
resW.WriteHeader(code)
json.NewEncoder(resW).Encode(payload)
}

func WithError(resW http.ResponseWriter, code int, message string){
	WithJSON(resW,code, message)
}