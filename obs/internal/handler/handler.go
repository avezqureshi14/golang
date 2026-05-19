package handler

import (
	"context"
	"math/rand"
	"net/http"
	"time"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/trace"
)

func HelloHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	tr := otel.Tracer("hello-handler")

	// Root business span
	ctx, span := tr.Start(ctx, "hello-handler")
	defer span.End()

	userID := rand.Intn(100)

	span.SetAttributes(
		attribute.Int("user.id", userID),
	)

	// Step 1: validation
	validate(ctx, tr)

	// Step 2: business logic
	process(ctx, tr)

	// Step 3: external call simulation
	callExternal(ctx, tr)

	w.Write([]byte("hello world"))
}

func validate(ctx context.Context, tr trace.Tracer) {
	_, span := tr.Start(ctx, "validate-request")
	defer span.End()

	time.Sleep(20 * time.Millisecond)
}

func process(ctx context.Context, tr trace.Tracer) {
	_, span := tr.Start(ctx, "business-logic")
	defer span.End()

	time.Sleep(100 * time.Millisecond)
}

func callExternal(ctx context.Context, tr trace.Tracer) {
	_, span := tr.Start(ctx, "external-api-call")
	defer span.End()

	time.Sleep(80 * time.Millisecond)
}
