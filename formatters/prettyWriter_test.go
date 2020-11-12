package formatters

import (
	"bytes"
	"io"
	"io/ioutil"
	"os"
	"path"
	"reflect"
	"strings"
	"testing"
)

func TestPrettyWriter(t *testing.T) {
	want := strings.Join([]string{
		"digraph {",
		`  label="test";`,
		`  "n1";`,
		`  "n2";`,
		`  "n1"->"n2";`,
		"}",
		"",
	}, "\n")
	t.Run("string writer", func(t *testing.T) {
		var writer strings.Builder

		err := writeSample(t, &writer)
		if err != nil {
			t.Fatal(err)
		}

		got := writer.String()

		if !reflect.DeepEqual(got, want) {
			t.Fatalf("got [\n%v\n] want [\n%v\n]", got, want)
		}
	})
	t.Run("bytes writer", func(t *testing.T) {
		var buf bytes.Buffer

		err := writeSample(t, &buf)
		if err != nil {
			t.Fatal(err)
		}

		var stringWriter strings.Builder
		_, err = buf.WriteTo(&stringWriter)
		if err != nil {
			t.Fatal(err)
		}

		got := stringWriter.String()

		if !reflect.DeepEqual(got, want) {
			t.Fatalf("got [\n%v\n] want [\n%v\n]", got, want)
		}
	})
	t.Run("file writer", func(t *testing.T) {
		filePath := path.Join(t.TempDir(), "graph.dot")
		fd, err := os.Create(filePath)
		if err != nil {
			t.Fatal(err)
		}

		err = writeSample(t, fd)
		if err != nil {
			t.Fatal(err)
		}

		fd.Close()

		bytes, err := ioutil.ReadFile(filePath)
		if err != nil {
			t.Fatal(err)
		}

		got := string(bytes)

		if !reflect.DeepEqual(got, want) {
			t.Fatalf("got [\n%v\n] want [\n%v\n]", got, want)
		}
	})
}

func writeSample(t *testing.T, writer io.Writer) error {
	t.Helper()

	prettyWriter := NewPrettyWriter(writer)

	writeChunks := []string{
		"digraph {",
		`label="test";`,
		`"n1";`,
		`"n2";`,
		`"n1"->"n2";`,
		"}",
	}

	for _, chunk := range writeChunks {
		_, err := prettyWriter.Write([]byte(chunk))
		if err != nil {
			return err
		}
	}

	return nil
}
