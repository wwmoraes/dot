package formatters

import (
	"errors"
	"io/ioutil"
	"math"
	"os"
	"path"
	"testing"

	"github.com/wwmoraes/dot"
	"github.com/wwmoraes/dot/dottest"
)

func TestPrettyWriter(t *testing.T) {
	t.Run("invalid writer", func(t *testing.T) {
		graph := dot.NewGraph(nil)
		graph.Node("n1").Edge(graph.Node("n2"))

		prettyWriter := NewPrettyWriter(nil)

		gotN, gotErr := graph.WriteTo(prettyWriter)

		wantN := int64(0)
		wantErr := ErrNoWriter

		if gotN != wantN {
			t.Errorf("got [\n%v\n] want [\n%v\n]", gotN, wantN)
		}

		if !errors.Is(gotErr, wantErr) {
			t.Errorf("got [\n%v\n] want [\n%v\n]", gotErr, wantErr)
		}
	})
	t.Run("bytes writer", func(t *testing.T) {
		graph := dot.NewGraph(nil)
		graph.Node("n1").Edge(graph.Node("n2"))

		wantString := "digraph {\n  \"n1\";\n  \"n2\";\n  \"n1\"->\"n2\";\n}\n"
		wantN := int64(len(wantString))
		var wantErr error = nil

		stringWriter := dottest.NewByteWriter(t, math.MaxInt32, wantErr)
		prettyWriter := NewPrettyWriter(stringWriter)

		gotN, gotErr := graph.WriteTo(prettyWriter)
		gotString := stringWriter.String()

		if gotN != wantN {
			t.Errorf("got [\n%v\n] want [\n%v\n]", gotN, wantN)
		}

		if !errors.Is(gotErr, wantErr) {
			t.Errorf("got [\n%v\n] want [\n%v\n]", gotErr, wantErr)
		}

		if gotString != wantString {
			t.Errorf("got [\n%v\n] want [\n%v\n]", gotString, wantString)
		}
	})
	t.Run("file writer", func(t *testing.T) {
		graph := dot.NewGraph(nil)
		graph.Node("n1").Edge(graph.Node("n2"))

		filePath := path.Join(t.TempDir(), "graph.dot")
		fd, err := os.Create(filePath)
		if err != nil {
			t.Fatal(err)
		}
		defer fd.Close()

		prettyWriter := NewPrettyWriter(fd)
		gotN, gotErr := graph.WriteTo(prettyWriter)

		bytes, err := ioutil.ReadFile(filePath)
		if err != nil {
			t.Fatal(err)
		}

		gotString := string(bytes)

		wantString := "digraph {\n  \"n1\";\n  \"n2\";\n  \"n1\"->\"n2\";\n}\n"
		wantN := int64(len(wantString))
		var wantErr error = nil

		if gotN != wantN {
			t.Errorf("got [\n%v\n] want [\n%v\n]", gotN, wantN)
		}

		if !errors.Is(gotErr, wantErr) {
			t.Errorf("got [\n%v\n] want [\n%v\n]", gotErr, wantErr)
		}

		if gotString != wantString {
			t.Errorf("got [\n%v\n] want [\n%v\n]", gotString, wantString)
		}
	})
}
