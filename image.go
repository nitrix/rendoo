package main

import (
	"image"
	"image/color"
	"image/png"
	"log"
	"os"
)

func saveImage(img image.Image) {
	output, err := os.OpenFile("output.png", os.O_CREATE | os.O_WRONLY, 0777)
	if err != nil {
		log.Fatalln("Unable to open output file:", err)
	}

	err = png.Encode(output, img)
	if err != nil {
		log.Fatalln("Something went wrong writing to the output file:", err)
	}
}

func newImage(width, height int) *image.RGBA {
	img := image.NewRGBA(image.Rectangle{Max: image.Point{X: width, Y: height}})
	for x := 0; x < width; x++ {
		for y := 0; y < height; y++ {
			img.Set(x, y, color.RGBA{R: 0, G: 0, B: 0, A: 0})
		}
	}
	return img
}

func flipImageVertically(img *image.RGBA) *image.RGBA {
	bounds := img.Bounds()
	rgba := newImage(bounds.Max.X, bounds.Max.Y)
	for x := 0; x < bounds.Max.X; x++ {
		for y := 0; y < bounds.Max.Y; y++ {
			rgba.Set(x, bounds.Max.Y - y, img.At(x, y))
		}
	}
	return rgba
}