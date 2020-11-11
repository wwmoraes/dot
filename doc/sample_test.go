package main

import (
	"errors"
	"io/ioutil"
	"os"
	"strings"
	"testing"
)

const dotFileName string = "sample.dot"

const expectedOutput string = `
digraph "" {
	subgraph "cluster_A" {
		label="Cluster A";
		"one";
		"two";
	}
	subgraph "cluster_B" {
		label="Cluster B";
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

func flatten(s string, t *testing.T) string {
	t.Helper()
	return strings.Replace((strings.Replace(s, "\n", "", -1)), "\t", "", -1)
}

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

	if got, want := flatten(string(fileData), t), flatten(expectedOutput, t); got != want {
		t.Errorf("got [\n%v\n] want [\n%v\n]", got, want)
	}
}
