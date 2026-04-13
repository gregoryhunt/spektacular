# Testing Patterns

## Framework
- Go standard library `testing` for unit/integration tests under `internal/...`.
- Pytest 8.4.1 for harbor verifier tests (`tests/harbor/*/tests/test_*.py`), installed via `uv tool install` in the Dockerfile.
- Harbor (`harbor run`) orchestrates the agent run and verifier container.

## File Conventions
- Go: `*_test.go` colocated with source under `internal/`.
- Harbor tests: `tests/harbor/<task-name>/` containing:
  - `task.toml` — harbor task metadata, timeouts, artifacts.
  - `instruction.md` — prompt given to the agent.
  - `environment/Dockerfile` — container the agent runs in; copies a prebuilt `spektacular` binary to `/usr/local/bin/`.
  - `solution/solve.sh` — reference solution used to sanity-check the test harness.
  - `tests/test.sh` — pytest entrypoint; writes `1` or `0` to `/logs/verifier/reward.txt`.
  - `tests/test_<name>.py` — pytest module with per-behaviour test classes.

## Unit Tests (Go)
- Standard `go test ./...` via `make test`.

## Harbor Integration Tests
- Run via `make harbor-test`. The Makefile:
  - Cross-compiles a linux/amd64 binary into `tests/harbor/spec-workflow/environment/spektacular`.
  - Extracts an OAuth token from `~/.claude/.credentials.json`.
  - Runs `harbor run -p <task-dir> -a claude-code -m claude-sonnet-4-6 -o tests/harbor/jobs`.
  - Cats the latest `verifier/test-stdout.txt`.

### Verifier conventions (see `tests/harbor/spec-workflow/tests/test_spec_workflow.py`)
- Project files are at `/app`, the agent transcript is at `/logs/agent/claude-code.txt`.
- Tests are split into one class per workflow step plus a `TestWorkflow` class for cross-cutting invariants.
- Helper functions at the top of the file:
  - `load_state()` reads `.spektacular/state.json`.
  - `parse_sections()` splits the produced spec markdown into a dict keyed by lowercase heading, stripping HTML comments.
  - `extract_tool_calls()` parses the JSONL transcript for `{"type":"assistant"}` messages and returns every `Bash` `tool_use` block's command.
  - `find_spektacular_calls()` greps bash commands for `spektacular spec new` / `spektacular spec goto --data '{"step":"<name>"}'`.
- The canonical step order is hard-coded in the verifier (`EXPECTED_STEP_ORDER`) and must match `internal/spec/steps.go` `Steps()`.
- A `MIN_SECTION_LENGTH` (100 chars) is used as the substantive-content bar.
- Per-step tests check: step is in `completed_steps`, the right CLI call is in the transcript, the step's artefact/section exists with enough content, and (where relevant) mentions topic-specific terms.

### Plan-test-specific extensions we'll need (future work)
- Detect `Skill` tool_use blocks (or `spektacular skill <name>` Bash calls) in the transcript.
- Detect `Task` / `Agent` tool_use blocks for sub-agent spawning.
- Parse `templates/plan-steps/*.md` to discover which skills/sub-agents each step expects.
- Walk per-step artefacts under `.spektacular/plans/<name>/` (e.g. `research.md`, `plan.md`).

## Running Tests
- All Go tests: `make test`
- A single Go package: `go test ./internal/plan/...`
- Harbor spec-workflow test: `make harbor-test`
