package main

import (
	// "bufio"
	"fmt"
	// "io/ioutil"
	"log"
	"os"
	"github.com/Ahtenus/go-raytracer/tracelib"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	filename := "out/hello_world.ppm"
	f, err := os.Create(filename)
	check(err)
	defer f.Close()

	imageWidth := 200
	imageHeight := 100

	fmt.Fprintf(f, "P3\n%d\n%d\n255\n", imageWidth, imageHeight)

	for j := imageHeight - 1; j >= 0; j-- {
		log.Printf("Scanlines remaining: %d", j)
		for i := 0; i < imageWidth; i++ {
			r := float64(i) / float64(imageWidth)
			g := float64(j) / float64(imageHeight)
			b := 0.2


			ir := int16(255.999 * r)
			ig := int16(255.999 * g)
			ib := int16(255.999 * b)

			fmt.Fprintf(f, "%d %d %d\n", ir, ig, ib)
			check(err)
		}
	}
	log.Printf("Done writing %s", filename)

	test := tracelib.Vec3{V: [3]float64{10.0, 5.0, 3.0}}
	log.Printf("%+v", test)
}
