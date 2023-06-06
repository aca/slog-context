package main

import (
	"context"

	lctx "github.com/aca/slog-context"
	"golang.org/x/exp/slog"
)

var UserID = lctx.Key("user")

func main() {
	lctx.SetDefaultTintDebugLogger(UserID)

	ctx := context.Background()
	ctx = context.WithValue(ctx, UserID, "john2")
	ctx = context.WithValue(ctx, UserID, "john2")

	A(ctx)
}

func A(ctx context.Context) {
	slog.InfoCtx(ctx, "called A")
	slog.DebugCtx(ctx, "called A")
	slog.ErrorCtx(ctx, "called A")
}

