package main

import "image"

type Face struct {
	Vertices [3]Vertex
	Textures [3]Vertex
	Normals [3]Vertex
}

// The trick to Barycentric Coordinates is to find the weights for V1, V2, and V3 that balance the following system of equations:
// Px = Wv1 * Xv1 + Wx2 * Xv2 + Wv3 * Xv3
// Py = Wx1 * Yv1 + Wv2 * Yv2 + Wv3 * Yv3
// Wv1 + Wv2 + Wv3 = 1
// Basically, for a given point P, we're finding weights that tell us how much of P's X coordinate is made of V1, V2, and V3, and also the same for P's Y coordinate.
// With a lot of re-arranging, we can solve Wv1, Wv2 and Wv3.
// Another thing about barycentric coordinates, is that if P is actually outside of the triangle, then at least one of W1, W2, or W3 will be negative!
func barycentric(pt, p1, p2, p3 image.Point) (float64, float64, float64) {
	p  := Vertex{X: float64(pt.X), Y: float64(pt.Y)}
	v1 := Vertex{X: float64(p1.X), Y: float64(p1.Y)}
	v2 := Vertex{X: float64(p2.X), Y: float64(p2.Y)}
	v3 := Vertex{X: float64(p3.X), Y: float64(p3.Y)}

	weightV1 := ((v2.Y - v3.Y) * (p.X - v3.X) + (v3.X - v2.X) * (p.Y - v3.Y)) / ((v2.Y - v3.Y) * (v1.X - v3.X) + (v3.X - v2.X) * (v1.Y - v3.Y))
	weightV2 := ((v3.Y - v1.Y) * (p.X - v3.X) + (v1.X - v3.X) * (p.Y - v3.Y)) / ((v2.Y - v3.Y) * (v1.X - v3.X) + (v3.X - v2.X) * (v1.Y - v3.Y))
	weightV3 := 1 - weightV1 - weightV2

	return weightV1, weightV2, weightV3
}