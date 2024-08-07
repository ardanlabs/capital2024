package testapi

import "github.com/ardanlabs/service/foundation/web"

// Routes adds specific routes for this group.
func Routes(app *web.App) {
	app.HandleFunc("GET /test", test)
}
