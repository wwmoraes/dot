package attributes

// Splines controls how, and if, edges are represented
type Splines *String

var (
	// SplinesEmpty edges are not drawn
	SplinesEmpty Splines = &String{""}
	// SplinesNone edges are not drawn
	SplinesNone Splines = &String{"none"}

	// SplinesFalse edges are drawn as line segment
	SplinesFalse Splines = &String{"false"}
	// SplinesLine edges are drawn as line segment
	SplinesLine Splines = &String{"line"}

	// SplinesTrue edges are drawn around nodes
	SplinesTrue Splines = &String{"true"}
	// SplinesSpline edges are drawn around nodes
	SplinesSpline Splines = &String{"spline"}

	// SplinesPolyline edges are drawn as polylines
	SplinesPolyline Splines = &String{"polyline"}
	// SplinesOrtho edges are drawn as polylines of axis-aligned segments
	SplinesOrtho Splines = &String{"ortho"}
	// SplinesCurved edges are drawn as curved arcs
	SplinesCurved Splines = &String{"curved"}
)
