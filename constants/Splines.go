package constants

// Splines controls how, and if, edges are represented
type Splines string

const (
	// SplinesEmpty edges are not drawn
	SplinesEmpty Splines = ""
	// SplinesNone edges are not drawn
	SplinesNone Splines = "none"

	// SplinesFalse edges are drawn as line segment
	SplinesFalse Splines = "false"
	// SplinesLine edges are drawn as line segment
	SplinesLine Splines = "line"

	// SplinesTrue edges are drawn around nodes
	SplinesTrue Splines = "true"
	// SplinesSpline edges are drawn around nodes
	SplinesSpline Splines = "spline"

	// SplinesPolyline edges are drawn as polylines
	SplinesPolyline Splines = "polyline"
	// SplinesOrtho edges are drawn as polylines of axis-aligned segments
	SplinesOrtho Splines = "ortho"
	// SplinesCurved edges are drawn as curved arcs
	SplinesCurved Splines = "curved"
)
