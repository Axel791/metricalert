package main

import (
	"flag"
	"fmt"
	"github.com/Axel791/metricalert/internal/agent/collector"
	"github.com/Axel791/metricalert/internal/agent/config"
	"github.com/Axel791/metricalert/internal/agent/model/dto"
	"github.com/Axel791/metricalert/internal/agent/sender"
	"github.com/Axel791/metricalert/internal/shared/validatiors"
	"time"
)

func main() {
	cfg, err := config.AgentLoadConfig()
	if err != nil {
		fmt.Printf("error loading config: %v", err)
		return
	}

	address := flag.String("a", fmt.Sprintf("%s", cfg.Address), "HTTP server address")
	reportInterval := flag.Duration("r", cfg.ReportInterval, "Frequency of sending metrics to the server")
	pollInterval := flag.Duration("p", cfg.PollInterval, "Frequency of collecting metrics from runtime")

	flag.Parse()

	if !validatiors.IsValidAddress(*address) {
		_ = fmt.Errorf("Invalid address: %s\n", *address)
		return
	}

	var metricsDTO dto.Metrics

	tickerCollector := time.NewTicker(*pollInterval)
	tickerSender := time.NewTicker(*reportInterval)

	metricClient := sender.NewMetricClient(*address)

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
				fmt.Printf("Error sending metrics: %v\n", err)
			}
		}
	}
}
