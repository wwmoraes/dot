package constants

// DirType type for drawing arrowheads
type DirType string

const (
	// DirTypeForward draws a glyph at the head end of the edge
	DirTypeForward DirType = "forward"
	// DirTypeBack draws a glyph at the tail end of the edge
	DirTypeBack DirType = "back"
	// DirTypeBoth draws a glyph at the both ends of the edge
	DirTypeBoth DirType = "both"
	// DirTypeNone no glyph is drawn on the edge
	DirTypeNone DirType = "none"
)
