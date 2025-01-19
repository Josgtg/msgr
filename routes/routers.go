package routes

import (
	"msgr/controller"

	"github.com/go-chi/chi/v5"
)

func CreateRouter() chi.Router {
	router := chi.NewRouter()
	router.Get("/health", controller.Health)
	router.Mount("/users", UserRouter())

	return router
}
