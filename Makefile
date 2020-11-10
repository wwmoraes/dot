SOURCES := $(wildcard *.go) $(wildcard */*.go)

.PHONY: lint
lint:
	golangci-lint run

.PHONY: build
build: $(SOURCES)
	go build ./...

.PHONY: clean
clean:
	go clean -cache -testcache ./...
	rm -f coverage*.out coverage*.html

.PHONY: test
test: $(SOURCES)
	go test -race -v ./...

.PHONY: coverage
coverage: coverage.html

coverage.out: $(SOURCES)
	go test -cover -coverprofile=$@ -v ./...

.PHONY: test-behavior
test-behavior: coverage-behavior.html

coverage-behavior.out: $(SOURCES)
	go test -race -v -run ".*Behavior" -cover -coverprofile=$@ ./...

%.html: %.out
	go tool cover -html=$< -o $@