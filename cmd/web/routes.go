package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
)

func (app *application) routes() http.Handler {

	mux := chi.NewRouter()

	mux.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"https://*", "http://*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		AllowCredentials: false,
		MaxAge:           300,
	}))

	mux.Post("/create", app.CreateMenu)
	mux.Get("/get-all-menu", app.getAllMenu)
	mux.Get("/get-opened-menu", app.getOpenMenu)
	mux.Post("/open-menu", app.OpenMenu)
	mux.Post("/update-menu", app.updateMenu)
	mux.Post("/update-rating", app.updateMenuRating)
	mux.Post("/add-order", app.addOrder)
	mux.Post("/get-order", app.getAllOrder)
	mux.Post("/update-order", app.updateOrder)
	mux.Post("/delete-order", app.deleteOrder)

	return mux
}
