package types

import "math"

//Vec3f Three-component floating point vector
type Vec3f struct {
	V0 float64
	V1 float64
	V2 float64
}

func (v *Vec3f) clone() Vec3f {
	return Vec3f{
		v.V0,
		v.V1,
		v.V2,
	}
}

func (v *Vec3f) Norm() float64 {
	return math.Sqrt(v.V0*v.V0 + v.V1*v.V1 + v.V2*v.V2)
}

//Normalize Get normal
func (v *Vec3f) Normalize() Vec3f {
	n := v.Norm()

	if n == 0 {
		return Vec3f{0, 0, 0}
	}

	return v.Scale(1.0 / n)
}

//Add Adding vector to vector
func (v *Vec3f) Add(other Vec3f) Vec3f {
	vv := v.clone()

	vv.V0 += other.V0
	vv.V1 += other.V1
	vv.V2 += other.V2

	return vv
}

//Sub Subtraction vector from vector
func (v *Vec3f) Sub(other Vec3f) Vec3f {
	vv := v.clone()

	vv.V0 -= other.V0
	vv.V1 -= other.V1
	vv.V2 -= other.V2

	return vv
}

//Mul Multiple vectors
func (v *Vec3f) Mul(other Vec3f) Vec3f {
	vv := v.clone()

	vv.V0 *= other.V0
	vv.V1 *= other.V1
	vv.V2 *= other.V2

	return vv
}

//Scale Change scale factor
func (v *Vec3f) Scale(factor float64) Vec3f {
	vv := v.clone()

	vv.V0 *= factor
	vv.V1 *= factor
	vv.V2 *= factor

	return vv
}

//Reverse Get negative vector
func (v *Vec3f) Reverse() Vec3f {
	vv := v.clone()

	vv.V0 = -v.V0
	vv.V1 = -v.V1
	vv.V2 = -v.V2

	return vv
}

func (v *Vec3f) Dot() float64 {
	return v.V0 + v.V1 + v.V2
}
