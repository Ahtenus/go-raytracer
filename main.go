package main

import (
	// "bufio"
	"fmt"
	// "io/ioutil"
	"log"
	"os"
	"github.com/Ahtenus/go-raytracer/vec"
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
			color := vec.New3(float64(i) / float64(imageWidth), float64(j) / float64(imageHeight), 0.2)
			color.WriteColor(f)
		}
	}
	log.Printf("Done writing %s", filename)
}
