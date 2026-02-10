package otel

import (
	"context"
	"log"

	"go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracegrpc"

	"github.com/sirupsen/logrus"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/sdk/trace"
)

func InitTracer() {
	rce, err := InitResource()
	if err != nil {
		log.Fatalf("init resource err: %v", err)
	}
	// console exporter
	// consoleExporter, err := stdouttrace.New(
	// 	stdouttrace.WithPrettyPrint(),
	// )
	if err != nil {
		logrus.Fatalf("failed to initialize stdouttrace consoleExporter: %v", err)
	}
	// jaeger exporter Jaeger HTTP 老方式，非 OTLP，OpenTelemetry 支持较差
	// jaegerExporter, err := jaeger.New(
	// 	jaeger.WithCollectorEndpoint(
	// 		jaeger.WithEndpoint("http://localhost:14268/api/traces"),
	// 	),
	// )
	// OTLP gRPC 推荐：OpenTelemetry 原生方式
	otlpExporter, err := otlptracegrpc.New(
		context.Background(),
		otlptracegrpc.WithEndpoint("localhost:4317"),
		otlptracegrpc.WithInsecure(), // Jaeger 本地默认没有 TLS
	)
	if err != nil {
		logrus.Fatalf("failed to initialize jaeger exporter: %v", err)
	}
	// 创建多个BatchSpanProcessor
	traceProvider := trace.NewTracerProvider(
		// trace.WithBatcher(consoleExporter),
		// trace.WithBatcher(jaegerExporter),
		trace.WithBatcher(otlpExporter),
		trace.WithResource(rce),
		trace.WithSampler(trace.AlwaysSample()),
	)
	otel.SetTracerProvider(traceProvider)
	log.Println("otel tracer initialized")
}
