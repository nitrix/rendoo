package main

import (
	"image"
	"image/color"
)

type Point struct {
	X, Y int
}

type Triangle struct {
	points [3]Point
}

func pointInTriangle(point Point, v1, v2, v3 Point) bool {
	w1, w2, _ := barycentric(point, v1, v2, v3)
	return w1 >= 0 && w1 <= 1 && w2 >= 0 && w2 <= 1 && w1 + w2 <= 1
}

func drawTriangle(img *image.RGBA, triangle Triangle, zBuffer []float64, texture image.Image, face Face) {
	width := img.Bounds().Dx()

	lightSource := Vertex{0, 0, 1}

	v1 := triangle.points[0]
	v2 := triangle.points[1]
	v3 := triangle.points[2]

	min, max := boundingBox(v1, v2, v3)
	for x := min.X; x <= max.X; x++ {
		for y := min.Y; y <= max.Y; y++ {
			if pointInTriangle(Point{x, y}, v1, v2, v3) {
				w1, w2, w3 := barycentric(Point{x,y}, v1, v2, v3)

				// Interpolate Z based on barycentric weights
				z := w1 * face.Vertices[0].Z + w2 * face.Vertices[1].Z + w3 * face.Vertices[2].Z

				// Interpolate normal based on barycentric weights
				normal := Vertex {
					X: w1 * face.Normals[0].X + w2 * face.Normals[1].X + w3 * face.Normals[2].X,
					Y: w1 * face.Normals[0].Y + w2 * face.Normals[1].Y + w3 * face.Normals[2].Y,
					Z: w1 * face.Normals[0].Z + w2 * face.Normals[1].Z + w3 * face.Normals[2].Z,
				}
				normal.normalize(1.0)

				// Texture coordinate
				txs := w1 * face.Textures[0].X + w2 * face.Textures[1].X + w3 * face.Textures[2].X
				tys := w1 * face.Textures[0].Y + w2 * face.Textures[1].Y + w3 * face.Textures[2].Y
				tx := int(txs * float64(texture.Bounds().Max.X))
				ty := int(float64(texture.Bounds().Max.Y) - tys * float64(texture.Bounds().Max.Y)) // Flip vertically!
				tcolor := texture.At(tx, ty)

				// Calculate light intensity
				intensity := normal.X * lightSource.X + normal.Y * lightSource.Y + normal.Z * lightSource.Z

				if intensity < 0 {
					continue
				}

				// Drawing according to Z-buffer
				if zBuffer[width*y+x] < z {
					zBuffer[width*y+x] = z
					r, g, b, _ := tcolor.RGBA()
					c := color.RGBA{
						R: uint8(float64(uint8(r)) * intensity),
						G: uint8(float64(uint8(g)) * intensity),
						B: uint8(float64(uint8(b)) * intensity),
						A: uint8(255),
					}
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