package routes

import (
	"log/slog"
	"msgr/controller"
	"msgr/middleware"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/httplog/v2"
)

func CreateRouter() chi.Router {
	router := chi.NewRouter()

	logger := httplog.NewLogger("server-logger", httplog.Options{
		Concise:          true,
		LogLevel:         slog.LevelDebug,
		MessageFieldName: "message",
	})
	router.Use(httplog.RequestLogger(logger))

	// Middleware to respond to OPTIONS and avoid CORS problems
	router.Use(middleware.Options)

	router.Mount("/api", apiRouter())

	return router
}

func apiRouter() chi.Router {
	router := chi.NewRouter()

	router.Get("/health", controller.Health)

	router.NotFound(controller.NotFound)
	router.MethodNotAllowed(controller.MethodNotAllowed)

	router.Mount("/users", userRouter())
	router.Mount("/chats", chatRouter())
	router.Mount("/messages", messageRouter())

	return router
}
