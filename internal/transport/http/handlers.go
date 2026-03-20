package http

import (
	"net/http"

	"github.com/michaeljmartin28/minikms/internal/core"
)

type Handler struct {
	Engine *core.Engine
}

func NewHandler(engine *core.Engine) *Handler {
	return &Handler{Engine: engine}
}

func (h *Handler) CreateKey(w http.ResponseWriter, r *http.Request)       {}
func (h *Handler) EnableKey(w http.ResponseWriter, r *http.Request)       {}
func (h *Handler) DisableKey(w http.ResponseWriter, r *http.Request)      {}
func (h *Handler) Encrypt(w http.ResponseWriter, r *http.Request)         {}
func (h *Handler) Decrypt(w http.ResponseWriter, r *http.Request)         {}
func (h *Handler) GenerateDataKey(w http.ResponseWriter, r *http.Request) {}
func (h *Handler) DecryptDataKey(w http.ResponseWriter, r *http.Request)  {}
func (h *Handler) RotateKey(w http.ResponseWriter, r *http.Request)       {}
