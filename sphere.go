package main

import (
	"math"
)

type sphere struct {
	center point3
	radius float64
}

func (s *sphere) hit(r ray, ray_tmin float64, ray_tmax float64, rec *hit_record) bool {
	oc := r.origin.sub(s.center)
	a := r.direction.length_squared()
	half_b := dot(oc, r.direction)
	c := oc.length_squared() - s.radius*s.radius

	discriminant := half_b*half_b - a*c
	if discriminant < 0 {
		return false
	}
	sqrtd := math.Sqrt(discriminant)

	//Find nearest root in acceptable range

	root := (-half_b - sqrtd) / a

	if root <= ray_tmin || ray_tmax <= root {
		root = (-half_b + sqrtd) / a
		if root <= ray_tmin || ray_tmax <= root {
			return false
		}
	}

	rec.t = root
	rec.p = r.at(rec.t)

	outward_normal := rec.p.sub(s.center).div(s.radius)

	// logger := log.New(os.Stderr, "", 0)
	// logger.Print(outward_normal)

	rec.set_face_normal(r, outward_normal)

	return true

}
