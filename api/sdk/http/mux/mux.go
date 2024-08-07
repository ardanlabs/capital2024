package mux

import (
	"context"

	"github.com/ardanlabs/service/api/domain/http/testapi"
	"github.com/ardanlabs/service/api/sdk/http/mid"
	"github.com/ardanlabs/service/foundation/logger"
	"github.com/ardanlabs/service/foundation/web"
)

// WebAPI constructs a http.Handler with all application routes bound.
func WebAPI(log *logger.Logger) *web.App {
	logger := func(ctx context.Context, msg string, args ...any) {
		log.Info(ctx, msg, args...)
	}

	app := web.NewApp(logger, mid.Logger(log), mid.Error(log), mid.Panic())

	testapi.Routes(app)

	return app
}
