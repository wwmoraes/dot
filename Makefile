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
coverage: coverage.out coverage.html
	go tool cover -func=$<

behavior: behavior.out behavior.html
	go tool cover -func=$<

coverage.out: $(SOURCES)
	go test -race -cover -coverprofile=$@ ./...

behavior.out: $(SOURCES)
	go test -race -run ".*Behavior" -coverprofile=$@ ./...

%.html: %.out
	go tool cover -html=$< -o $@

%.png: %.dot
	dot -Tpng -o $@ $<

sample/%.dot: sample/sample.go
	cd $(shell dirname $<) && go run $(shell basename $<)
