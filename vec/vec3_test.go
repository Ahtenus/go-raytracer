package vec

import (
	"testing"
)

func TestVec3(t *testing.T) {
	v := Vec3{10.0, 5.0, 3.0}
	v2 := Vec3{10.0, 5.0, 3.0}

	if v != v2 {
		t.Errorf("Neq")
	}

	a := Vec3{1.0, 2.0, 3.0}
	b := v.Add(a)
	c := b.Sub(a)

	if v != v2 {
		t.Errorf("Not equal: %+v %+v", v, v2)
	}

	if v != c {
		t.Errorf("Not equal: %+v %+v", v, v2)
	}

	v = v.Mul(3)
	v = v.Div(3)

	if v != v2 {
		t.Errorf("Not equal: %+v %+v", v, v2)
	}
}
