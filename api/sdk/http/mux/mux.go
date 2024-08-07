package mux

import (
	"github.com/ardanlabs/service/api/domain/http/testapi"
	"github.com/ardanlabs/service/foundation/web"
)

// WebAPI constructs a http.Handler with all application routes bound.
func WebAPI() *web.App {
	app := web.NewApp()

	testapi.Routes(app)

	return app
}
