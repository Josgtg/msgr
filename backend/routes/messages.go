package routes

import (
	"msgr/controller"
	"msgr/middleware"

	"github.com/go-chi/chi/v5"
)

func messageRouter() chi.Router {
	router := chi.NewRouter()

	router.Group(func(r chi.Router) {
		r.Use(middleware.Admin)
		r.Get("/", controller.GetAllMessages)
	})

	router.Group(func(r chi.Router) {
		r.Use(middleware.CheckSession)
		r.Get("/{id}", controller.GetMessage)
		r.Post("/", controller.InsertMessage)
		r.Delete("/{id}", controller.DeleteMessage)
	})

	return router
}
