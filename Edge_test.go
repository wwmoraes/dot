package dot

import (
	"fmt"
	"math"
	"testing"

	"github.com/wwmoraes/dot/dottest"
)

func TestEdge_String(t *testing.T) {
	// TODO String needs to be implemented, and will break this test when done so
	graph := NewGraph(nil)
	edge := graph.Node("n1").Edge(graph.Node("n2"))

	if got, want := dottest.MustGetFlattenSerializableString(t, edge), `"n1"->"n2";`; got != want {
		t.Errorf("got [\n%v\n] want [\n%v\n]", got, want)
	}
}

func TestEdge_ObjectInterface(t *testing.T) {
	graph := NewGraph(nil)
	graph.Node("n1").Edge(graph.Node("n2"))

	if got, want := dottest.MustGetFlattenSerializableString(t, graph), `digraph {"n1";"n2";"n1"->"n2";}`; got != want {
		t.Errorf("got [\n%v\n] want [\n%v\n]", got, want)
	}
}

func TestEdge_StyleHelpers(t *testing.T) {

	type test struct {
		input string
		want  string
	}

	tests := []test{
		{input: "solid", want: `digraph {"%[1]s";"%[2]s";"%[1]s"->"%[2]s"[style="solid"];}`},
		{input: "bold", want: `digraph {"%[1]s";"%[2]s";"%[1]s"->"%[2]s"[style="bold"];}`},
		{input: "dashed", want: `digraph {"%[1]s";"%[2]s";"%[1]s"->"%[2]s"[style="dashed"];}`},
		{input: "dotted", want: `digraph {"%[1]s";"%[2]s";"%[1]s"->"%[2]s"[style="dotted"];}`},
	}

	for _, tc := range tests {

		di := NewGraph(nil)
		n1 := di.Node("A")
		n2 := di.Node("B")

		switch tc.input {
		case "solid":
			di.Edge(n1, n2).Solid()
		case "bold":
			di.Edge(n1, n2).Bold()
		case "dashed":
			di.Edge(n1, n2).Dashed()
		case "dotted":
			di.Edge(n1, n2).Dotted()
		}

		if got, want := dottest.MustGetFlattenSerializableString(t, di), fmt.Sprintf(tc.want, n1.ID(), n2.ID()); got != want {
			t.Errorf("got [%v] want [%v]", got, want)
		}
	}
}

func TestEdge_WriteTo(t *testing.T) {
	tests := []struct {
		name       string
		limit      int
		wantErr    error
		wantString string
	}{
		{
			name:       "zero data written",
			limit:      0,
			wantErr:    dottest.ErrLimit,
			wantString: "",
		},
		{
			name:       "partially written - edge",
			limit:      1,
			wantErr:    dottest.ErrLimit,
			wantString: `"n1"->"n2"`,
		},
		{
			name:       "fully written",
			limit:      math.MaxInt32,
			wantErr:    nil,
			wantString: `"n1"->"n2"[label="test"];`,
		},
	}

	graph := NewGraph(nil)
	edge := graph.Edge(graph.Node("n1"), graph.Node("n2"))
	edge.SetAttributeString("label", "test")

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			wantN := int64(len(tt.wantString))
			dottest.TestByteWrite(t, edge, tt.limit, tt.wantErr, wantN, tt.wantString)
		})
	}
}
