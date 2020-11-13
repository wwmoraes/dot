package dot

import (
	"math"
	"testing"

	"github.com/wwmoraes/dot/dottest"
)

func TestNode_String(t *testing.T) {
	// TODO String needs to be implemented, and will break this test when done so

	graph := NewGraph(nil)
	n1 := graph.Node("n1")

	if got, want := n1.String(), n1.ID(); got != want {
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

	graph := NewGraph(nil)
	node := graph.Node("n1")
	node.SetAttributeString("label", "test")

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			wantN := int64(len(tt.wantString))
			dottest.TestByteWrite(t, node, tt.limit, tt.wantErr, wantN, tt.wantString)
		})
	}
}
