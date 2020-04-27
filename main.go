package main

import (
	// "bufio"
	"fmt"
	// "io/ioutil"
	"github.com/Ahtenus/go-raytracer/vec"
	"log"
	"os"
)

func main() {
	filename := "out/hello_world.ppm"
	f, err := os.Create(filename)
	check(err)
	defer f.Close()

	imageWidth := 200
	imageHeight := 100

	lowerLeft := vec.Vec(-2.0, -1.0, -1.0)
	horizontal := vec.Vec(4.0, 0.0, 0.0)
	vertical := vec.Vec(0.0, 2.0, 0.0)
	origin := vec.Vec(0.0, 0.0, 0.0)

	fmt.Fprintf(f, "P3\n%d\n%d\n255\n", imageWidth, imageHeight)

	for j := imageHeight - 1; j >= 0; j-- {
		log.Printf("Scanlines remaining: %d", j)
		for i := 0; i < imageWidth; i++ {
			u := float64(i) / float64(imageWidth)
			v := float64(j) / float64(imageHeight)
			ray := vec.Ray{Orig: origin, Dir: lowerLeft.Add(horizontal.Mul(u)).Add(vertical.Mul(v))}
			colorRay(ray).WriteColor(f)
		}
	}
	log.Printf("Done writing %s", filename)
}

func colorRay(r vec.Ray) vec.Vec3 {
	if hitSphere(vec.Vec(0.0, 0.0, 1.0), 0.5, r) {
		return vec.Vec(1.0, 0.0, 0.0)
	}
	unitDirection := r.Dir.Unit()
	t := 0.5*unitDirection.Y + 1.0
	return vec.Vec(1.0, 1.0, 1.0).Mul(1.0 - t).Add(vec.Vec(0.5, 0.7, 1.0).Mul(t))
}

func hitSphere(center vec.Vec3, radius float64, ray vec.Ray) bool {
	oc := ray.Orig.Sub(center)
	a := ray.Dir.Dot(ray.Dir)
	b := 2.0 * oc.Dot(ray.Dir)
	c := oc.Dot(oc) - radius*radius
	discriminant := b*b - 4*a*c
	return discriminant > 0
}
func check(e error) {
	if e != nil {
		panic(e)
	}
}
