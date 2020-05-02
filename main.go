package main

import (
	"fmt"
	"github.com/Ahtenus/go-raytracer/vec"
	"log"
	"os"
	"os/exec"
)

func main() {
	filename := "out/hello_world"
	f, err := os.Create(filename + ".ppm")
	check(err)
	defer f.Close()

	imageWidth := 200
	imageHeight := 100

	lowerLeft := vec.Vec(-2.0, -1.0, -1.0)
	horizontal := vec.Vec(4.0, 0.0, 0.0)
	vertical := vec.Vec(0.0, 2.0, 0.0)
	origin := vec.Vec(0.0, 0.0, 0.0)

	s1 := vec.NewSphere(vec.Vec(0, 0, -1), 0.5)
	// s2 := vec.NewSphere(vec.Vec(0, -100.5, -1), 100)
	world := []vec.Hittable{&s1}

	fmt.Fprintf(f, "P3\n%d\n%d\n255\n", imageWidth, imageHeight)

	for j := imageHeight - 1; j >= 0; j-- {
		log.Printf("Scanlines remaining: %d", j)
		for i := 0; i < imageWidth; i++ {
			u := float64(i) / float64(imageWidth)
			v := float64(j) / float64(imageHeight)
			ray := vec.Ray{Orig: origin, Dir: lowerLeft.Add(horizontal.Mul(u)).Add(vertical.Mul(v))}
			rayColor(world, ray).WriteColor(f)
		}
	}
	log.Printf("Done writing %s", filename + ".ppm")
	convertToPng(filename)
	log.Printf("Done writing %s", filename + ".png")
}

func hit(world []vec.Hittable, ray vec.Ray, tMin float64, tMax float64) (isHit bool, t float64, pos vec.Vec3, normal vec.Vec3, frontFace bool) {
	isHit = false
	for _, hittable := range world {
		isHit, t, pos, normal, frontFace = hittable.Hit(ray, tMin, tMax)
		if isHit {
			tMax = t
		}
	}
	return
}

func rayColor(world []vec.Hittable, r vec.Ray) vec.Vec3 {
	isHit, t, _, normal, _ := hit(world, r, 0, 1E6)
	if isHit { // 
		return (vec.Vec(1,1,1).Add(normal)).Mul(0.5)
	}
	unitDirection := r.Dir.Unit()
	t = 0.5*unitDirection.Y + 1.0
	return vec.Vec(1.0, 1.0, 1.0).Mul(1.0 - t).Add(vec.Vec(0.5, 0.7, 1.0).Mul(t))
}

func convertToPng(filename string) {
	cmd := exec.Command("convert", filename + ".ppm", filename + ".png")
	stdout, err := cmd.Output()
    if err != nil {
        log.Println(err.Error())
        return
    }

    fmt.Print(string(stdout))
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
