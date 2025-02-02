package handlers

import (
	"Go-Tutorial-Api/internal/middleware"
	"github.com/go-chi/chi"
	chimiddle "github.com/go-chi/chi/middleware"
)

func Handler(r *chi.Mux) {
	// Global middleware
	r.Use(chimiddle.StripSlashes) // This enables the removal of trailing slashes from incoming requests to avoid 404 errors
	r.Route("/account", func(router chi.Router) {
		router.Use(middleware.Authorization)
		router.Get("/coins", GetCoinBalance)
	})
}
