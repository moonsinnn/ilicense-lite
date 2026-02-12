package otel

import (
	"go.opentelemetry.io/otel/sdk/resource"
	semconv "go.opentelemetry.io/otel/semconv/v1.4.0"
)

func InitResource() (*resource.Resource, error) {
	return resource.NewWithAttributes(
		semconv.SchemaURL,
		semconv.ServiceNameKey.String("ilicense-lite"),
		semconv.ServiceVersionKey.String("1.0"),
		semconv.DeploymentEnvironmentKey.String("dev"),
		semconv.CloudRegionKey.String("zh-cn"),
	), nil
}
