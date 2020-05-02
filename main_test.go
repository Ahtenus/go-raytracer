package main

import (
	"testing"
	"github.com/Ahtenus/go-raytracer/vec"
	"bytes"
	"io/ioutil"
)


func TestHelloWorld(t *testing.T) {
	imageWidth := 200
	imageHeight := 100

	s1 := vec.NewSphere(vec.Vec(0, 0, -1), 0.5)
	s2 := vec.NewSphere(vec.Vec(0, -100.5, -1), 100)
	world := []vec.Hittable{&s1, &s2}

	buf := new(bytes.Buffer)
	drawScene(imageWidth, imageHeight, world, buf)

	expected, err := ioutil.ReadFile("test/hello_world.ppm")
    if err != nil {
		t.Fatal(err)
    }
	if(buf.String() != string(expected)) {
		t.Error("render not as expected")
	}
}