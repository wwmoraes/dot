<h1 align="center">dot</h1>

<blockquote align="center">
a lightweight, pure golang graphviz-compatible dot language implementation
</blockquote>

<div align="center">

[![Status](https://img.shields.io/badge/status-active-success.svg)]()
[![GitHub Issues](https://img.shields.io/github/issues/wwmoraes/dot.svg)](https://github.com/wwmoraes/dot/issues)
[![GitHub Pull Requests](https://img.shields.io/github/issues-pr/wwmoraes/dot.svg)](https://github.com/wwmoraes/dot/pulls)
[![License](https://img.shields.io/badge/license-MIT-blue.svg)](/LICENSE)

[![Go Report Card](https://goreportcard.com/badge/github.com/wwmoraes/dot)](https://goreportcard.com/report/github.com/wwmoraes/dot)
[![GoDoc](https://godoc.org/github.com/wwmoraes/dot?status.svg)](https://pkg.go.dev/github.com/wwmoraes/dot)

[![Quality Gate Status](https://sonarcloud.io/api/project_badges/measure?project=wwmoraes_dot&metric=alert_status)](https://sonarcloud.io/dashboard?id=wwmoraes_dot)
[![Reliability Rating](https://sonarcloud.io/api/project_badges/measure?project=wwmoraes_dot&metric=reliability_rating)](https://sonarcloud.io/dashboard?id=wwmoraes_dot)
[![Bugs](https://sonarcloud.io/api/project_badges/measure?project=wwmoraes_dot&metric=bugs)](https://sonarcloud.io/dashboard?id=wwmoraes_dot)
[![Security Rating](https://sonarcloud.io/api/project_badges/measure?project=wwmoraes_dot&metric=security_rating)](https://sonarcloud.io/dashboard?id=wwmoraes_dot)
[![Vulnerabilities](https://sonarcloud.io/api/project_badges/measure?project=wwmoraes_dot&metric=vulnerabilities)](https://sonarcloud.io/dashboard?id=wwmoraes_dot)

[![Maintainability Rating](https://sonarcloud.io/api/project_badges/measure?project=wwmoraes_dot&metric=sqale_rating)](https://sonarcloud.io/dashboard?id=wwmoraes_dot)
[![Coverage](https://sonarcloud.io/api/project_badges/measure?project=wwmoraes_dot&metric=coverage)](https://sonarcloud.io/dashboard?id=wwmoraes_dot)
[![Code Smells](https://sonarcloud.io/api/project_badges/measure?project=wwmoraes_dot&metric=code_smells)](https://sonarcloud.io/dashboard?id=wwmoraes_dot)
[![Technical Debt](https://sonarcloud.io/api/project_badges/measure?project=wwmoraes_dot&metric=sqale_index)](https://sonarcloud.io/dashboard?id=wwmoraes_dot)

</div>

## 📝 Table of Contents

- [About](#about)
- [Getting Started](#getting_started)
- [Usage](#usage)
- [Contributing](../CONTRIBUTING.md)
- [Authors](#authors)
- [Acknowledgments](#acknowledgement)

## 🧐 About <a name = "about"></a>

> WARNING: this package is a WIP and will introduce breaking changes while on
> major version zero.

Dot provides interfaces and ready-to-use concrete types to create
[graphviz](graphviz)-compatible graphs using its [dot language](dotlanguage).

This package was inspired/initially forked from [emicklei/dot](emicklei-dot),
but has too many breaking changes compared to the original - namely interface
usage and other distinct design decisions - that I decided to maintain it
separately. If you need a simpler, no-brainy option, use
[emicklei's dot package](emicklei-dot).

## 🏁 Getting Started <a name = "getting_started"></a>

Clone the repository and then run `make` to build, test, generate coverage and
lint the code.

### Prerequisites

Golang 1.15 as of now, but should work with older golang versions.

No packages are needed.

## 🎈 Usage <a name = "usage"></a>

Add it to your modules with

```shell
go get -u github.com/wwmoraes/dot
```

And then:

```go
package main

import (
  "os"
  "github.com/wwmoraes/dot"
  "github.com/wwmoraes/dot/attributes"
)

func main() {
  graph := dot.NewGraph(nil)
  clusterA := graph.Subgraph(&attributes.GraphOptions{ ID: "Cluster A", Cluster: true })
  clusterA.SetAttributeString("label", "Cluster A")
  clusterB := graph.Subgraph(&attributes.GraphOptions{ ID: "Cluster B", Cluster: true })
  clusterB.SetAttributeString("label", "Cluster B")

  clusterA.
    Node("one").
    Edge(clusterA.Node("two")).
    Edge(clusterB.Node("three")).
    Edge(graph.Node("Outside")).
    Edge(clusterB.Node("four")).
    Edge(clusterA.Node("one"))

  graph.Write(os.Create("sample.dot"))
}
```

The attributes sub-package has all supported keys defined as variables, and can
be used instead of plain strings to avoid both duplication and errors:

```go
graph := dot.NewGraph(nil)
graph.Node("n1").SetAttributeString(attributes.KeyLabel, "my label")
```

You can also set literals and HTML values using the helper functions:

```go
graph := dot.NewGraph(nil)
graph.Node("n1").SetAttributeLiteral(attributes.KeyLabel, `my left label\l`)
graph.Node("n2").SetAttributeHTML(attributes.KeyLabel, `<b>a bold label</b>`)
```

## ✍️ Authors <a name = "authors"></a>

- [@emicklei](https://github.com/emicklei) - Original package
- [@wwmoraes](https://github.com/wwmoraes) - Modified version

## 🎉 Acknowledgements <a name = "acknowledgement"></a>

- [@emicklei](https://github.com/emicklei) for the amazing original package,
which initially helped me [get rid of a graphviz cgo package](goccy-go-graphviz)
I used on the [kubegraph project](https://github.com/wwmoraes/kubegraph)
- [@damianopetrungaro](https://github.com/damianopetrungaro) for the reviews and
discussion about golang ways - my personal master Yoda!

[graphviz]: https://graphviz.org
[dotlanguage]: http://www.graphviz.org/doc/info/lang.html
[emicklei-dot]: https://github.com/emicklei/dot
[goccy-go-graphviz]: https://github.com/goccy/go-graphviz
