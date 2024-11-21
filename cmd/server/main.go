package main

import (
	"github.com/Axel791/metricalert/internal/handlers"
	"github.com/Axel791/metricalert/internal/storage/repository"
	"net/http"
)

func main() {
	mux := http.NewServeMux()

	storage := repository.NewMetricRepository()

	mux.Handle("/update/", handlers.NewUpdateMetricHandler(storage))

	err := http.ListenAndServe(":8080", mux)
	if err != nil {
		panic(err)
	}
}
