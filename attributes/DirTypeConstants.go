package attributes

// DirType type for drawing arrowheads
type DirType *String

var (
	// DirTypeForward draws a glyph at the head end of the edge
	DirTypeForward DirType = &String{"forward"}

	// DirTypeBack draws a glyph at the tail end of the edge
	DirTypeBack DirType = &String{"back"}

	// DirTypeBoth draws a glyph at the both ends of the edge
	DirTypeBoth DirType = &String{"both"}

	// DirTypeNone no glyph is drawn on the edge
	DirTypeNone DirType = &String{"none"}
)
