package router

import (
	"encoding/json"
	"fmt"
	"github.com/matthiasbruns/golang_utils/env"
	"net/http"
)

func JsonError(w http.ResponseWriter, err string, code int) {
	JsonResponse(w, code, err)
}

func JsonSuccess(w http.ResponseWriter, body string) {
	JsonResponse(w, http.StatusOK, body)
}

func JsonResponse(w http.ResponseWriter, code int, json string) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.Header().Set("X-Content-Type-Options", "nosniff")
	w.WriteHeader(code)
	_, _ = fmt.Fprintf(w, json)
}

func TryMarshalOr500(w http.ResponseWriter, input interface{}) *[]byte {
	js, err := json.Marshal(input)
	if err != nil {
		RespondWithError(w, err, http.StatusInternalServerError, "Could not marshal json")
		return nil
	}
	return &js
}

func RespondWithError(w http.ResponseWriter, err error, httpError int, message string) {
	var devMsg *string
	if env.IsDev() && err != nil {
		m := err.Error()
		devMsg = &m
	}

	errResponse := ErrorResponse{
		Message:      message,
		DebugMessage: devMsg,
	}

	js := TryMarshalOr500(w, errResponse)
	if js == nil {
		return
	}

	JsonError(w, string(*js), httpError)
}
