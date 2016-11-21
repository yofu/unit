package unit

// Length is unit of length.
type Length struct {
	name  string
	value float64
}

var (
	// M is meter.
	M = Length{
		name:  "m",
		value: 1.0,
	}
	// Cm is centimeter.
	Cm = Length{
		name:  "cm",
		value: 100.0,
	}
	// Mm is millimeter.
	Mm = Length{
		name:  "mm",
		value: 1000.0,
	}
)

// String is for Unit interface.
func (l Length) String() string {
	return l.name
}

// Value is for Unit interface
func (l Length) Value() float64 {
	return l.value
}
