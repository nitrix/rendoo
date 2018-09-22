package main

import (
	"fmt"
	"image"
	"image/png"
	"log"
	"os"
	"time"
)

func absInt(n int) int {
	if n < 0 {
		return n * -1
	} else {
		return n
	}
}

func main() {
	rect := image.Rectangle{Max: image.Point{X: 800, Y: 800}}
	img := newImage(rect)

	obj, err := loadObjFromFile("models/african_head.obj")
	if err != nil {
		log.Fatalln("Unable to load obj file:", err)
	}

	textureFile, err := os.Open("textures/african_head_diffuse.png")
	if err != nil {
		log.Fatalln("Unable to read texture from file:", err)
	}
	texture, err := png.Decode(textureFile)
	if err != nil {
		log.Fatalln("Unable to decode texture:", err)
	}

	texture = flipImageVertically(texture.Bounds(), texture)

	now := time.Now()
	fps := 0
	for time.Since(now) <= time.Second {
		render(img, obj, texture)
		fps++
	}
	fmt.Println("FPS:", fps)

	img = flipImageVertically(rect, img)
	saveImage(img)
}

func render(img *image.RGBA, obj *Obj, texture image.Image) {
	rect := img.Bounds()
	width := rect.Dx()
	height := rect.Dy()

	zBuffer := make([]float64, width*height)
	for i := 0; i < len(zBuffer); i++ {
		zBuffer[i] = -1.0
	}

	for _, face := range obj.Faces {
		triangle := Triangle{}

		for i := 0; i < 3; i++ {
			vertex := face.Vertices[i]

			x := (vertex.X + 1.0) * float64(width - 1) / 2
			y := (vertex.Y + 1.0) * float64(height - 1) / 2

			triangle.points[i].X = int(x)
			triangle.points[i].Y = int(y)
		}

		drawTriangle(
			img,
			triangle,
			zBuffer,
			texture,
			face,
		)
	}
}