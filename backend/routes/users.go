package routes

import (
	"msgr/controller"
	"msgr/middleware"

	"github.com/go-chi/chi/v5"
)

func userRouter() chi.Router {
	router := chi.NewRouter()

	router.Post("/register", controller.Register)
	router.Post("/login", controller.LogIn)

	router.Group(func(r chi.Router) {
		r.Use(middleware.Admin)
		r.Get("/", controller.GetAllUsers)
	})

	router.Group(func(r chi.Router) {
		r.Use(middleware.SameUserID)
		r.Get("/{id}", controller.GetUser)
		r.Get("/{id}/chats", controller.GetUserChats)
		r.Delete("/{id}", controller.DeleteUser)
	})

	return router
}
