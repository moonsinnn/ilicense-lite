package otel

import (
	"log"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/exporters/prometheus"
	"go.opentelemetry.io/otel/exporters/stdout/stdoutmetric"
	"go.opentelemetry.io/otel/sdk/metric"
)

func InitMetric() {
	stdoutExporter, err := stdoutmetric.New(stdoutmetric.WithPrettyPrint())
	if err != nil {
		log.Fatalf("failed to initialize stdoutmetric promExporter: %v", err)
	}
	// Prometheus promExporter
	promExporter, err := prometheus.New()
	if err != nil {
		log.Fatalf("init prometheus promExporter err: %v", err)
	}
	provider := metric.NewMeterProvider(
		metric.WithReader(metric.NewPeriodicReader(stdoutExporter)),
		metric.WithReader(promExporter),
	)
	otel.SetMeterProvider(provider)
}
