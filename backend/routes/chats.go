package routes

import (
	"msgr/controller"

	"github.com/go-chi/chi/v5"
)

func chatRouter() chi.Router {
	router := chi.NewRouter()

	// FIXME: Must pass through a middleware to check admin
	router.Get("/", controller.GetAllChats)

	router.Get("/{id}", controller.GetChat)
	router.Post("/", controller.InsertChat)
	router.Delete("/{id}", controller.DeleteChat)

	router.Get("/user/{id}", controller.GetUserChats)

	return router
}
