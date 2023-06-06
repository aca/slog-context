# slog-context

Handler for [slog](https://pkg.go.dev/golang.org/x/exp/slog) for `context.Context` integration. Replacement for  [zerolog.Ctx](https://github.com/rs/zerolog#contextcontext-integration).

Instead of
```
slog.Info("called A", "user", "john", "traceID", "32423423423423")
```

Just 
```
slog.InfoCtx(ctx, "called A")
```

And use `context.Context` for contextual logging.
```
time=2023-06-06T19:11:24.245+09:00 level=INFO msg="called A" traceID=1772351682304946582 user="john"
```

## Examples

Basic
```
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
```

```
$ go run main.go
time=2023-06-06T22:12:35.540+09:00 level=INFO msg="called A" user=john
```

HTTP Middleware
```
package main

import (
	"context"
	"log"
	"math/rand"
	"net/http"
	"os"

	lctx "github.com/aca/slog-context"
	"golang.org/x/exp/slog"
)

func main() {
	ctxLogger := lctx.NewContextHandler(
		slog.NewTextHandler(os.Stdout, nil),
		[]lctx.ContextKey{
			TraceID,
		},
	)
	slog.SetDefault(slog.New(ctxLogger))

	http.HandleFunc("/ping", TraceMW(Ping))
	log.Fatal(http.ListenAndServe(":8080", nil))
}

var TraceID = lctx.Key("traceID")

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
```
```
$ curl localhost:8080/ping

$ go run main.go
time=2023-06-06T22:15:31.590+09:00 level=INFO msg=ping traceID=5040138652115587287
```
