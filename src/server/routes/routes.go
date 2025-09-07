package routes

import "github.com/go-chi/chi/v5"

// Setup all the routes here
func SetupRoutes(r *chi.Mux) {
	r.Get("/healthz", Healthz)

	r.Route("/tasks", func(r chi.Router) {
		r.Post("/new", CreateNewTask)
		r.Get("/all", RetrieveAllTasks)
		r.Get("/{id}", GetSingleTask)
		r.Put("/{id}", UpdateTask)
		r.Delete("/{id}", RemoveTaskById)
	})
}
