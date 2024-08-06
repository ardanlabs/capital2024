# Check to see if we can use ash, in Alpine images, or default to BASH.
SHELL_PATH = /bin/ash
SHELL = $(if $(wildcard $(SHELL_PATH)),/bin/ash,/bin/bash)

run:
	go run api/cmd/services/sales/main.go | go run api/cmd/tooling/logfmt/main.go

run-help:
	go run api/cmd/services/sales/main.go --help | go run api/cmd/tooling/logfmt/main.go

run-build:
	go run -ldflags "-X main.build=$(shell git rev-parse --short HEAD)" api/cmd/services/sales/main.go | go run api/cmd/tooling/logfmt/main.go

# ==============================================================================
# Modules support

tidy:
	go mod tidy
	go mod vendor