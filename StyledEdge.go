package dot

// StyledEdge is implemented by dot-compatible edge values which have
// convenience styling methods
type StyledEdge interface {
	Edge
	// Solid sets the edge style to solid
	Solid() Edge
	// Solid sets the edge style to bold
	Bold() Edge
	// Solid sets the edge style to dashed
	Dashed() Edge
	// Solid sets the edge style to dotted
	Dotted() Edge
}
