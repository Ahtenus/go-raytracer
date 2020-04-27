package vec

// Ray represents a ray along the line p(t) = Orig + t * Dir 
type Ray struct {
	Orig Vec3
	Dir Vec3
}

// At Point at p(t) = Orig + t * Dir 
func (r Ray) At(t float64) Vec3 {
	return r.Orig.Add(r.Dir.Mul(t))
}