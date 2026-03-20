package http

import (
	"encoding/json"
	"net/http"

	"github.com/michaeljmartin28/minikms/internal/core"
)

type Handler struct {
	Engine *core.Engine
}

func NewHandler(engine *core.Engine) *Handler {
	return &Handler{Engine: engine}
}

func WriteJSON(w http.ResponseWriter, status int, v any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(v)
}

func (h *Handler) CreateKey(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var request CreateKeyRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		WriteError(w, ErrBadJson)
		return
	}

	// log.Printf("request: %+v\n", request)

	if request.Name == "" || request.Algorithm == "" {
		WriteError(w, ErrMissingFields)
		return
	}

	algorithm, err := core.ParseAlgorithm(request.Algorithm)
	if err != nil {
		WriteError(w, err)
		return
	}

	coreRequest := core.CreateKeyRequest{Name: request.Name, Algorithm: algorithm}
	response, err := h.Engine.CreateKey(ctx, coreRequest)
	if err != nil {
		WriteError(w, err)
		return
	}

	WriteJSON(w, http.StatusCreated, response)
}

func (h *Handler) EnableKey(w http.ResponseWriter, r *http.Request)       {}
func (h *Handler) DisableKey(w http.ResponseWriter, r *http.Request)      {}
func (h *Handler) Encrypt(w http.ResponseWriter, r *http.Request)         {}
func (h *Handler) Decrypt(w http.ResponseWriter, r *http.Request)         {}
func (h *Handler) GenerateDataKey(w http.ResponseWriter, r *http.Request) {}
func (h *Handler) DecryptDataKey(w http.ResponseWriter, r *http.Request)  {}
func (h *Handler) RotateKey(w http.ResponseWriter, r *http.Request)       {}
