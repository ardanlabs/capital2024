package testapi

import (
	"context"
	"encoding/json"
	"net/http"

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
	status := status{
		Status: "OK",
	}

	return status
}
