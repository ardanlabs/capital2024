package testapi

import (
	"context"
	"encoding/json"
	"math/rand"
	"net/http"

	"github.com/ardanlabs/service/app/sdk/errs"
	"github.com/ardanlabs/service/foundation/web"
)

type status struct {
	Status string
}

func (s status) Encode() ([]byte, string, error) {
	data, err := json.Marshal(s)

	return data, "application/json", err
}

func test(ctx context.Context, r *http.Request) web.Encoder {
	if n := rand.Intn(100); n%2 == 0 {
		return errs.Newf(errs.InvalidArgument, "You are very bad: %s", "Bill")
	}

	status := status{
		Status: "OK",
	}

	return status
}
