package unit

import (
	"bytes"
	"fmt"
	"reflect"
	"sort"
)

// Value represents a value with unit.
type Value struct {
	Value       float64
	Numerator   []Unit
	Denominator []Unit
}

// NewValue returns new Value.
func NewValue(val float64, num []Unit, den []Unit) *Value {
	return &Value{
		Value:       val,
		Numerator:   num,
		Denominator: den,
	}
}

func unitlist(units ...Unit) ([]Unit, []Unit) {
	num := make([]Unit, 0)
	den := make([]Unit, 0)
	sep := false
	for _, u := range units {
		if u == Sep {
			sep = true
			continue
		}
		if sep {
			den = append(den, u)
		} else {
			num = append(num, u)
		}
	}
	return num, den
}

// Define defines new Value.
func Define(val float64, units ...Unit) *Value {
	num, den := unitlist(units...)
	return &Value{
		Value:       val,
		Numerator:   num,
		Denominator: den,
	}
}

// Power returns unit to the power of num.
func Power(u Unit, num int) []Unit {
	rtn := make([]Unit, num)
	for i := 0; i < num; i++ {
		rtn[i] = u
	}
	return rtn
}

// HasSameUnit returns if u1 & u2 has the same unit.
func HasSameUnit(u1, u2 *Value) bool {
	if len(u1.Numerator) != len(u2.Numerator) || len(u1.Denominator) != len(u2.Denominator) {
		return false
	}
	for i := 0; i < len(u1.Numerator); i++ {
		if reflect.TypeOf(u1.Numerator[i]) != reflect.TypeOf(u2.Numerator[i]) {
			return false
		}
	}
	for i := 0; i < len(u1.Denominator); i++ {
		if reflect.TypeOf(u1.Denominator[i]) != reflect.TypeOf(u2.Denominator[i]) {
			return false
		}
	}
	return true
}

func coefficient(coeffs ...float64) float64 {
	sort.Sort(sort.Reverse(sort.Float64Slice(coeffs)))
	coeff := 1.0
	for _, c := range coeffs {
		coeff *= c
	}
	return coeff
}

// ConvertUnit converts val from unit1:num1/den1 to unit2:num2/den2.
func ConvertUnit(val float64, num1, den1, num2, den2 []Unit) (float64, error) {
	if len(num1) != len(num2) || len(den1) != len(den2) {
		return 0.0, fmt.Errorf("unit error")
	}
	coeffs := make([]float64, 0)
	for i, n := range num2 {
		if reflect.TypeOf(n) != reflect.TypeOf(num1[i]) {
			return 0.0, fmt.Errorf("unit error")
		}
		coeffs = append(coeffs, n.Value()/num1[i].Value())
	}
	for i, d := range den2 {
		if reflect.TypeOf(d) != reflect.TypeOf(den1[i]) {
			return 0.0, fmt.Errorf("unit error")
		}
		coeffs = append(coeffs, den1[i].Value()/d.Value())
	}
	return coefficient(coeffs...) * val, nil
}

// String is for Stringer interface.
func (v *Value) String() string {
	var otp bytes.Buffer
	otp.WriteString(fmt.Sprintf("%14.10f [", v.Value))
	for i, n := range v.Numerator {
		if i == 0 {
			otp.WriteString(n.String())
		} else {
			otp.WriteString(" " + n.String())
		}
	}
	if len(v.Denominator) > 0 {
		otp.WriteString("/")
		for i, d := range v.Denominator {
			if i == 0 {
				otp.WriteString(d.String())
			} else {
				otp.WriteString(" " + d.String())
			}
		}
	}
	otp.WriteString("]")
	return otp.String()
}

// Reduce reduces unit.
func (v *Value) Reduce() *Value {
	num := make([]Unit, 0)
	den := make([]Unit, 0)
	check := make([]bool, len(v.Denominator))
	coeffs := make([]float64, 0)
nloop:
	for _, n := range v.Numerator {
		for i, d := range v.Denominator {
			if check[i] {
				continue
			}
			if reflect.TypeOf(n) == reflect.TypeOf(d) {
				coeffs = append(coeffs, d.Value()/n.Value())
				check[i] = true
				continue nloop
			}
		}
		num = append(num, n)
	}
	for i, d := range v.Denominator {
		if !check[i] {
			den = append(den, d)
		}
	}
	return NewValue(v.Value*coefficient(coeffs...), num, den)
}

// As returns a value converted to given unit.
func (v *Value) As(units ...Unit) (float64, error) {
	num, den := unitlist(units...)
	return ConvertUnit(v.Value, v.Numerator, v.Denominator, num, den)
}

// Multiply multiplies v by u, and reduces unit.
func (v *Value) Multiply(u *Value) *Value {
	num := append(v.Numerator, u.Numerator...)
	den := append(v.Denominator, u.Denominator...)
	return NewValue(v.Value*u.Value, num, den).Reduce()
}
