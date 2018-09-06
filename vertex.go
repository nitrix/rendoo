package main

import "math"

type Vertex struct {
	X float64
	Y float64
	Z float64
}

func (v *Vertex) normalize(target float64) {
	 n := target / math.Sqrt(v.X*v.X + v.Y*v.Y + v.Z*v.Z)
	 *v = Vertex{
	 	X: v.X * n,
	 	Y: v.Y * n,
	 	Z: v.Z * n,
	 }
}