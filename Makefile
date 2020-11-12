SOURCES := $(wildcard *.go) $(wildcard */*.go)

.DEFAULT_GOAL := all

.PHONY: all
all: build test test-behavior coverage lint

.PHONY: lint
lint:
	golangci-lint run

.PHONY: cluster
cluster: sample/cluster.png
sample/cluster.png: sample/cluster.dot
sample/cluster.dot: sample/cluster.go
	cd sample && go run cluster.go

.PHONY: build
build: $(SOURCES)
	go build ./...

.PHONY: clean
clean:
	go clean -cache -testcache ./...
	rm -f coverage*.out coverage*.html

.PHONY: test
test: $(SOURCES)
	go test -race ./...

.PHONY: test-v
test-v: $(SOURCES)
	go test -race -v ./...

.PHONY: coverage
coverage: coverage.html

coverage.out: $(SOURCES)
	go test -cover -coverprofile=$@ ./...

.PHONY: test-behavior
test-behavior: coverage-behavior.html

coverage-behavior.out: $(SOURCES)
	go test -race -run ".*Behavior" -cover -coverprofile=$@ ./...

%.html: %.out
	go tool cover -html=$< -o $@

%.png: %.dot
	dot -Tpng -o $@ $<