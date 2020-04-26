package vec

import (
	"fmt"
	"io"
)
// Vec3 Represents a 3D vector
type Vec3 struct {
	V [3]float64
}

// New3 create a Vec3
func New3(v1 float64, v2 float64, v3 float64) Vec3 {
	return Vec3{[3]float64{v1, v2, v3}}
}

func (a * Vec3) sub(b Vec3) {
	a.V[0] -= b.V[0]
	a.V[1] -= b.V[1]
	a.V[2] -= b.V[2]
}

func (a * Vec3) add(b Vec3) {
	a.V[0] += b.V[0]
	a.V[1] += b.V[1]
	a.V[2] += b.V[2]
}

func (a * Vec3) mul(t float64) {
	a.V[0] *= t
	a.V[1] *= t
	a.V[2] *= t
}

func (a * Vec3) div(t float64) {
	a.mul(1/t)
}

func (a * Vec3) lengthSquared() float64 {
	sum := 0.0
	for _, v := range a.V {
		sum += v * v
	}
	return sum
}

func (a * Vec3) WriteColor(out io.Writer) {
	r := uint8(255.999 * a.V[0])
	g := uint8(255.999 * a.V[1])
	b := uint8(255.999 * a.V[2])
	fmt.Fprintf(out,"%d %d %d\n", r, g, b)
}