# Project Commands

## Build
- `make build` — Builds the `spektacular` binary into `./bin/spektacular` with version linker flags.
- `make cross` — Cross-compiles binaries for darwin/linux/windows (amd64 & arm64) into `./bin/`.
- `go build .` — Standard Go build of the root package.

## Test
- `make test` — Runs all Go unit/integration tests via `go test ./...`.
- `go test ./internal/...` — Scoped tests for internal packages.
- `make harbor-test` — Builds a linux/amd64 `spektacular` binary into `tests/harbor/spec-workflow/environment/`, runs the spec-workflow harbor task with `claude-sonnet-4-6`, and prints the verifier's `test-stdout.txt`. Requires `harbor` CLI and an Anthropic OAuth token extracted from `~/.claude/.credentials.json`.

## Lint
- `make lint` — `go vet ./...`.

## Dev / Install
- `make install-local` — Builds then copies the binary to `/usr/local/bin/spektacular` (requires sudo).
- `make clean` — Removes `./bin`.

## Other
- `go run . <subcommand>` — Run CLI directly without building. E.g. `go run . plan new --data '{"name":"..."}'`, `go run . skill <name>`, `go run . plan goto --data '{"step":"..."}'`.
