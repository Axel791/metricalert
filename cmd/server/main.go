package main

import (
	"flag"
	"fmt"
	"github.com/Axel791/metricalert/internal/server/config"
	"github.com/Axel791/metricalert/internal/server/handlers"
	"github.com/Axel791/metricalert/internal/server/storage/repositories"
	"github.com/Axel791/metricalert/internal/shared/validatiors"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"net/http"
)

func main() {
	cfg, err := config.ServerLoadConfig()
	if err != nil {
		fmt.Printf("error loading config: %v", err)
	}

	fmt.Printf("Server address: %s\n", cfg.Address)

	addr := flag.String("a", cfg.Address, "HTTP server address (default: localhost:8080)")

	flag.Parse()

	fmt.Printf("Server address 2: %s\n", *addr)

	if !validatiors.IsValidAddress(*addr, false) {
		fmt.Printf("invalid address: %s\n", *addr)
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
	err = http.ListenAndServe(*addr, router)
	fmt.Printf("Server started: %s\n", *addr)
	if err != nil {
		panic(err)
	}
}
