package routes

import (
	"msgr/controller"
	"msgr/middleware"

	"github.com/go-chi/chi/v5"
)

func chatRouter() chi.Router {
	router := chi.NewRouter()

	router.Group(func(r chi.Router) {
		r.Use(middleware.Admin)
		r.Get("/", controller.GetAllChats)
	})

	router.Group(func(r chi.Router) {
		r.Use(middleware.CheckSession)
		r.Get("/{id}", controller.GetChat)
		r.Post("/", controller.InsertChat)
		r.Delete("/{id}", controller.DeleteChat)
		r.Get("/{id}/messages", controller.GetMessagesByChat)
	})

	return router
}
