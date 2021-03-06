package main

import (
	"fmt"
	"github.com/Ahtenus/go-raytracer/vec"
	"log"
	"os"
	"os/exec"
	"math"
	"io"
)

func main() {
	filename := "out/hello_world"
	f, err := os.Create(filename + ".ppm")
	check(err)
	defer f.Close()

	imageWidth := 200
	imageHeight := 100

	s1 := vec.NewSphere(vec.Vec(0, 0, -1), 0.5)
	s2 := vec.NewSphere(vec.Vec(0, -100.5, -1), 100)
	world := []vec.Hittable{&s1, &s2}

	drawScene(imageWidth, imageHeight, world, f)
	
	log.Printf("Done writing %s", filename + ".ppm")
	convertToPng(filename)
	log.Printf("Done writing %s", filename + ".png")
}

func drawScene(imageWidth int, imageHeight int, world []vec.Hittable, f io.Writer) {
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
			rayColor(world, ray).WriteColor(f)
		}
	}
}

func hit(world []vec.Hittable, ray vec.Ray, tMin float64, tMax float64) (isHit bool, t float64, pos vec.Vec3, normal vec.Vec3, frontFace bool) {
	t = tMax
	for _, hittable := range world {
		isHitCur, tCur, posCur, normalCur, frontFaceCur := hittable.Hit(ray, tMin, t)
		if isHitCur {
			isHit = isHitCur
			t = tCur
			pos = posCur
			normal = normalCur
			frontFace = frontFaceCur
		}
	}
	return
}

func rayColor(world []vec.Hittable, r vec.Ray) vec.Vec3 {
	isHit, t, _, normal, _ := hit(world, r, 0, 1E6)
	if isHit {
		if(math.Abs(normal.Length() - 1.0) > 1E-6) {
			log.Printf("Normal of length %f", normal.Length())
		}
		return (vec.Vec(1,1,1).Add(normal)).Div(2)
	}
	t = 0.5*r.Dir.Unit().Y + 1.0
	return vec.Vec(1.0, 1.0, 1.0).Mul(1.0 - t).Add(vec.Vec(0.5, 0.7, 1.0).Mul(t))
}

func convertToPng(filename string) {
	cmd := exec.Command("convert", filename + ".ppm", filename + ".png")
	stdout, err := cmd.Output()
    if err != nil {
        log.Println(err.Error())
        return
    }
    log.Print(string(stdout))
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
