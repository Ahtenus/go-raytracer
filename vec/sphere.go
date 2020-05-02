package vec

import "math"

// Sphere a sphere
type Sphere struct {
	Center Vec3
	Radius float64
}

// Hit where and how a ray hits a sphere
func (s *Sphere) Hit(ray Ray, tMin float64, tMax float64) (isHit bool, t float64, pos Vec3, normal Vec3, frontFace bool) {
	oc := ray.Orig.Sub(s.Center)
	a := ray.Dir.LengthSquared()
	halfB := oc.Dot(ray.Dir)
	c := oc.LengthSquared() - s.Radius*s.Radius
	discriminant := halfB*halfB - a*c
	if discriminant <= 0 {
		return
	}

	root := math.Sqrt(discriminant)
	for _, m := range []float64{-1, 1} {
		t = (-halfB - m*root) / a
		isHit = t < tMax && t > tMin
		if isHit {
			pos = ray.At(t)
			normal = pos.Sub(s.Center).Div(s.Radius)
			frontFace = ray.Dir.Dot(normal) < 0
			if !frontFace {
				// Point normal toward ray
				normal = normal.Mul(-1)
			}
			return
		}
	}
	return
}

// NewSphere with center and radius
func NewSphere(center Vec3, radius float64) Sphere {
	return Sphere{Center: center, Radius: 100}
}