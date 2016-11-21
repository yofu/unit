package unit

// Weight is unit of weight.
type Weight struct {
	name  string
	value float64
}

var (
	// T is ton.
	T = Weight{
		name:  "t",
		value: 0.001,
	}
	// Kg is kilogram.
	Kg = Weight{
		name:  "kg",
		value: 1.0,
	}
	// G is gram.
	G = Weight{
		name:  "g",
		value: 1000.0,
	}
)

// String is for Unit interface.
func (w Weight) String() string {
	return w.name
}

// Value is for Unit interface
func (w Weight) Value() float64 {
	return w.value
}
