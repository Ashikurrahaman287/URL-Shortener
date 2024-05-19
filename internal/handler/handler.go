// internal/handler/handler.go
package handler

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com//Ashikurrahaman287//url-shortener/internal/model"
	"github.com//Ashikurrahaman287//url-shortener/internal/store"
	"github.com/go-chi/chi/v5"
)

type Handler struct {
	Store *store.URLStore
}

func NewHandler(store *store.URLStore) *Handler {
	return &Handler{Store: store}
}

func (h *Handler) CreateShortURL(w http.ResponseWriter, r *http.Request) {
	var request struct {
		Original string `json:"original"`
		Expiry   int64  `json:"expiry"`
	}

	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	short := generateShortURL(request.Original) // Implement this function
	url := &model.URL{
		ID:        short,
		Original:  request.Original,
		Short:     short,
		CreatedAt: time.Now(),
		ExpiresAt: time.Now().Add(time.Duration(request.Expiry) * time.Second),
		Clicks:    0,
	}

	h.Store.Save(url)
	response := map[string]string{"short": short}
	json.NewEncoder(w).Encode(response)
}

func (h *Handler) Redirect(w http.ResponseWriter, r *http.Request) {
	short := chi.URLParam(r, "short")
	url, exists := h.Store.Find(short)
	if !exists || time.Now().After(url.ExpiresAt) {
		http.Error(w, "URL not found or expired", http.StatusNotFound)
		return
	}

	url.Clicks++
	http.Redirect(w, r, url.Original, http.StatusFound)
}

// Placeholder for generating short URL. Replace with your implementation.
func generateShortURL(original string) string {
	return "shortURL"
}
