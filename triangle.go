package main

import (
	"image"
	"image/color"
)

type Triangle struct {
	points [3]image.Point
}

func drawTriangle(img *image.RGBA, triangle Triangle, zBuffer []float64, texture image.Image, face Face) {
	width := img.Bounds().Dx()
	height := img.Bounds().Dy()

	lightSource := Vertex3{0, 0, 1}

	v1 := triangle.points[0]
	v2 := triangle.points[1]
	v3 := triangle.points[2]

	min, max := boundingBox(v1, v2, v3)
	for x := min.X; x <= max.X; x++ {
		if x < 0 || x >= width {
			break
		}

		for y := min.Y; y <= max.Y; y++ {
			if y < 0 || y >= height {
				break
			}

			p := image.Point{X: x, Y: y}
			w1, w2, w3 := barycentric(p, v1, v2, v3)

			// If point in triangle
			if w1 >= 0 && w1 <= 1 && w2 >= 0 && w2 <= 1 && w1+w2 <= 1 {
				// Interpolate depth based on barycentric weights
				depth := w1*face.Vertices[0].Z + w2*face.Vertices[1].Z + w3*face.Vertices[2].Z

				// Interpolate normal based on barycentric weights
				normal := Vertex3{
					X: w1*face.Normals[0].X + w2*face.Normals[1].X + w3*face.Normals[2].X,
					Y: w1*face.Normals[0].Y + w2*face.Normals[1].Y + w3*face.Normals[2].Y,
					Z: w1*face.Normals[0].Z + w2*face.Normals[1].Z + w3*face.Normals[2].Z,
				}
				normal.normalize(1.0)

				// Interpolate texture based on barycentric weights
				txs := w1*face.Textures[0].X + w2*face.Textures[1].X + w3*face.Textures[2].X
				tys := w1*face.Textures[0].Y + w2*face.Textures[1].Y + w3*face.Textures[2].Y
				tx := int(txs * float64(texture.Bounds().Max.X))
				ty := int(tys * float64(texture.Bounds().Max.Y))
				tcolor := texture.At(tx, ty)

				// Calculate light intensity
				intensity := normal.X*lightSource.X + normal.Y*lightSource.Y + normal.Z*lightSource.Z

				if intensity < 0 {
					continue
				}

				// Drawing according to Z-buffer
				if zBuffer[width*y+x] < depth {
					zBuffer[width*y+x] = depth
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

func boundingBox(v1, v2, v3 image.Point) (image.Point, image.Point) {
	min := image.Point{
		X: minInt(minInt(v1.X, v2.X), v3.X),
		Y: minInt(minInt(v1.Y, v2.Y), v3.Y),
	}

	max := image.Point{
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
