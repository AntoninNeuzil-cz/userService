
.PHONY: build
build:
	go build -o user-service cmd/server/server.go

.PHONY: run
run: generate build
	./user-service

.PHONY: lint
lint:
	golangci-lint --timeout=5m run

.PHONY: vendor
vendor:
	go mod tidy
	go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest
	go install github.com/oapi-codegen/oapi-codegen/v2/cmd/oapi-codegen@latest
	go install github.com/atombender/go-jsonschema@v0.16.0

.PHONY: generate
generate:
	go generate ./...
	oapi-codegen --config oapi-codegen.config.yaml api/api.yaml > internal/gen/api.gen.go

.PHONY: run-docker
run-docker:
	docker build -t user-service .
	docker run --rm -p 8080:8080 --name user-service-container user-service