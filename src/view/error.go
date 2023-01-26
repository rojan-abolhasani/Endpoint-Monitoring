package view

import (
	"encoding/json"
	"monitor/model"
	"net/http"
)

type ErrHandle model.ErrResponse

// we create a handler from the error response
func (err ErrHandle) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(err.ErrorCode)
	json.NewEncoder(w).Encode(err)
}
