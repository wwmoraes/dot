<h1 align="center">dot</h1>

<blockquote align="center">
a lightweight, pure golang graphviz-compatible dot language implementation
</blockquote>

<div align="center">

[![Status](https://img.shields.io/badge/status-active-success.svg)][![FOSSA Status](https://app.fossa.com/api/projects/git%2Bgithub.com%2Fwwmoraes%2Fdot.svg?type=shield)](https://app.fossa.com/projects/git%2Bgithub.com%2Fwwmoraes%2Fdot?ref=badge_shield)
()
![GitHub tag (latest SemVer)](https://img.shields.io/github/v/tag/wwmoraes/dot)
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

Golang 1.12+ and ideally a modules-enabled application. Dot has no dependencies.

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
)

func main() {
  graph := dot.New()
  graph.SetAttributeString("label", "an amazing graph")
  clusterA := graph.Subgraph(WithID("Cluster A"), WithCluster())
  clusterA.SetAttributeString("label", "Cluster A")
  clusterB := graph.Subgraph(WithID("Cluster B"), WithCluster())
  clusterB.SetAttributeString("label", "Cluster B")

  clusterA.
    Node("one").
    Edge(clusterA.Node("two")).
    Edge(clusterB.Node("three")).
    Edge(graph.Node("Outside")).
    Edge(clusterB.Node("four")).
    Edge(clusterA.Node("one"))

  fd, _ := os.Create("sample.dot")
  graph.WriteTo(fd)
}
```

The constants sub-package has all supported keys defined as variables, and can
be used instead of plain strings to avoid both duplication and errors:

```go
package main

import (
  "os"
  "github.com/wwmoraes/dot"
  "github.com/wwmoraes/dot/constants"
)

func main() {
  graph := dot.New()
  graph.SetAttributeString(constants.KeyLabel, "a graph")
  node := graph.Node("n1")
  node.SetAttributeString(constants.KeyLabel, "a node")
  edge := graph.Edge(graph.Node("n2"), graph.Node("n3"))
  edge.SetAttributeString(constants.KeyLabel, "a edge")

  fd, _ := os.Create("sample.dot")
  graph.WriteTo(fd)
}
```

You can also set literals and HTML values using the helper functions:

```go
package main

import (
  "os"
  "github.com/wwmoraes/dot"
  "github.com/wwmoraes/dot/constants"
)

func main() {
  graph := dot.New()
  graph.Node("n1").SetAttributeLiteral(constants.KeyLabel, `a left label\l`)
  graph.Node("n2").SetAttributeHTML(constants.KeyLabel, `<b>a bold label</b>`)
  fd, _ := os.Create("sample.dot")
  graph.WriteTo(fd)
}
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


## License
[![FOSSA Status](https://app.fossa.com/api/projects/git%2Bgithub.com%2Fwwmoraes%2Fdot.svg?type=large)](https://app.fossa.com/projects/git%2Bgithub.com%2Fwwmoraes%2Fdot?ref=badge_large)