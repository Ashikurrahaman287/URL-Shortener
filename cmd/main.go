// cmd/main.go
package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/your-username/url-shortener/internal/handler"
	"github.com/your-username/url-shortener/internal/store"
)

func main() {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	urlStore := store.NewURLStore()
	h := handler.NewHandler(urlStore)

	r.Post("/shorten", h.CreateShortURL)
	r.Get("/{short}", h.Redirect)

	http.ListenAndServe(":8080", r)
}
