package unit

import "fmt"

// Unit is an interface for unit.
type Unit interface {
	fmt.Stringer
	Value() float64
}
