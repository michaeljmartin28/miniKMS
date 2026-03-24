package httpsrv

import (
	"encoding/json"
	"errors"
	"net/http"
)

type ErrorResponse struct {
	Error string `json:"error"`
}

var ErrBadJson = errors.New("invalid json in request")
var ErrMissingFields = errors.New("missing fields in request")
var ErrBadAlgorithm = errors.New("unsupported algorithm chosen")

func WriteError(w http.ResponseWriter, err error) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusFromError(err))
	json.NewEncoder(w).Encode(ErrorResponse{Error: err.Error()})
}

func statusFromError(err error) int {
	switch {

	case errors.Is(err, ErrBadJson),
		errors.Is(err, ErrMissingFields),
		errors.Is(err, ErrBadAlgorithm):
		return http.StatusBadRequest

	default:
		return http.StatusInternalServerError
	}
}
