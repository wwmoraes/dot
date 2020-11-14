package constants

// PrimitiveShape edge arrow types
type PrimitiveShape string

const (
	// ShapeBox box arrow shape
	ShapeBox PrimitiveShape = "box"
	// ShapeCrow crow arrow shape
	ShapeCrow PrimitiveShape = "crow"
	// ShapeCurve curve arrow shape
	ShapeCurve PrimitiveShape = "curve"
	// ShapeICurve icurve arrow shape
	ShapeICurve PrimitiveShape = "icurve"
	// ShapeDiamond diamond arrow shape
	ShapeDiamond PrimitiveShape = "diamond"
	// ShapeDot dot arrow shape
	ShapeDot PrimitiveShape = "dot"
	// ShapeInv inv arrow shape
	ShapeInv PrimitiveShape = "inv"
	// ShapeNone none arrow shape
	ShapeNone PrimitiveShape = "none"
	// ShapeNormal normal arrow shape
	ShapeNormal PrimitiveShape = "normal"
	// ShapeTee tee arrow shape
	ShapeTee PrimitiveShape = "tee"
	// ShapeVee vee arrow shape
	ShapeVee PrimitiveShape = "vee"

	// ShapeEDiamond backwards-compatible ediamond shape
	ShapeEDiamond PrimitiveShape = "ediamond"
	// ShapeOpen backwards-compatible ediamond open
	ShapeOpen PrimitiveShape = "open"
	// ShapeHalfOpen backwards-compatible ediamond halfopen
	ShapeHalfOpen PrimitiveShape = "halfopen"
	// ShapeEmpty backwards-compatible ediamond empty
	ShapeEmpty PrimitiveShape = "empty"
	// ShapeInvEmpty backwards-compatible ediamond invempty
	ShapeInvEmpty PrimitiveShape = "invempty"
)
