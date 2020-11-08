package attributes

// PrimitiveShape edge arrow types
type PrimitiveShape *String

var (
	// ShapeBox box arrow shape
	ShapeBox PrimitiveShape = &String{"box"}
	// ShapeCrow crow arrow shape
	ShapeCrow PrimitiveShape = &String{"crow"}
	// ShapeCurve curve arrow shape
	ShapeCurve PrimitiveShape = &String{"curve"}
	// ShapeICurve icurve arrow shape
	ShapeICurve PrimitiveShape = &String{"icurve"}
	// ShapeDiamond diamond arrow shape
	ShapeDiamond PrimitiveShape = &String{"diamond"}
	// ShapeDot dot arrow shape
	ShapeDot PrimitiveShape = &String{"dot"}
	// ShapeInv inv arrow shape
	ShapeInv PrimitiveShape = &String{"inv"}
	// ShapeNone none arrow shape
	ShapeNone PrimitiveShape = &String{"none"}
	// ShapeNormal normal arrow shape
	ShapeNormal PrimitiveShape = &String{"normal"}
	// ShapeTee tee arrow shape
	ShapeTee PrimitiveShape = &String{"tee"}
	// ShapeVee vee arrow shape
	ShapeVee PrimitiveShape = &String{"vee"}

	// ShapeEDiamond backwards-compatible ediamond shape
	ShapeEDiamond PrimitiveShape = &String{"ediamond"}
	// ShapeOpen backwards-compatible ediamond open
	ShapeOpen PrimitiveShape = &String{"open"}
	// ShapeHalfOpen backwards-compatible ediamond halfopen
	ShapeHalfOpen PrimitiveShape = &String{"halfopen"}
	// ShapeEmpty backwards-compatible ediamond empty
	ShapeEmpty PrimitiveShape = &String{"empty"}
	// ShapeInvEmpty backwards-compatible ediamond invempty
	ShapeInvEmpty PrimitiveShape = &String{"invempty"}
)
