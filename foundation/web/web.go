package web

import (
	"context"
	"net/http"
)

// Encoder defines behavior that can encode a data model and provide
// the content type for that encoding.
type Encoder interface {
	Encode() (data []byte, contentType string, err error)
}

// Logger represents a function that will be called to add information
// to the logs.
type Logger func(ctx context.Context, msg string, args ...any)

// HandlerFunc represents a function that handles a http request within our own
// little mini framework.
type HandlerFunc func(ctx context.Context, r *http.Request) Encoder

// App is the entrypoint into our application and what configures our context
// object for each of our http handlers. Feel free to add any configuration
// data/logic on this App struct.
type App struct {
	*http.ServeMux
	log Logger
}

// NewApp creates an App value that handle a set of routes for the application.
func NewApp(log Logger) *App {
	return &App{
		ServeMux: http.NewServeMux(),
		log:      log,
	}
}

// HandleFunc sets a handler function for a given HTTP method and path pair
// to the application server mux.
func (a *App) HandleFunc(pattern string, handlerFunc HandlerFunc) {

	h := func(w http.ResponseWriter, r *http.Request) {

		// WE CAN DO WHAT WE WANT HERE

		ctx := r.Context()

		resp := handlerFunc(ctx, r)
		if err := Respond(ctx, w, resp); err != nil {
			a.log(ctx, "web-respond", "ERROR", err)
			return
		}

		// WE CAN DO WHAT WE WANT HERE
	}

	a.ServeMux.HandleFunc(pattern, h)
}
