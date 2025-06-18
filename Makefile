
.PHONY: build
build:
	go build -ldflags "-X main.GitVersion=${GIT_VERSION}" -o user-service -v cmd/server/server.go

.PHONY: run
run: generate build
	./user-service

.PHONY: lint
lint:
	golangci-lint --timeout=5m run

.PHONY: vendor
vendor:
	go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest
	go install github.com/oapi-codegen/oapi-codegen/v2/cmd/oapi-codegen@latest
	go install github.com/atombender/go-jsonschema@v0.16.0

.PHONY: generate
generate:
	go generate ./...
	oapi-codegen --config oapi-codegen.config.yaml api/api.yaml > internal/gen/api/api.gen.go
	go-jsonschema -p hotel pkg/hotel/hotel.schema.json > pkg/hotel/hotel.gen.go

