package mid

import (
	"context"
	"net/http"
	"path"

	"github.com/ardanlabs/service/app/sdk/errs"
	"github.com/ardanlabs/service/foundation/logger"
	"github.com/ardanlabs/service/foundation/web"
)

// isError tests if the Encoder has an error inside of it.
func isError(e web.Encoder) error {
	err, isError := e.(error)
	if isError {
		return err
	}
	return nil
}

func Error(log *logger.Logger) web.MidFunc {
	m := func(handler web.HandlerFunc) web.HandlerFunc {
		h := func(ctx context.Context, r *http.Request) web.Encoder {
			resp := handler(ctx, r)

			err := isError(resp)
			if err == nil {
				return resp
			}

			appErr, ok := err.(*errs.Error)
			if !ok {
				appErr = errs.Newf(errs.Internal, "Internal Server Error")
			}

			log.Error(ctx, "handled error during request",
				"err", err,
				"source_err_file", path.Base(appErr.FileName),
				"source_err_func", path.Base(appErr.FuncName))

			// Send the error to the transport package so the error can be
			// used as the response.

			return appErr
		}

		return h
	}

	return m
}
