package main

import (
	"image"
	"image/color"
	"log"
)

func absInt(n int) int {
	if n < 0 {
		return n * -1
	} else {
		return n
	}
}

func main() {
	width, height := 800, 800
	img := newImage(width, height, colors["black"])

	obj, err := loadObjFromFile("models/african_head.obj")
	if err != nil {
		log.Fatalln("Unable to load obj file:", err)
	}

	// drawWireframe(img, obj)
	// drawSolid(img, obj)
	// drawSolidShading(img, obj)
	drawSolidShadingWithZBuffer(img, obj)

	img = flipImageVertically(img)
	saveImage(img)
}

func drawSolidShadingWithZBuffer(img *image.RGBA, obj *Obj) {
	rect := img.Bounds()
	width := rect.Dx()
	height := rect.Dy()

	lightSource := Vertex{0, 0, -1}
	zBuffer := make([]float64, width*height)
	for i := 0; i < len(zBuffer); i++ {
		zBuffer[i] = -1.0
	}

	for _, face := range obj.Faces {
		screenCoordinates := [3]Point{}

		for i := 0; i < 3; i++ {
			vertex := face.Vertices[i]

			x := (vertex.X + 1.0) * float64(width - 1) / 2
			y := (vertex.Y + 1.0) * float64(height - 1) / 2

			screenCoordinates[i].X = int(x)
			screenCoordinates[i].Y = int(y)
		}

		faceNormal := face.Normal()
		faceNormal.normalize(1.0)
		intensity := faceNormal.X * lightSource.X + faceNormal.Y * lightSource.Y + faceNormal.Z * lightSource.Z

		// Back face culling if the intensity of the light is negative (light is coming from behind the face)
		if intensity > 0 {
			// drawFilledTriangle(img, screenCoordinates[0], screenCoordinates[1], screenCoordinates[2], color.RGBA{uint8(intensity * 255), uint8(intensity * 255), uint8(intensity * 255), 255})
			drawFilledTriangleZBuffer(img, screenCoordinates[0], screenCoordinates[1], screenCoordinates[2], zBuffer, face, color.RGBA{uint8(intensity * 255), uint8(intensity * 255), uint8(intensity * 255), 255})
		}
	}
}

func drawSolidShading(img *image.RGBA, obj *Obj) {
	rect := img.Bounds()
	width := rect.Dx()
	height := rect.Dy()

	lightSource := Vertex{0, 0, -1}

	for _, face := range obj.Faces {
		screenCoordinates := [3]Point{}

		for i := 0; i < 3; i++ {
			vertex := face.Vertices[i]

			x := (vertex.X + 1.0) * float64(width) / 2
			y := (vertex.Y + 1.0) * float64(height) / 2

			screenCoordinates[i].X = int(x)
			screenCoordinates[i].Y = int(y)
		}

		faceNormal := face.Normal()
		faceNormal.normalize(1.0)
		intensity := faceNormal.X * lightSource.X + faceNormal.Y * lightSource.Y + faceNormal.Z * lightSource.Z

		// Back face culling if the intensity of the light is negative (light is coming from behind the face)
		if intensity > 0 {
			drawFilledTriangle(img, screenCoordinates[0], screenCoordinates[1], screenCoordinates[2], color.RGBA{uint8(intensity * 255), uint8(intensity * 255), uint8(intensity * 255), 255})
		}
	}
}

func drawSolid(img *image.RGBA, obj *Obj) {
	rect := img.Bounds()
	width := rect.Dx()
	height := rect.Dy()

	for _, face := range obj.Faces {
		screenCoordinates := [3]Point{}

		for i := 0; i < 3; i++ {
			vertex := face.Vertices[i]

			x := (vertex.X + 1.0) * float64(width) / 2
			y := (vertex.Y + 1.0) * float64(height) / 2

			screenCoordinates[i].X = int(x)
			screenCoordinates[i].Y = int(y)
		}

		drawFilledTriangle(img, screenCoordinates[0], screenCoordinates[1], screenCoordinates[2], colors["white"])
	}
}

func drawWireframe(img *image.RGBA, obj *Obj) {
	rect := img.Bounds()
	width := rect.Dx()
	height := rect.Dy()

	for _, face := range obj.Faces {
		for i := 0; i < 3; i++ {
			fromVertex := face.Vertices[i]
			toVertex := face.Vertices[(i+1) % 3]

			fromX := (fromVertex.X + 1.0) * float64(width) / 2
			fromY := (fromVertex.Y + 1.0) * float64(height) / 2
			toX := (toVertex.X + 1.0) * float64(width) / 2
			toY := (toVertex.Y + 1.0) * float64(height) / 2

			drawLine(img, int(fromX), int(fromY), int(toX), int(toY), colors["white"])
		}
	}
}