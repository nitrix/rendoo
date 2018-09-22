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

func newImage(rect image.Rectangle) *image.RGBA {
	img := image.NewRGBA(rect)
	for x := 0; x < rect.Max.X; x++ {
		for y := 0; y < rect.Max.Y; y++ {
			img.Set(x, y, color.RGBA{R: 0, G: 0, B: 0, A: 255})
		}
	}
	return img
}

func flipImageVertically(rect image.Rectangle, img image.Image) *image.RGBA {
	rgba := newImage(rect)

	for x := 0; x < rect.Max.X; x++ {
		for y := 0; y < rect.Max.Y; y++ {
			rgba.Set(x, rect.Max.Y - y, img.At(x, y))
		}
	}

	return rgba
}