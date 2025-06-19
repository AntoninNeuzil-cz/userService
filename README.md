# User Service

User Service is a Go-based application that provides APIs to save and retrieve user data. It uses a database for persistent storage and follows the OpenAPI 3.0 specification.

## Features

- Save user data into the database.
- Retrieve user data by external ID.
- OpenAPI 3.0 specification for API documentation.

## Prerequisites

- Go 1.24 or later
- Docker
- Make

## Project Structure

- `cmd/server/server.go`: Entry point for the application.
- `internal/handler/handler.go`: Contains API handlers for user operations.
- `api/api.yaml`: OpenAPI specification for the service.
- `Makefile`: Build and run automation.

## Installation dependencies

- Command: `make vendor`

## Generating OpenAPI Documentation

- Command: `make generate`

## Build the application

- Command: `make build`

## Running service

- Command `make run` starts service locally
- Command `make run-docker` starts service in Docker container

## Running Tests
- Command: `make test` runs tests
- Curls for testing in postman:
  - Save: `curl --location 'http://localhost:8080/save' \
--header 'Content-Type: application/json' \
--data-raw '{
    "external_id": "123e4567-e89b-12d3-a456-426614174000",
    "name": "John Doe",
    "email": "john.doe@example.com",
    "date_of_birth": "1990-01-01T00:00:00Z"
}'`
  - Get `curl --location 'http://localhost:8080/123e4567-e89b-12d3-a456-426614174000'`

## Environment

- The application will be accessible on port 8080 by default.
- address: http://localhost:8080
