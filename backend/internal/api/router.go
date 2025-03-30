package api

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"

	"github.com/awnzl/to-do-app/internal/api/handlers/lists"
	"github.com/awnzl/to-do-app/internal/api/handlers/todos"
	"github.com/awnzl/to-do-app/internal/service"
)

func NewRouter(svc service.TodoService) *chi.Mux {
	r := chi.NewRouter()

	// Middleware
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.RequestID)

	// API routes
	r.Route("/api/v1", func(r chi.Router) {
		// Lists endpoints
		r.Route("/lists", func(r chi.Router) {
			listsHandler := lists.NewHandler(svc)
			listsHandler.RegisterRoutes(r)

			// Nested todos endpoints
			r.Route("/{listID}/todos", func(r chi.Router) {
				todosHandler := todos.NewHandler(svc)
				todosHandler.RegisterCreateRoute(r)
				todosHandler.RegisterRoutes(r)
			})
		})

		// Individual todo endpoints
		r.Route("/todos", func(r chi.Router) {
			todosHandler := todos.NewHandler(svc)
			todosHandler.RegisterRoutes(r)
		})
	})

	return r
}
