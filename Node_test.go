package dot

import (
	"math"
	"testing"

	"github.com/wwmoraes/dot/dottest"
)

func TestNode_String(t *testing.T) {
	graph, err := New()
	if err != nil {
		t.Fatal("graph is nil, expected a valid instance")
	}

	n1 := graph.Node("n1")

	if got, want := dottest.MustGetSerializableString(t, n1), `"n1";`; got != want {
		t.Errorf("got [%v] want [%v]", got, want)
	}
}

func TestNode_WriteTo(t *testing.T) {
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
			name:       "partially written - node ID",
			limit:      1,
			wantErr:    dottest.ErrLimit,
			wantString: `"n1"`,
		},
		{
			name:       "fully written",
			limit:      math.MaxInt32,
			wantErr:    nil,
			wantString: `"n1"[label="test"];`,
		},
	}

	graph, err := New()
	if err != nil {
		t.Fatal("graph is nil, expected a valid instance")
	}

	node := graph.Node("n1")
	node.SetAttributeString("label", "test")

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			wantN := int64(len(tt.wantString))
			dottest.TestByteWrite(t, node, tt.limit, tt.wantErr, wantN, tt.wantString)
		})
	}
}
