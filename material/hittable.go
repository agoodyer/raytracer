package material

import (
	. "raytracer/common"
)

type Hit_record struct {
	P          Point3
	Normal     Vec3
	Mat        Material
	T          float64
	Front_face bool
}

type Hittable interface {
	Hit(r *Ray, ray_t Interval, rec *Hit_record) bool
	Bounding_box() Aabb
}

func (h *Hit_record) Set_face_normal(r *Ray, outward_normal Vec3) {
	h.Front_face = Dot(r.Direction, outward_normal) < 0
	if h.Front_face {
		h.Normal = outward_normal
	} else {
		h.Normal = NewVec3(0, 0, 0).Sub(outward_normal)
	}

	// logger := log.New(os.Stderr, "", 0)
	// logger.Print(h.normal)

}

type Hittable_list struct {
	Objects []Hittable
	bbox    Aabb
}

func (l *Hittable_list) Add(h Hittable) {
	l.Objects = append(l.Objects, h)
	l.bbox = Merge(l.bbox, h.Bounding_box())
}

func (l *Hittable_list) clear() {
	l.Objects = l.Objects[:0]
}

func (l *Hittable_list) Hit(r *Ray, ray_t Interval, rec *Hit_record) bool {

	var temp_rec Hit_record
	hit_anything := false
	closest_so_far := ray_t.Max

	for _, object := range l.Objects {

		if object.Hit(r, NewInterval(ray_t.Min, closest_so_far), &temp_rec) {
			hit_anything = true
			closest_so_far = temp_rec.T
			*rec = temp_rec
		}

	}

	return hit_anything

}

func (l *Hittable_list) Bounding_box() Aabb {
	return l.bbox
}
