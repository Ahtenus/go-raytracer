package vec

import (
	"fmt"
	"io"
	"math"
)
// Vec3 Represents a 3D vector
type Vec3 struct {
	V [3]float64
}

// New3 create a Vec3
func New3(v1 float64, v2 float64, v3 float64) Vec3 {
	return Vec3{[3]float64{v1, v2, v3}}
}

// Sub subtract and returns a new vector
func (a Vec3) Sub(b Vec3) Vec3 {
	a.V[0] -= b.V[0]
	a.V[1] -= b.V[1]
	a.V[2] -= b.V[2]
	return a
}

// Add adds and returns a new vector
func (a Vec3) Add(b Vec3) Vec3 {
	a.V[0] += b.V[0]
	a.V[1] += b.V[1]
	a.V[2] += b.V[2]
	return a
}

// Mul multiplies and returns a new vector
func (a Vec3) Mul(t float64) Vec3 {
	a.V[0] *= t
	a.V[1] *= t
	a.V[2] *= t
	return a
}

// Div divides and returns a new vector
func (a Vec3) Div(t float64) Vec3 {
	return a.Mul(1/t)
}

// X get x coordinate 
func (a Vec3) X() float64 {
	return a.V[0]
}

// Y get y coordinate 
func (a Vec3) Y() float64 {
	return a.V[1]
}

// Z get z coordinate 
func (a Vec3) Z() float64 {
	return a.V[2]
}

func (a Vec3) lengthSquared() float64 {
	sum := 0.0
	for _, v := range a.V {
		sum += v * v
	}
	return sum
}

func (a Vec3) length() float64 {
	return math.Sqrt(a.lengthSquared())
}

// Unit return a vector with the same direction but with length 1
func (a Vec3) Unit() Vec3 {
	return a.Div(a.length())
}

// Dot dot product of a and b
func (a Vec3) Dot(b Vec3) float64 {
	return a.V[0] * b.V[0] + a.V[1] * b.V[1] + a.V[2] * b.V[2]
}

// WriteColor of vector to 
func (a Vec3) WriteColor(out io.Writer) {
	r := uint8(255.999 * a.V[0])
	g := uint8(255.999 * a.V[1])
	b := uint8(255.999 * a.V[2])
	fmt.Fprintf(out,"%d %d %d\n", r, g, b)
}
