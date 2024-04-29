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
	Bbox   Aabb
}

func NewSphere(center Point3, radius float64, material Material) Sphere {
	rvec := NewVec3(radius, radius, radius)

	// logger := log.New(os.Stderr, "", 0)
	// logger.Print("!!!!")

	return Sphere{Center: center, Radius: radius, Mat: material, Bbox: NewAabbFromPoints(center.Sub(rvec), center.Add(rvec))}
}

func (s *Sphere) Hit(r *Ray, ray_t Interval, rec *Hit_record) bool {
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
	s.Get_sphere_uv(&outward_normal, &rec.U, &rec.V)

	// fmt.Print(rec.U, rec.V, "\n")

	rec.Mat = s.Mat

	return true

}

func (s *Sphere) Get_sphere_uv(p *Point3, u *float64, v *float64) {

	theta := math.Acos(-p.Y())
	phi := math.Atan2(-p.Z(), p.X()) + math.Pi

	*u = (phi / (2 * math.Pi))
	*v = theta / math.Pi

	// fmt.Print(*u, *v, "\n")

}

func (s *Sphere) Bounding_box() Aabb {
	return s.Bbox
}
