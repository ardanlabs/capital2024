package mid

import (
	"context"
	"net/http"
	"time"

	"github.com/ardanlabs/service/foundation/logger"
	"github.com/ardanlabs/service/foundation/web"
)

func Logger(log *logger.Logger) web.MidFunc {
	m := func(handler web.HandlerFunc) web.HandlerFunc {

		h := func(ctx context.Context, r *http.Request) web.Encoder {
			now := time.Now()

			log.Info(ctx, "request started", "method", r.Method, "path", r.URL.Path, "remoteaddr", r.RemoteAddr)

			resp := handler(ctx, r)

			log.Info(ctx, "request completed", "method", r.Method, "path", r.URL.Path, "remoteaddr", r.RemoteAddr,
				"since", time.Since(now).String())

			return resp
		}

		return h
	}

	return m
}
