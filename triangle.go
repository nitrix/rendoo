package main

import (
	"image"
	"image/color"
)

type Point struct {
	X, Y int
}

func sign(v1, v2, v3 Point) int {
	return (v1.X - v3.X) * (v2.Y - v3.Y) - (v2.X - v3.X) * (v1.Y - v3.Y)
}

func pointInTriangle(point Point, v1, v2, v3 Point) bool {
	b1 := sign(point, v1, v2) < 0.0
	b2 := sign(point, v2, v3) < 0.0
	b3 := sign(point, v3, v1) < 0.0

	return b1 == b2 && b2 == b3
}

/*
func pointInTriangle(point Point, v1, v2, v3 Point) bool {
	w1, w2, w3 := barycentric(point, v1, v2, v3)
	return w1 >= 0 && w2 >= 0 && w3 >= 0
}
*/

func drawFilledTriangle(img *image.RGBA, v1, v2, v3 Point, c color.Color) {
	min, max := boundingBox(v1, v2, v3)
	for x := min.X; x <= max.X; x++ {
		for y := min.Y; y <= max.Y; y++ {
			if pointInTriangle(Point{x, y}, v1, v2, v3) {
				img.Set(x, y, c)
			}
		}
	}
}

func drawFilledTriangleZBuffer(img *image.RGBA, v1, v2, v3 Point, zBuffer []float64, face Face, c color.Color) {
	width := img.Bounds().Dx()

	min, max := boundingBox(v1, v2, v3)
	for x := min.X; x <= max.X; x++ {
		for y := min.Y; y <= max.Y; y++ {
			if pointInTriangle(Point{x, y}, v1, v2, v3) {
				w1, w2, w3 := barycentric(Point{x,y}, v1, v2, v3)
				z := w1 * face.Vertices[0].Z + w2 * face.Vertices[1].Z + w3 * face.Vertices[2].Z

				if zBuffer[width*y+x] < z {
					zBuffer[width*y+x] = z
					img.Set(x, y, c)
				}
			}
		}
	}
}

func boundingBox(v1, v2, v3 Point) (Point, Point) {
	min := Point{
		X: minInt(minInt(v1.X, v2.X), v3.X),
		Y: minInt(minInt(v1.Y, v2.Y), v3.Y),
	}

	max := Point{
		X: maxInt(maxInt(v1.X, v2.X), v3.X),
		Y: maxInt(maxInt(v1.Y, v2.Y), v3.Y),
	}

	return min, max
}

func minInt(a, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}

func maxInt(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}