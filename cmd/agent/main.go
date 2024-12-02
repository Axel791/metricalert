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
		fmt.Printf("error loading config: %v\n", err)
		return
	}

	address := flag.String("a", cfg.Address, "HTTP server address")
	reportInterval := flag.Int("r", int(cfg.ReportInterval.Seconds()), "Frequency of sending metrics to the server (in seconds)")
	pollInterval := flag.Int("p", int(cfg.PollInterval.Seconds()), "Frequency of collecting metrics from runtime (in seconds)")

	flag.Parse()

	fmt.Printf("Server address: %s\n", *address)
	fmt.Printf("Report interval: %s\n", *address)
	fmt.Printf("Polling interval: %s\n", *pollInterval)

	if !validatiors.IsValidAddress(*address, true) {
		fmt.Printf("invalid address: %s\n", *address)
		return
	}

	reportDuration := time.Duration(*reportInterval) * time.Second
	pollDuration := time.Duration(*pollInterval) * time.Second

	var metricsDTO dto.Metrics

	tickerCollector := time.NewTicker(reportDuration)
	tickerSender := time.NewTicker(pollDuration)

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
				fmt.Printf("error sending metrics: %v\n", err)
			}
		}
	}
}
