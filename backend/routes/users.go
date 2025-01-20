package routes

import (
	"msgr/controller"

	"github.com/go-chi/chi/v5"
)

func userRouter() chi.Router {
	router := chi.NewRouter()

	// FIXME: Must pass through a middleware to check admin
	router.Get("/", controller.GetAllUsers)

	router.Get("/{id}", controller.GetUser)
	router.Post("/", controller.InsertUser)
	router.Delete("/{id}", controller.DeleteUser)

	return router
}
