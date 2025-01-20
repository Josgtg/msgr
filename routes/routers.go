package routes

import (
	"msgr/controller"

	"github.com/go-chi/chi/v5"
)

func CreateRouter() chi.Router {
	router := chi.NewRouter()

	router.Get("/health", controller.Health)

	router.Mount("/api", apiRouter())

	return router
}

func apiRouter() chi.Router {
	router := chi.NewRouter()

	router.NotFound(controller.NotFound)
	router.MethodNotAllowed(controller.MethodNotAllowed)

	router.Mount("/users", userRouter())
	router.Mount("/chats", chatRouter())
	router.Mount("/messages", messageRouter())

	return router
}
