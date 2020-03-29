package types

import "math"

type Sphere struct {
	Center Vec3f
	Radius float64
	Material Material
}

func (s *Sphere) RayIntersect(originalPoint Vec3f, directionPoint Vec3f, t0 *float64) bool {
	L := s.Center.Sub(originalPoint)
	tcaVec := L.Mul(directionPoint); tca := tcaVec.Dot()
	d2Vec := L.Mul(L); d2Vec = d2Vec.Sub(tcaVec.Scale(tca)); d2 := d2Vec.Dot()
	radius2 := s.Radius * s.Radius

	if d2 > radius2 {
		return false
	}

	thc := math.Sqrt(radius2 - d2)
	*t0 = tca - thc
	t1 := tca + thc

	if (*t0) < 0. {
		*t0 = t1
	}

	if (*t0) < 0 {
		return false
	}

	return true
}
