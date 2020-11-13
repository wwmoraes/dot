SOURCES := $(wildcard *.go) $(wildcard */*.go)
PNGS := sample/plain.png sample/pretty.png
.DEFAULT_GOAL := all

.PHONY: all
all: build test coverage lint

.PHONY: lint
lint:
	golangci-lint run

.PHONY: pngs
pngs: $(PNGS)

.PHONY: build
build: $(SOURCES)
	go build ./...

.PHONY: clean
clean:
	find . -name "*.dot" -delete
	find . -name "*.png" -delete
	rm -f coverage*.out coverage*.html
	rm -rf site

.PHONY: clean-cache
clean-cache:
	go clean -cache -testcache ./...

.PHONY: test
test: $(SOURCES)
	go test -race ./...

.PHONY: test-v
test-v: $(SOURCES)
	go test -race -v ./...

.PHONY: coverage
coverage: coverage.html
	@touch doc.go
	@$(MAKE) $<

.PHONY: coverage-html
coverage-html: coverage.html

coverage.out: $(SOURCES)
	go test -race -cover -coverprofile=$@ ./...

.PHONY: coverage-behavior
coverage-behavior: coverage-behavior.html

coverage-behavior.out: $(SOURCES)
	go test -race -run ".*Behavior" -coverprofile=$@ ./...

%.html: %.out
	go tool cover -html=$< -o $@

%.png: %.dot
	dot -Tpng -o $@ $<

%.dot: %.go
	cd $(shell dirname $<) && go run $(shell basename $<)