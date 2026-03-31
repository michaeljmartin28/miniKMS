package httpsrv

import (
	"encoding/json"
	"io"
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

func DecodeRequest[T any](r io.Reader) (T, error) {

	var value T
	decoder := json.NewDecoder(r)
	if err := decoder.Decode(&value); err != nil {
		return value, err
	}
	return value, nil

}

func (h *Handler) CreateKey(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	req, err := DecodeRequest[CreateKeyRequest](r.Body)
	if err != nil {
		WriteError(w, ErrBadJSON)
		return
	}

	if req.Name == "" || req.Algorithm == "" {
		WriteError(w, ErrMissingFields)
		return
	}

	algorithm, err := core.ParseAlgorithm(req.Algorithm)
	if err != nil {
		WriteError(w, err)
		return
	}

	coreRequest := core.CreateKeyRequest{Name: req.Name, Algorithm: algorithm}
	response, err := h.Engine.CreateKey(ctx, coreRequest)
	if err != nil {
		WriteError(w, err)
		return
	}
	WriteJSON(w, http.StatusCreated, response)
}

func (h *Handler) EnableKey(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	keyID := r.PathValue("id")

	meta, err := h.Engine.EnableKey(ctx, keyID)
	if err != nil {
		WriteError(w, err)
		return
	}

	WriteJSON(w, http.StatusOK, meta)
}
func (h *Handler) DisableKey(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	keyID := r.PathValue("id")

	meta, err := h.Engine.DisableKey(ctx, keyID)
	if err != nil {
		WriteError(w, err)
		return
	}

	WriteJSON(w, http.StatusOK, meta)
}
func (h *Handler) Encrypt(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	keyID := r.PathValue("id")
	req, err := DecodeRequest[EncryptRequest](r.Body)
	if err != nil {
		WriteError(w, ErrBadJSON)
		return
	}

	encryptRequest := core.EncryptRequest{
		KeyID:          keyID,
		Plaintext:      []byte(req.Plaintext),
		AdditionalData: []byte(req.AdditionalData),
	}

	resp, err := h.Engine.Encrypt(ctx, encryptRequest)
	if err != nil {
		WriteError(w, err)
		return
	}
	WriteJSON(w, http.StatusOK, resp)

}
func (h *Handler) Decrypt(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	keyID := r.PathValue("id")

	req, err := DecodeRequest[DecryptRequest](r.Body)
	if err != nil {
		WriteError(w, ErrBadJSON)
		return
	}

	decryptRequest := core.DecryptRequest{
		KeyID:          keyID,
		Ciphertext:     []byte(req.Ciphertext),
		AdditionalData: []byte(req.AdditionalData),
		Version:        req.Version,
	}

	resp, err := h.Engine.Decrypt(ctx, decryptRequest)
	if err != nil {
		WriteError(w, err)
		return
	}
	WriteJSON(w, http.StatusOK, resp)
}
func (h *Handler) GenerateDataKey(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	keyID := r.PathValue("id")

	req, err := DecodeRequest[GenerateDataKeyRequest](r.Body)
	if err != nil {
		WriteError(w, ErrBadJSON)
		return
	}

	genDEKRequest := core.GenerateDataKeyRequest{
		KeyID:          keyID,
		AdditionalData: []byte(req.AdditionalData),
	}

	resp, err := h.Engine.GenerateDataKey(ctx, genDEKRequest)
	if err != nil {
		WriteError(w, err)
		return
	}
	WriteJSON(w, http.StatusOK, resp)
}
func (h *Handler) DecryptDataKey(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	keyID := r.PathValue("id")

	req, err := DecodeRequest[DecryptDataKeyRequest](r.Body)
	if err != nil {
		WriteError(w, ErrBadJSON)
		return
	}

	decryptDEKRequest := core.DecryptDataKeyRequest{
		KeyID:          keyID,
		EncryptedDEK:   []byte(req.EncryptedDEK),
		Version:        req.Version,
		AdditionalData: []byte(req.AdditionalData),
	}

	resp, err := h.Engine.DecryptDataKey(ctx, decryptDEKRequest)
	if err != nil {
		WriteError(w, err)
		return
	}
	WriteJSON(w, http.StatusOK, resp)
}
func (h *Handler) RotateKey(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	keyID := r.PathValue("id")

	version, err := h.Engine.RotateKey(ctx, keyID)
	if err != nil {
		WriteError(w, err)
		return
	}

	resp := RotateKeyResponse{Version: version}

	WriteJSON(w, http.StatusOK, resp)
}
