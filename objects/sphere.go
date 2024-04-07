package objects

import (
	"math"
	. "raytracer/common"
	. "raytracer/material"
)

type Sphere struct {
	Center Point3
	Radius float64
	Mat    Material
}

func (s *Sphere) Hit(r Ray, ray_t Interval, rec *Hit_record) bool {
	oc := r.Origin.Sub(s.Center)
	a := r.Direction.Length_squared()
	half_b := Dot(oc, r.Direction)
	c := oc.Length_squared() - s.Radius*s.Radius

	discriminant := half_b*half_b - a*c
	if discriminant < 0 {
		return false
	}
	sqrtd := math.Sqrt(discriminant)

	//Find nearest root in acceptable range

	root := (-half_b - sqrtd) / a

	if !ray_t.Surrounds(root) {
		root = (-half_b + sqrtd) / a
		if !ray_t.Surrounds(root) {
			return false
		}
	}

	rec.T = root
	rec.P = r.At(rec.T)

	outward_normal := rec.P.Sub(s.Center).Div(s.Radius)

	rec.Set_face_normal(r, outward_normal)

	rec.Mat = s.Mat

	return true

}
