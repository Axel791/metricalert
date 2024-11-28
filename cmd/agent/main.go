package main

import (
	"fmt"
	"github.com/Axel791/metricalert/internal/agent/collector"
	"github.com/Axel791/metricalert/internal/agent/model/dto"
	"github.com/Axel791/metricalert/internal/agent/sender"
	"time"
)

const (
	pollInterval   = time.Second * 2
	reportInterval = time.Second * 10
	requestURL     = "http://0.0.0.0"
	requestPort    = 8080
)

func main() {
	var metricsDTO dto.Metrics

	tickerCollector := time.NewTicker(pollInterval)
	tickerSender := time.NewTicker(reportInterval)

	metricClient := sender.NewMetricClient(requestURL, requestPort)

	defer tickerCollector.Stop()
	defer tickerSender.Stop()

	for {
		select {
		case <-tickerCollector.C:
			metric := collector.Collector()

			metricsDTO = dto.Metrics{
				Alloc:         float64(metric.Alloc) / 1024,
				BuckHashSys:   float64(metric.BuckHashSys) / 1024,
				Frees:         float64(metric.Frees),
				GCCPUFraction: metric.GCCPUFraction,
			}
			fmt.Println("Собранные метрики:", metricsDTO)

		case <-tickerSender.C:
			fmt.Println("Отправка метрик", metricsDTO)

			err := metricClient.SendMetrics(metricsDTO)
			if err != nil {
				_ = fmt.Errorf("error: %v", err)
			}
		}
	}
}
