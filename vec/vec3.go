package vec

import (
	"fmt"
	"io"
	"math"
)
// Vec3 Represents a 3D vector
type Vec3 struct {
	X, Y, Z float64
}

// Vec Create new Vec3
func Vec(x float64, y float64, z float64) Vec3 {
	return Vec3{X: x, Y: y, Z: z}
}
// Sub subtract and returns a new vector
func (a Vec3) Sub(b Vec3) Vec3 {
	a.X -= b.X
	a.Y -= b.Y
	a.Z -= b.Z
	return a
}

// Add adds and returns a new vector
func (a Vec3) Add(b Vec3) Vec3 {
	a.X += b.X
	a.Y += b.Y
	a.Z += b.Z
	return a
}

// Mul multiplies and returns a new vector
func (a Vec3) Mul(t float64) Vec3 {
	a.X *= t
	a.Y *= t
	a.Z *= t
	return a
}

// Div divides and returns a new vector
func (a Vec3) Div(t float64) Vec3 {
	return a.Mul(1/t)
}

func (a Vec3) lengthSquared() float64 {
	return a.X*a.X + a.Y*a.Y + a.Z * a.Z
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
	return a.X * b.X + a.Y * b.Y + a.Z * b.Z
}

// WriteColor of vector to 
func (a Vec3) WriteColor(out io.Writer) {
	r := uint8(255.999 * a.X)
	g := uint8(255.999 * a.Y)
	b := uint8(255.999 * a.Z)
	fmt.Fprintf(out,"%d %d %d\n", r, g, b)
}
