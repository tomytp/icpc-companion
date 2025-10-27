BIN ?= comp

.PHONY: build run snapshot release tidy

build:
	go build -o $(BIN) ./cmd/comp

run: build
	./$(BIN) --help

snapshot:
	goreleaser release --snapshot --skip-publish --clean

release:
	goreleaser release --clean

tidy:
	go mod tidy

