package handlers

import (
	"github.com/Axel791/metricalert/internal/storage"
	"net/http"
	"testing"
)

func TestUpdateMetricHandler_ServeHTTP(t *testing.T) {
	type fields struct {
		storage storage.Store
	}
	type args struct {
		w http.ResponseWriter
		r *http.Request
	}
	var tests []struct {
		name   string
		fields fields
		args   args
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := &UpdateMetricHandler{
				storage: tt.fields.storage,
			}
			h.ServeHTTP(tt.args.w, tt.args.r)

		})
	}
}
