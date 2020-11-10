package dot

import "testing"

func TestNode_String(t *testing.T) {
	// TODO String needs to be implemented, and will break this test when done so

	graph := NewGraph(nil)
	n1 := graph.Node("n1")

	if got, want := n1.String(), n1.ID(); got != want {
		t.Errorf("got [%v] want [%v]", got, want)
	}
}
