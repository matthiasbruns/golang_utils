package router

import (
	"encoding/json"
	"fmt"
	"matthiasbruns/golang_utils/env"
	"net/http"
)

func jsonError(w http.ResponseWriter, err string, code int) {
	jsonResponse(w, code, err)
}

func jsonSuccess(w http.ResponseWriter, body string) {
	jsonResponse(w, http.StatusOK, body)
}

func jsonResponse(w http.ResponseWriter, code int, json string) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.Header().Set("X-Content-Type-Options", "nosniff")
	w.WriteHeader(code)
	_, _ = fmt.Fprintf(w, json)
}

func tryMarshalOr500(w http.ResponseWriter, input interface{}) *[]byte {
	js, err := json.Marshal(input)
	if err != nil {
		respondWithError(w, err, http.StatusInternalServerError, "Could not marshal json")
		return nil
	}
	return &js
}

func respondWithError(w http.ResponseWriter, err error, httpError int, message string) {
	var devMsg *string
	if env.IsDev() && err != nil {
		m := err.Error()
		devMsg = &m
	}

	errResponse := ErrorResponse{
		Message:      message,
		DebugMessage: devMsg,
	}

	js := tryMarshalOr500(w, errResponse)
	if js == nil {
		return
	}

	jsonError(w, string(*js), httpError)
}

