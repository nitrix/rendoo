package main

import "math"

type Vertex2 struct {
	X float64
	Y float64
}

type Vertex3 struct {
	X float64
	Y float64
	Z float64
}

func (v Vertex3) minus(o Vertex3) Vertex3 {
	return Vertex3{
		X: v.X - o.X,
		Y: v.Y - o.Y,
		Z: v.Z - o.Z,
	}
}

func (v Vertex3) cross(o Vertex3) Vertex3 {
	return Vertex3 {
		X: (v.Y * o.Z) - (v.Z * o.Y),
		Y: (v.Z * o.X) - (v.X * o.Z),
		Z: (v.X * o.Y) - (v.Y * o.X),
	}
}

type Vertex4 struct {
	X float64
	Y float64
	Z float64
	W float64
}

func (v Vertex3) normalize(target float64) Vertex3 {
	n := target / math.Sqrt(v.X*v.X+v.Y*v.Y+v.Z*v.Z)
	return Vertex3{
		X: v.X * n,
		Y: v.Y * n,
		Z: v.Z * n,
	}
}

func (v *Vertex4) transform(m Matrix4) {
	*v = Vertex4{
		X: v.X*m.m11 + v.Y*m.m12 + v.Z*m.m13 + v.W*m.m14,
		Y: v.X*m.m21 + v.Y*m.m22 + v.Z*m.m23 + v.W*m.m24,
		Z: v.X*m.m31 + v.Y*m.m32 + v.Z*m.m33 + v.W*m.m34,
		W: v.X*m.m41 + v.Y*m.m42 + v.Z*m.m43 + v.W*m.m44,
	}
}

func (v Vertex4) lower() Vertex3 {
	return Vertex3{
		X: v.X / v.W,
		Y: v.Y / v.W,
		Z: v.Z / v.W,
	}
}
