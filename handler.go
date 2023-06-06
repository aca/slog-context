package context

import (
	"context"
	"os"
	"time"

	"github.com/lmittmann/tint"
	"golang.org/x/exp/slog"
)

type ContextKey string

type ContextHandler struct {
	handler slog.Handler
	keys    []ContextKey
}

func Key(v string) ContextKey {
	return ContextKey(v)
}

func NewContextHandler(h slog.Handler, k []ContextKey) *ContextHandler {
	return &ContextHandler{
		handler: h,
		keys:    k,
	}
}

func SetDefaultTintDebugLogger(keys... ContextKey) {
	tintHandler := tint.NewHandler(os.Stderr, &tint.Options{
        Level:      slog.LevelDebug,
        TimeFormat: time.Kitchen,
		AddSource: true,
    })

	h := NewContextHandler(tintHandler, keys)
	slog.SetDefault(slog.New(h))
}

func (h *ContextHandler) Handle(ctx context.Context, r slog.Record) error {
	for _, k := range h.keys {
		v := ctx.Value(k)
		if v != nil {
			r.AddAttrs(slog.Attr{
				Key:   string(k),
				Value: slog.AnyValue(v),
			})
		}
	}

	return h.handler.Handle(ctx, r)
}

func (h *ContextHandler) Enabled(ctx context.Context, level slog.Level) bool {
	return h.handler.Enabled(ctx, level)
}

func (h *ContextHandler) WithAttrs(attrs []slog.Attr) slog.Handler {
	return h.handler.WithAttrs(attrs)
}

func (h *ContextHandler) WithGroup(name string) slog.Handler {
	return h.handler.WithGroup(name)
}
