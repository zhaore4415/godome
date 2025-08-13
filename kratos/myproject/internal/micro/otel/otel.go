package otel

import (
	"context"
	"time"

	"myproject/internal/proto/hello"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/exporters/otlp/otlpmetric/otlpmetricgrpc"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracegrpc"
	"go.opentelemetry.io/otel/propagation"
	"go.opentelemetry.io/otel/sdk/metric"
	"go.opentelemetry.io/otel/sdk/resource"
	"go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.26.0"
)

func SetupOTel(ctx context.Context, bs *hello.Bootstrap) error {
	// 设置资源。
	res, err := newResource(bs.Service.Name+"."+bs.Service.Env, bs.Service.Version)
	if err != nil {
		return err
	}

	// 设置传播器。
	prop := newPropagator()
	otel.SetTextMapPropagator(prop)

	// 设置跟踪提供程序。
	tracerProvider, err := newTraceProvider(ctx, res, bs.Otlp.Endpoint)
	if err != nil {
		return err
	} else {
		otel.SetTracerProvider(tracerProvider)
	}

	// 设置度量仪提供程序。
	meterProvider, err := newMeterProvider(ctx, res, bs.Otlp.Endpoint, 10)
	if err != nil {
		return err
	} else {
		otel.SetMeterProvider(meterProvider)
	}

	return nil
}

func newResource(serviceName, serviceVersion string) (*resource.Resource, error) {
	return resource.Merge(resource.Default(),
		resource.NewWithAttributes(semconv.SchemaURL,
			semconv.ServiceName(serviceName),
			semconv.ServiceVersion(serviceVersion),
		))
}

func newPropagator() propagation.TextMapPropagator {
	return propagation.NewCompositeTextMapPropagator(
		propagation.TraceContext{},
		propagation.Baggage{},
	)
}

func newTraceProvider(ctx context.Context, res *resource.Resource, ep string) (*trace.TracerProvider, error) {
	traceExporter, err := otlptracegrpc.New(ctx, otlptracegrpc.WithInsecure(), otlptracegrpc.WithEndpoint(ep))
	if err != nil {
		return nil, err
	}

	traceProvider := trace.NewTracerProvider(
		trace.WithResource(res),
		trace.WithBatcher(traceExporter, trace.WithBatchTimeout(time.Second*5)),
	)
	return traceProvider, nil
}

func newMeterProvider(ctx context.Context, res *resource.Resource, ep string, intervalSecond int64) (*metric.MeterProvider, error) {
	metricExporter, err := otlpmetricgrpc.New(ctx, otlpmetricgrpc.WithInsecure(), otlpmetricgrpc.WithEndpoint(ep))
	if err != nil {
		return nil, err
	}

	if intervalSecond < 10 {
		intervalSecond = 10
	}
	var interval = time.Duration(intervalSecond * int64(time.Second))
	meterProvider := metric.NewMeterProvider(
		metric.WithResource(res),
		metric.WithReader(metric.NewPeriodicReader(metricExporter, metric.WithInterval(interval))),
	)
	return meterProvider, nil
}
