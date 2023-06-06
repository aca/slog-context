package main

import (
	"context"
	"log"
	"math/rand"
	"net/http"
	"os"

	slogctx "github.com/aca/slog-context"
	"golang.org/x/exp/slog"
)

func main() {
	ctxLogger := slogctx.NewContextHandler(
		slog.NewTextHandler(os.Stdout, nil),
		[]slogctx.ContextKey{
			TraceID,
		},
	)
	slog.SetDefault(slog.New(ctxLogger))

	http.HandleFunc("/ping", TraceMW(Ping))
	log.Fatal(http.ListenAndServe(":8080", nil))
}

var TraceID slogctx.ContextKey = "traceID"

// Middleware to inject traceID for each request
func TraceMW(h http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		ctx = context.WithValue(ctx, TraceID, rand.Int())
		h(w, r.WithContext(ctx))
	}
}

func Ping(w http.ResponseWriter, r *http.Request) {
	slog.InfoCtx(r.Context(), "ping")
}
