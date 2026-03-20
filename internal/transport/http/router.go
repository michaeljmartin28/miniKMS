package http

import "net/http"

func NewRouter(h *Handler) http.Handler {
	mux := http.NewServeMux()

	mux.HandleFunc("POST /keys", h.CreateKey)
	mux.HandleFunc("POST /keys/{id}/enable", h.EnableKey)
	mux.HandleFunc("POST /keys/{id}/disable", h.DisableKey)
	mux.HandleFunc("POST /keys/{id}/encrypt", h.Encrypt)
	mux.HandleFunc("POST /keys/{id}/decrypt", h.Decrypt)
	mux.HandleFunc("POST /keys/{id}/generate-data-key", h.GenerateDataKey)
	mux.HandleFunc("POST /keys/{id}/decrypt-data-key", h.DecryptDataKey)
	mux.HandleFunc("POST /keys/{id}/rotate", h.RotateKey)

	return mux
}
