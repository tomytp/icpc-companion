# Repository Guidelines

## Project Structure & Module Organization
- Source: Go CLI in `cmd/comp/`; domain code in `internal/` (`config`, `server`, `platform`, `fs`, `runner`, `tester`, `util`).
- Platforms: `internal/platform` includes `generic`, plus specific `codeforces` and `vjudge` resolvers.
- Config: JSON at `~/.config/puc_rio_cp/config.json`.

## Build, Run, and Test
- Build: `go build -o comp ./cmd/comp` (binary `comp`).
- Help: `./comp --help` and `./comp <cmd> --help`.
- Configure: `./comp setup` (set `base_path`, optional template/makefile).
- Listen: `./comp solve` (HTTP on `localhost:10043`, creates folders/tests, opens VS Code after first problem).
- Dry listen: `./comp listen` (prints formatted JSON payloads, no writes).
- Test: `./comp test [-d]` (builds latest `*.cpp` via `make`, compares `in/*` to `out/*`).
- Run interactively: `./comp run [-d]` (builds latest `*.cpp`, runs with stdin/stdout; supports `< in/a1`).

## Coding Style & Naming
- Go 1.21+, idiomatic Go; keep packages focused and small.
- Names: packages `lowercase`, files `snake_case.go`, exported symbols with clear docs.
- Paths: folder structure `base_path/<platform>/<group>/`; solution `<file>.cpp`; tests in `in/` and `out/`.

## Testing Guidelines
- Primary testing is I/O based for problems. Name tests as `in/<name>1`, `out/<name>1`, `in/<name>2`, etc.
- CI builds releases via GoReleaser; add unit tests for URL resolvers as table tests when expanding platforms.

## Commit & PR Guidelines
- Commits: imperative, concise subject (≤72 chars), include rationale when needed.
- PRs: describe what/why, verification steps (commands/paths), and logs or screenshots for behavior changes.

## Release & Packaging
- GoReleaser: `.goreleaser.yaml` builds darwin/linux (amd64/arm64), Homebrew formula, `.deb` packages.
- Package name: `puc-companion` (Homebrew formula and .deb).
- Tag to release: `git tag vX.Y.Z && git push --tags` (CI publishes). Fill tap owner/repo and optional Cloudsmith settings.
