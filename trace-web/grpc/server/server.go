package main

import (
	"context"
	"fmt"
	"net"
	"time"
	"trace/grpc/pb"

	"go.opentelemetry.io/contrib/instrumentation/google.golang.org/grpc/otelgrpc"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracegrpc"
	"go.opentelemetry.io/otel/propagation"
	"go.opentelemetry.io/otel/sdk/resource"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.17.0"
	"go.opentelemetry.io/otel/trace"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

const (
	serverName        = "gRPC-laeger-Demo"
	jaegerRPCEndpoint = "127.0.0.1:4317"
)

var tracer = otel.Tracer("grpc-example")

func newJaegerTraceProvider(c context.Context) (*sdktrace.TracerProvider, error) {
	exp, err := otlptracegrpc.New(c, otlptracegrpc.WithEndpoint(jaegerRPCEndpoint), otlptracegrpc.WithInsecure())
	if err != nil {
		return nil, err
	}
	res, err := resource.New(c, resource.WithAttributes(semconv.ServiceName(serverName)))
	if err != nil {
		return nil, err
	}
	traceProvider := sdktrace.NewTracerProvider(
		sdktrace.WithResource(res),
		sdktrace.WithSampler(sdktrace.AlwaysSample()),
		sdktrace.WithBatcher(exp, sdktrace.WithBatchTimeout(time.Second)),
	)
	return traceProvider, nil
}

func initTracer(ctx context.Context) (*sdktrace.TracerProvider, error) {
	tp, err := newJaegerTraceProvider(ctx)
	if err != nil {
		return nil, err
	}

	otel.SetTracerProvider(tp)
	otel.SetTextMapPropagator(
		propagation.NewCompositeTextMapPropagator(propagation.TraceContext{}, propagation.Baggage{}),
	)
	return tp, nil
}

type server struct {
	pb.UnimplementedGreeterServiceServer
}

func (s *server) SayHello(ctx context.Context, args *pb.SayHelloRequest) (*pb.SayHelloResponse, error) {
	md, _ := metadata.FromIncomingContext(ctx)
	_, span := tracer.Start(
		ctx,
		"SayHello",
		trace.WithAttributes(
			attribute.String("name", args.GetName()),
			attribute.StringSlice("client-id", md.Get("client-id")),
			attribute.StringSlice("user-id", md.Get("user-id")),
		),
	)
	defer span.End()
	return &pb.SayHelloResponse{
		Reply: "Hello" + args.Name,
	}, nil
}

func main() {
	ctx := context.Background()
	initTracer(ctx)
	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		fmt.Printf("listen tcp port failed, err=%v\n", err)
		return
	}
	s := grpc.NewServer(grpc.StatsHandler(otelgrpc.NewServerHandler()))
	pb.RegisterGreeterServiceServer(s, &server{})
	err = s.Serve(listener)
	if err != nil {
		fmt.Printf("Failed to serve: %v\n", err)
		return
	}
}
