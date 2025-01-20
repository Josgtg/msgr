package routes

import (
	"msgr/controller"

	"github.com/go-chi/chi/v5"
)

func messageRouter() chi.Router {
	router := chi.NewRouter()

	// FIXME: Must pass through a middleware to check admin
	router.Get("/", controller.GetAllMessages)

	router.Get("/{id}", controller.GetMessage)
	router.Post("/", controller.InsertMessage)
	router.Delete("/{id}", controller.DeleteMessage)

	return router
}
