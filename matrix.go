package main

type Matrix4 struct {
	m11, m12, m13, m14 float64
	m21, m22, m23, m24 float64
	m31, m32, m33, m34 float64
	m41, m42, m43, m44 float64
}

func Identity4() Matrix4 {
	return Matrix4{
		m11: 1.0,
		m22: 1.0,
		m33: 1.0,
		m44: 1.0,
	}
}

func (m Matrix4) Dot(o Matrix4) Matrix4 {
	return Matrix4{
		m11: (m.m11 * o.m11) + (m.m12 * o.m21) + (m.m13 * o.m31) + (m.m14 * o.m41),
		m12: (m.m11 * o.m12) + (m.m12 * o.m22) + (m.m13 * o.m32) + (m.m14 * o.m42),
		m13: (m.m11 * o.m13) + (m.m12 * o.m23) + (m.m13 * o.m33) + (m.m14 * o.m43),
		m14: (m.m11 * o.m14) + (m.m12 * o.m24) + (m.m13 * o.m34) + (m.m14 * o.m44),

		m21: (m.m21 * o.m21) + (m.m22 * o.m21) + (m.m23 * o.m31) + (m.m24 * o.m41),
		m22: (m.m21 * o.m12) + (m.m22 * o.m22) + (m.m23 * o.m32) + (m.m24 * o.m42),
		m23: (m.m21 * o.m13) + (m.m22 * o.m23) + (m.m23 * o.m33) + (m.m24 * o.m43),
		m24: (m.m21 * o.m14) + (m.m22 * o.m24) + (m.m23 * o.m34) + (m.m24 * o.m44),

		m31: (m.m31 * o.m21) + (m.m32 * o.m21) + (m.m33 * o.m31) + (m.m34 * o.m41),
		m32: (m.m31 * o.m12) + (m.m32 * o.m22) + (m.m33 * o.m32) + (m.m34 * o.m42),
		m33: (m.m31 * o.m13) + (m.m32 * o.m23) + (m.m33 * o.m33) + (m.m34 * o.m43),
		m34: (m.m31 * o.m14) + (m.m32 * o.m24) + (m.m33 * o.m34) + (m.m34 * o.m44),

		m41: (m.m41 * o.m21) + (m.m42 * o.m21) + (m.m43 * o.m31) + (m.m44 * o.m41),
		m42: (m.m41 * o.m12) + (m.m42 * o.m22) + (m.m43 * o.m32) + (m.m44 * o.m42),
		m43: (m.m41 * o.m13) + (m.m42 * o.m23) + (m.m43 * o.m33) + (m.m44 * o.m43),
		m44: (m.m41 * o.m14) + (m.m42 * o.m24) + (m.m43 * o.m34) + (m.m44 * o.m44),
	}
}