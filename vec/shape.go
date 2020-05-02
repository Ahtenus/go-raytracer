package vec

// Hittable shape that can be hit by a ray
type Hittable interface {
	Hit(ray Ray, tMin float64, tMax float64) (
		isHit bool,
		t float64,
		pos Vec3,
		normal Vec3,
		frontFace bool,
	)
}
