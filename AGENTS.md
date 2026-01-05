# Repository Guidelines

## Project Structure & Module Organization
- `main.go` wires routes, services, and database connections; the API serves from `:3000`.
- `internal/` contains the app modules:
  - `internal/db/` database access and connection helpers.
  - `internal/services/` domain logic for doctors, patients, prescriptions, visits.
  - `internal/handlers/` HTTP handlers and request/response wiring.
  - `internal/models/` data models.
  - `internal/utils/` helpers like hashing and JWT.
- `API_DOCUMENTATION.md` describes available endpoints.
- `tmp/` is for local artifacts; do not commit new build outputs.

## Build, Test, and Development Commands
- `go run ./main.go` starts the API on port 3000 (requires `DATABASE_URL`).
- `go build ./...` builds all packages to catch compile errors.
- `go test ./...` runs tests (currently none; add as you implement new features).
- `gofmt -w .` formats Go files using standard Go conventions.
- `go vet ./...` runs basic static checks.

## Coding Style & Naming Conventions
- Use `gofmt` formatting (tabs for indentation, standard Go layout).
- Keep filenames in `camelCase` to match existing files (e.g., `doctorService.go`).
- Prefer short, focused functions and keep handler logic thin; place business logic in `internal/services/`.

## Testing Guidelines
- Use Go’s `testing` package for unit tests; place `_test.go` alongside the package under test.
- Name tests descriptively, e.g., `TestCreatePrescription`.
- Add coverage for handlers and services when adding new endpoints.

## Commit & Pull Request Guidelines
- Recent commits use short, imperative messages like “changed port to 3000”. Keep messages concise and action-focused.
- PRs should include a clear description, linked issues (if any), and how the change was tested (commands + results).
- Include API changes in `API_DOCUMENTATION.md` when endpoints or payloads change.

## Configuration & Security
- Set `DATABASE_URL` in your environment before running locally.
- Credentials or secrets should stay out of the repo; use local env files if needed.
