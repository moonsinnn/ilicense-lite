package otel

import "ilicense-lite/library/otel"

func InitOTEL() {
	otel.InitTracer()
	otel.InitMetric()
}
