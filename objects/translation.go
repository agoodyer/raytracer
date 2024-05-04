package objects

import (
	. "raytracer/common"
	. "raytracer/material"
)

type Translation struct {
	object Hittable
	offset Vec3
	bbox   Aabb
}

func NewTranslation(object Hittable, offset Vec3) Translation {

	bbox := object.Bounding_box()
	bbox = bbox.Add(offset)
	return Translation{object: object, offset: offset, bbox: bbox}

}

func (tr *Translation) Hit(r *Ray, ray_t Interval, rec *Hit_record) bool {

	ray_offset_r := NewRay(r.Origin.Sub(tr.offset), r.Direction)

	if tr.object.Hit(&ray_offset_r, ray_t, rec) {
		return false
	}
	rec.P.Add(tr.offset)
	return true

}

func (tr *Translation) Bounding_box() Aabb {
	return tr.bbox
}
