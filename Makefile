SOURCES := $(wildcard *.go)

.PHONY: build
build: $(SOURCES)
	go build ./...

.PHONY: clean
clean:
	go clean -cache -testcache ./...
	rm -f coverage.out coverage.html

.PHONY: test
test: $(SOURCES)
	go test -race -v ./...

.PHONY: coverage
coverage: coverage.html

coverage.html: coverage.out
	go tool cover -html=$< -o $@

coverage.out: $(SOURCES)
	go test -cover -coverprofile=$@ -v ./...