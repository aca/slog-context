package main

import (
	"context"
	"os"

	lctx "github.com/aca/slog-context"
	"golang.org/x/exp/slog"
)

var UserKey = lctx.Key("user")

func main() {
	ctxLogger := lctx.NewContextHandler(
		slog.NewTextHandler(os.Stdout, nil), 
		[]lctx.ContextKey{UserKey},
	)
	slog.SetDefault(slog.New(ctxLogger))

	ctx := context.Background()
	ctx = context.WithValue(ctx, UserKey, "john")

	A(ctx)
}

func A(ctx context.Context) {
	slog.InfoCtx(ctx, "called A")
}
