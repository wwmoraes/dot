package main

import (
	"errors"
	"io/ioutil"
	"os"
	"testing"

	"github.com/wwmoraes/dot/dottest"
)

const dotFileName string = "sample.dot"

const expectedOutput string = `
digraph {
  subgraph "cluster_A" {
    graph [label="Cluster A"];
    "one";
    "two";
    {rank=same;"one";"two";}
  }
  subgraph "cluster_B" {
    graph [label="Cluster B"];
    "four";
    "three";
  }
  "Outside";
  "Outside"->"four";
  "four"->"one";
  "one"->"two";
  "three"->"Outside";
  "two"->"three";
}
`

func TestSample(t *testing.T) {
	if err := os.Remove(dotFileName); err != nil {
		if !errors.Is(err, os.ErrNotExist) {
			t.Fatal(err)
		}
	}

	main()

	fileInfo, err := os.Stat("sample.dot")
	if err != nil {
		t.Fatal(err)
	}

	if !fileInfo.Mode().IsRegular() {
		t.Fatalf("%s is not a regular file", dotFileName)
	}

	fileData, err := ioutil.ReadFile("sample.dot")
	if err != nil {
		t.Fatal(err)
	}

	if got, want := dottest.Flatten(t, string(fileData)), dottest.Flatten(t, expectedOutput); got != want {
		t.Errorf("got [\n%v\n] want [\n%v\n]", got, want)
	}
}
