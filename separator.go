package unit

// Separator is an unit separator.
type Separator struct {
	name  string
	value float64
}

// Sep is an unit separator.
var Sep = Separator{
	name:  "separator",
	value: 0.0,
}

// String is for Unit interface.
func (s Separator) String() string {
	return s.name
}

// Value is for Unit interface.
func (s Separator) Value() float64 {
	return s.value
}
