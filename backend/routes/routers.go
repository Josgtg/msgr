package routes

import (
	"log/slog"
	"msgr/controller"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/httplog/v2"
)

func CreateRouter() chi.Router {
	logger := httplog.NewLogger("httplog-example", httplog.Options{
		Concise:          true,
		LogLevel:         slog.LevelDebug,
		MessageFieldName: "message",
	})

	router := chi.NewRouter()

	router.Use(httplog.RequestLogger(logger))

	// Middleware to respond to OPTIONS and avoid CORS problems
	router.Use(func(h http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if r.Method == "OPTIONS" {
				controller.RespondJSON(w, http.StatusOK, nil)
			}
			h.ServeHTTP(w, r)
		})
	})

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
