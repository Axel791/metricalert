package main

import (
	"flag"
	"fmt"
	"github.com/Axel791/metricalert/internal/server/handlers"
	"github.com/Axel791/metricalert/internal/server/storage/repositories"
	"github.com/Axel791/metricalert/internal/shared/validatiors"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"net/http"
)

func main() {
	addr := flag.String("a", "localhost:8080", "HTTP server address (default: localhost:8080)")

	flag.Parse()

	if !validatiors.IsValidAddress(*addr) {
		fmt.Errorf("Invalid address: %s\n", *addr)
		return
	}

	router := chi.NewRouter()

	router.Use(middleware.Logger)
	router.Use(middleware.Recoverer)

	storage := repositories.NewMetricRepository()

	router.Method(
		http.MethodPost, "/update/{metricType}/{name}/{value}", handlers.NewUpdateMetricHandler(storage),
	)
	router.Method(http.MethodGet, "/value/{metricType}/{name}", handlers.NewGetMetricHandler(storage))

	err := http.ListenAndServe(*addr, router)
	if err != nil {
		panic(err)
	}
}
