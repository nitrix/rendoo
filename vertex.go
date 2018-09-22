package main

import "math"

type Vertex4 struct {
	X float64
	Y float64
	Z float64
	W float64
}

func (v *Vertex4) normalize(target float64) {
	 n := target / math.Sqrt(v.X*v.X + v.Y*v.Y + v.Z*v.Z)
	 *v = Vertex4{
	 	X: v.X * n,
	 	Y: v.Y * n,
	 	Z: v.Z * n,
	 }
}

func (v *Vertex4) transform(m Matrix4) {
	*v = Vertex4{
		X: v.X * m.m11 + v.Y * m.m12 + v.Z * m.m13 + v.W * m.m14,
		Y: v.X * m.m21 + v.Y * m.m22 + v.Z * m.m23 + v.W * m.m24,
		Z: v.X * m.m31 + v.Y * m.m32 + v.Z * m.m33 + v.W * m.m34,
		W: v.X * m.m41 + v.Y * m.m42 + v.Z * m.m43 + v.W * m.m44,
	}
}