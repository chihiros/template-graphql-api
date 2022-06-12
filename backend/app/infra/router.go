package infra

import (
	"tamaribacms/ent"
	"tamaribacms/interfaces/controller"

	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/chi/v5"
)

func NewRouter(conn *ent.Client) *chi.Mux {
	r := chi.NewRouter()
	r.Use(middleware.Logger)

	controller := controller.NewController(conn)
	r.Route("/api", func(r chi.Router) {
		r.Route("/users", func(r chi.Router) {
			r.Get("/", controller.Get)
			r.Get("/query", controller.GetByID)
			r.Post("/", controller.Post)
			r.Put("/", controller.Put)
			r.Delete("/", controller.Delete)
		})
	})

	return r
}
