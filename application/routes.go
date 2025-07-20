package application

import (
	"net/http"

	"github.com/MustafaPatharia/go-microservice/handler"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func loadRoutes() *chi.Mux {
	router := chi.NewRouter()

	router.Use(middleware.Logger)

	router.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	})

	router.Route("/orders", loadOrdersRotes)

	return router
}

func loadOrdersRotes(r chi.Router) {
	orderHandler := &handler.Order{}

	r.Post("/", orderHandler.Create)
	r.Get("/", orderHandler.GetList)
	r.Get("/{id}", orderHandler.GetById)
	r.Put("/{id}", orderHandler.Update)
	r.Delete("/{id}", orderHandler.Delete)
}
