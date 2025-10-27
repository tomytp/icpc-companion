# puc-companion (comp)

Competitive Programming Companion CLI in Go. Listens for problems sent by CP Companion, organizes folders and tests, and provides quick run/test helpers for C++.

Highlights
- Listens on `localhost:10043` and creates problem folders/tests.
- Generic URL grouping + specific logic for Codeforces and VJudge.
- `comp test` compares `in/*` vs `out/*` with colorized output.
- `comp run` compiles and runs interactively (stdin/stdout).

Install
- Homebrew (macOS/Linux):
  - `brew tap tomytp/homebrew-tap`
  - `brew install puc-companion`
  - Binary name: `comp`
- Debian/Ubuntu (APT via Cloudsmith):
  - `curl -1sLf 'https://dl.cloudsmith.io/public/tomytp/icpc-companion/setup.deb.sh' | sudo -E bash`
  - `sudo apt-get update && sudo apt-get install puc-companion`

Quick Start
1) Configure paths:
   - `comp setup`  (set base_path, optional C++ template and makefile)
2) Listen for problems (opens VS Code after the first problem):
   - `comp solve`
3) Send a test payload (example – Codeforces):
   - `curl -sS -X POST http://localhost:10043 -H 'Content-Type: application/json' \
      -d '{"name":"CF A","url":"https://codeforces.com/contest/1234/problem/A","tests":[{"input":"1\n","output":"1\n"}]}'`
4) Work in the created folder, then:
   - `comp test` (or `comp test -d` for debug)
   - `comp run`  (or `comp run -d`, supports `< in/a1`)

Commands
- `comp setup`        Configure base path and templates
- `comp solve`        Listen and create problems/tests
- `comp listen`       Dry mode: pretty-print incoming JSON
- `comp test [-d]`    Build latest `*.cpp` and compare outputs
- `comp run  [-d]`    Build latest `*.cpp` and run interactively
- `comp completion`   Generate shell completion scripts

Development
- Build locally: `go build -o comp ./cmd/comp`
- Snapshot artifacts: `make snapshot` (no publish)
- Release via CI: `make release VERSION=X.Y.Z` (pushes `vX.Y.Z` tag)

Paths & Layout
- Config file: `~/.config/puc_rio_cp/config.json`
- Problem structure: `<base_path>/<platform>/<group>/`
- Tests: `in/<name>1`, `out/<name>1`, `in/<name>2`, ...

License
- MIT (see `LICENSE`).
