package tracelib

import (
	"testing"
)

func TestVec3(t *testing.T) {
	v := Vec3{[3]float64{10.0, 5.0, 3.0}}
	v2 := Vec3{[3]float64{10.0, 5.0, 3.0}}


	if v != v2 {
		t.Errorf("Neq")
	}
	
	a := Vec3{[3]float64{1.0, 2.0, 3.0}}
	v.add(a)
	v.sub(a)
	
	if v != v2 {
		t.Errorf("Not equal: %+v %+v", v, v2)
	}

	v.mul(3)
	v.div(3)

	if v != v2 {
		t.Errorf("Not equal: %+v %+v", v, v2)
	}
}