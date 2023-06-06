package main

import (
	"context"
	"os"

	slogctx "github.com/aca/slog-context"
	"golang.org/x/exp/slog"
)

var UserID slogctx.ContextKey = "user"

func main() {
	ctxLogger := slogctx.NewContextHandler(slog.NewTextHandler(os.Stdout, nil), []slogctx.ContextKey{UserID})
	slog.SetDefault(slog.New(ctxLogger))

	ctx := context.Background()
	ctx = context.WithValue(ctx, UserID, "john")

	A(ctx)
}

func A(ctx context.Context) {
	slog.InfoCtx(ctx, "called A")
	slog.Info("called A", "user", "john", "traceID", "32423423423423")
}
