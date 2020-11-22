# Welcome

!!! warning
    this package is a WIP and will introduce breaking changes while on major
    version zero.

[![Status](https://img.shields.io/badge/status-active-success.svg)](https://github.com/wwmoraes/dot)
![GitHub tag (latest SemVer)](https://img.shields.io/github/v/tag/wwmoraes/dot)
[![Go Report Card](https://goreportcard.com/badge/github.com/wwmoraes/dot)](https://goreportcard.com/report/github.com/wwmoraes/dot)
[![GoDoc](https://godoc.org/github.com/wwmoraes/dot?status.svg)](https://pkg.go.dev/github.com/wwmoraes/dot)

[![Quality Gate Status](https://sonarcloud.io/api/project_badges/measure?project=wwmoraes_dot&metric=alert_status)](https://sonarcloud.io/dashboard?id=wwmoraes_dot)
[![Reliability Rating](https://sonarcloud.io/api/project_badges/measure?project=wwmoraes_dot&metric=reliability_rating)](https://sonarcloud.io/dashboard?id=wwmoraes_dot)
[![Security Rating](https://sonarcloud.io/api/project_badges/measure?project=wwmoraes_dot&metric=security_rating)](https://sonarcloud.io/dashboard?id=wwmoraes_dot)
[![Maintainability Rating](https://sonarcloud.io/api/project_badges/measure?project=wwmoraes_dot&metric=sqale_rating)](https://sonarcloud.io/dashboard?id=wwmoraes_dot)

Dot is a lightweight, pure golang graphviz-compatible dot language
implementation focused on generating dot/gv files that [Graphviz](graphviz) can
then convert into any of its output formats supported (e.g. png, jpeg, svg, pdf).

Dot provides both interfaces and ready-to-use concrete types that represent
[dot language](dot-language) resources - namely Graphs, Nodes and Edges, plus all
attributes.

This package was inspired/initially forked from [emicklei/dot](emicklei-dot),
but has too many breaking changes compared to the original - namely interface
usage and other distinct design decisions - so it seemed better to maintain it
separately. If you need a simpler, no-brainy option, use
[emicklei's dot package](emicklei-dot).

## Installation

Fetch the package using go

```shell
go get -u github.com/wwmoraes/dot
```

## Usage

```go
package main

import (
  "os"
  "github.com/wwmoraes/dot"
)

func main() {
  graph := dot.New()

  graph.Node("n1").SetAttributeString("label", "hello dot!")

  fd, _ := os.Create("sample.dot")
  defer fd.Close()

  graph.WriteTo(fd)
}
```

[graphviz]: https://graphviz.org
[dot-language]: http://www.graphviz.org/doc/info/lang.html
[emicklei-dot]: https://github.com/emicklei/dot
