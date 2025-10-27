BIN ?= comp

.PHONY: build run snapshot release tidy tag

build:
	go build -o $(BIN) ./cmd/comp

run: build
	./$(BIN) --help

snapshot:
	goreleaser release --snapshot --skip-publish --clean

# Create and push a version tag to trigger CI release
release: tag

tag:
	@if [ -z "$(VERSION)" ]; then echo "VERSION required. Usage: make release VERSION=0.1.0"; exit 1; fi
	@git tag v$(VERSION)
	@git push origin v$(VERSION)

tidy:
	go mod tidy
