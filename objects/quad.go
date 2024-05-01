package objects

import (
	"math"
	. "raytracer/common"
	. "raytracer/material"
)

type Quad struct {
	q      Point3
	u, v   Vec3
	mat    Material
	Bbox   Aabb
	normal Vec3
	d      float64
	w      Vec3
}

func NewQuad(q Point3, u Vec3, v Vec3, material Material) Quad {

	n := Cross(u, v)
	normal := Unit_vector(n)
	d := Dot(normal, q)
	w := n.Div(Dot(n, n))
	quad := Quad{q: q, u: u, v: v, mat: material, w: w, d: d, normal: normal}
	quad.setBbox()
	return quad
}

func (quad *Quad) setBbox() {

	bbox_diagonal1 := NewAabbFromPoints(quad.q, quad.q.Add(quad.u).Add(quad.v))
	bbox_diagonal2 := NewAabbFromPoints(quad.q.Add(quad.u), quad.q.Add(quad.v))
	quad.Bbox = Merge(bbox_diagonal1, bbox_diagonal2)
}

func (quad *Quad) Hit(r *Ray, ray_t Interval, rec *Hit_record) bool {

	denom := Dot(quad.normal, r.Direction)

	// fmt.Println(quad.Bbox)

	if math.Abs(denom) < 1e-8 {
		return false
	}

	t := (quad.d - Dot(quad.normal, r.Origin)) / denom
	if !ray_t.Contains(t) {
		// fmt.Printf("no ")
		return false
	}

	intersection := r.At(t)
	planar_hitpt_vector := intersection.Sub(quad.q)

	alpha := Dot(quad.w, Cross(planar_hitpt_vector, quad.v))
	beta := Dot(quad.w, Cross(quad.u, planar_hitpt_vector))

	if !is_interior(alpha, beta, rec) {

		return false
	}

	rec.T = t
	rec.P = intersection
	rec.Mat = quad.mat
	rec.Set_face_normal(r, quad.normal)

	// fmt.Print("!!!!")
	return true
}

func is_interior(a float64, b float64, rec *Hit_record) bool {
	unit_interval := NewInterval(0, 1)

	if !unit_interval.Contains(a) || !unit_interval.Contains(b) {
		return false
	}

	rec.U = a
	rec.V = b

	return true

}

func (quad *Quad) Bounding_box() Aabb {
	return quad.Bbox
}

func NewBox(a Point3, b Point3, mat Material) Hittable_list {

	var sides Hittable_list

	min := NewPoint3(math.Min(a.X(), b.X()), math.Min(a.Y(), b.Y()), math.Min(a.Z(), b.Z()))
	max := NewPoint3(math.Max(a.X(), b.X()), math.Max(a.Y(), b.Y()), math.Max(a.Z(), b.Z()))

	dx := NewVec3(max.X()-min.X(), 0, 0)
	dy := NewVec3(0, max.Y()-min.Y(), 0)
	dz := NewVec3(0, 0, max.Z()-min.Z())

	front := NewQuad(NewPoint3(min.X(), min.Y(), max.Z()), dx, dy, mat)
	right := NewQuad(NewPoint3(max.X(), min.Y(), max.Z()), dz.Mult(-1), dy, mat)
	back := NewQuad(NewPoint3(max.X(), min.Y(), min.Z()), dx.Mult(-1), dy, mat)
	left := NewQuad(NewPoint3(min.X(), min.Y(), min.Z()), dz, dy, mat)
	top := NewQuad(NewPoint3(min.X(), max.Y(), max.Z()), dx, dz.Mult(-1), mat)
	bottom := NewQuad(NewPoint3(min.X(), min.Y(), min.Z()), dx, dz, mat)

	sides.Add(&front)
	sides.Add(&right)
	sides.Add(&back)
	sides.Add(&left)
	sides.Add(&top)
	sides.Add(&bottom)
	return sides

}

type Tri struct {
	q Quad
}

func NewTri(q Point3, u Vec3, v Vec3, material Material) Tri {
	return Tri{q: NewQuad(q, u, v, material)}
}

func (tri *Tri) Hit(r *Ray, ray_t Interval, rec *Hit_record) bool {
	denom := Dot(tri.q.normal, r.Direction)

	// fmt.Println(tri.q.Bbox)

	if math.Abs(denom) < 1e-8 {
		return false
	}

	t := (tri.q.d - Dot(tri.q.normal, r.Origin)) / denom
	if !ray_t.Contains(t) {
		// fmt.Printf("no ")
		return false
	}

	intersection := r.At(t)
	planar_hitpt_vector := intersection.Sub(tri.q.q)

	alpha := Dot(tri.q.w, Cross(planar_hitpt_vector, tri.q.v))
	beta := Dot(tri.q.w, Cross(tri.q.u, planar_hitpt_vector))

	if !is_interiorTri(alpha, beta, rec) {

		return false
	}

	rec.T = t
	rec.P = intersection
	rec.Mat = tri.q.mat
	rec.Set_face_normal(r, tri.q.normal)

	// fmt.Print("!!!!")
	return true

}

func is_interiorTri(a float64, b float64, rec *Hit_record) bool {

	if a < 0 || b < 0 || a+b > 1 {
		return false
	}

	rec.U = a
	rec.V = b

	return true

}

func (tri *Tri) Bounding_box() Aabb {
	return tri.q.Bbox
}

func Triangle(a Point3, b Point3, c Point3, material Material) Tri {

	return NewTri(a, c.Sub(a), b.Sub(a), material)

}
