package util

import (
	"context"
	"go.opentelemetry.io/otel/trace"
)

func GetTraceInfo(ctx context.Context) (string, string) {
	spanCtx := trace.SpanContextFromContext(ctx)
	if spanCtx.IsValid() {
		return spanCtx.TraceID().String(), spanCtx.SpanID().String()
	}
	return "", ""
}
