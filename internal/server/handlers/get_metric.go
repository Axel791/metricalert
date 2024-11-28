package handlers

import (
	"fmt"
	"github.com/Axel791/metricalert/internal/server/storage"
	"github.com/go-chi/chi/v5"
	"net/http"
)

type GetMetricHandler struct {
	storage storage.Store
}

func NewGetMetricHandler(storage storage.Store) *GetMetricHandler {
	return &GetMetricHandler{storage}
}

func (h *GetMetricHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	metricType := chi.URLParam(r, "metricType")
	name := chi.URLParam(r, "name")

	if name == "" {
		http.Error(w, "invalid metric name", http.StatusNotFound)
		return
	}

	var value interface{}

	if metricType == Counter || metricType == Gauge {
		value = h.storage.GetMetric(name)
	} else {
		http.Error(w, "invalid metric type", http.StatusBadRequest)
		return
	}

	if value == nil {
		http.Error(w, "metric not found", http.StatusNotFound)
	}

	valueStr, ok := value.(string)
	if !ok {
		valueStr = fmt.Sprintf("%v", value)
	}

	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	w.Header().Set("Content-Length:", valueStr)
	w.WriteHeader(http.StatusOK)

}
