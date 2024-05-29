package main

import (
	"flag"
	"fmt"
	"image"
	"image/color"
	"image/png"
	"log"
	"math"
	"math/rand/v2"
	"os"
)

func main() {
	size := flag.Int("size", 1024, "image size")
	flag.Parse()

	filename := flag.Arg(0)

	if filename == "" {
		log.Fatal(fmt.Errorf("please provide filename"))
	}

	bounds := image.Rect(0, 0, *size, *size)

	img := image.NewNRGBA(bounds)

	for x := range bounds.Dx() {
		for y := range bounds.Dy() {
			c := color.NRGBA{
				R: rand.N[uint8](math.MaxUint8),
				G: rand.N[uint8](math.MaxUint8),
				B: rand.N[uint8](math.MaxUint8),
				A: math.MaxUint8,
			}
			img.SetNRGBA(x, y, c)
		}
	}

	f, err := os.Create(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	err = png.Encode(f, img)
	if err != nil {
		log.Fatal(err)
	}
}
