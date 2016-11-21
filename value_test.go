package unit

import (
	"math"
	"testing"
)

func TestMultiply(t *testing.T) {
	p := Define(1.0, G, Sep, Cm, Cm, Cm)
	v := Define(1.0, Mm, Mm, Mm)
	rtn := p.Multiply(v)
	expected := 0.001
	if math.Abs(rtn.Value-expected) > 1e-16 {
		t.Errorf("got %v, want %v", rtn.Value, expected)
	}
}

func TestAs(t *testing.T) {
	p := Define(1.0, G, Sep, Cm, Cm, Cm)
	rtn, err := p.As(T, Sep, M, M, M)
	if err != nil {
		t.Fatal(err)
	}
	expected := 1.0
	if math.Abs(rtn-expected) > 1e-16 {
		t.Errorf("got %v, want %v", rtn, expected)
	}
}
