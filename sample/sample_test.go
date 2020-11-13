package main

import (
	"errors"
	"io/ioutil"
	"os"
	"testing"

	"github.com/wwmoraes/dot/dottest"
)

const expectedOutput string = `digraph {
  subgraph "cluster_A" {
    graph [label="Cluster A"];
    "one";
    "two";
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

type mainOutput struct {
	filename       string
	expectedOutput string
}

func TestSample_Cluster(t *testing.T) {
	outputs := []mainOutput{
		{
			filename:       "plain.dot",
			expectedOutput: dottest.Flatten(t, expectedOutput),
		},
		{
			filename:       "pretty.dot",
			expectedOutput: expectedOutput,
		},
	}
	testMain(t, main, outputs)
}

func testMain(tb testing.TB, main func(), outputs []mainOutput) {
	tb.Helper()

	for _, output := range outputs {
		if err := os.Remove(output.filename); err != nil {
			if !errors.Is(err, os.ErrNotExist) {
				tb.Fatal(err)
			}
		}
	}

	main()

	for _, output := range outputs {
		fileInfo, err := os.Stat(output.filename)
		if err != nil {
			tb.Fatal(err)
		}

		if !fileInfo.Mode().IsRegular() {
			tb.Fatalf("%s is not a regular file", output.filename)
		}

		fileData, err := ioutil.ReadFile(output.filename)
		if err != nil {
			tb.Fatal(err)
		}

		if got, want := string(fileData), output.expectedOutput; got != want {
			tb.Errorf("got [\n%v\n] want [\n%v\n]", got, want)
		}
	}
}
