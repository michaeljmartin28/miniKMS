package httpsrv

import (
	"log"
	"net/http"
)

func LoggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("%s %s Body: %+v", r.Method, r.URL.Path, r.Body)

		next.ServeHTTP(w, r)
	})
}

func NewRouter(h *Handler) http.Handler {
	mux := http.NewServeMux()

	mux.HandleFunc("POST /v1/keys", h.CreateKey)
	mux.HandleFunc("POST /v1/keys/{id}/enable", h.EnableKey)
	mux.HandleFunc("POST /v1/keys/{id}/disable", h.DisableKey)
	mux.HandleFunc("POST /v1/keys/{id}/encrypt", h.Encrypt)
	mux.HandleFunc("POST /v1/keys/{id}/decrypt", h.Decrypt)
	mux.HandleFunc("POST /v1/keys/{id}/generate-data-key", h.GenerateDataKey)
	mux.HandleFunc("POST /v1/keys/{id}/decrypt-data-key", h.DecryptDataKey)
	mux.HandleFunc("POST /v1/keys/{id}/rotate", h.RotateKey)

	return LoggingMiddleware(mux)
}
