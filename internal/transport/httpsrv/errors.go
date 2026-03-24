package httpsrv

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/michaeljmartin28/minikms/internal/core"
)

type ErrorResponse struct {
	Error string `json:"error"`
}

var (
	ErrBadJSON       = errors.New("invalid json")
	ErrMissingFields = errors.New("missing required fields")
	ErrBadMethod     = errors.New("invalid http method")
)

func WriteError(w http.ResponseWriter, err error) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusFromError(err))
	json.NewEncoder(w).Encode(ErrorResponse{Error: err.Error()})
}

func statusFromError(err error) int {
	switch {
	case errors.Is(err, ErrBadJSON),
		errors.Is(err, ErrMissingFields),
		errors.Is(err, core.ErrBadAlgorithm),
		errors.Is(err, core.ErrInvalidVersion):
		return http.StatusBadRequest
	case errors.Is(err, ErrBadMethod):
		return http.StatusMethodNotAllowed
	case errors.Is(err, core.ErrKeyNotFound):
		return http.StatusNotFound
	case errors.Is(err, core.ErrKeyDisabled):
		return http.StatusConflict
	default:
		return http.StatusInternalServerError
	}
}
