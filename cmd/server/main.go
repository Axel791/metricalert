package main

import (
	handlers2 "github.com/Axel791/metricalert/internal/server/handlers"
	"github.com/Axel791/metricalert/internal/server/storage/repositories"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"net/http"
)

func main() {
	router := chi.NewRouter()

	router.Use(middleware.Logger)
	router.Use(middleware.Recoverer)

	storage := repositories.NewMetricRepository()

	router.Method(
		http.MethodPost, "/update/{metricType}/{name}/{value}", handlers2.NewUpdateMetricHandler(storage),
	)
	router.Method(http.MethodGet, "/value/{metricType}/{name}", handlers2.NewGetMetricHandler(storage))

	err := http.ListenAndServe(":8080", router)
	if err != nil {
		panic(err)
	}
}
